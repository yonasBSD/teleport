/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package app

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"strings"

	"github.com/gravitational/trace"
	"github.com/julienschmidt/httprouter"

	"github.com/gravitational/teleport/api/types"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/events"
	"github.com/gravitational/teleport/lib/httplib"
	"github.com/gravitational/teleport/lib/utils"
)

type fragmentRequest struct {
	StateValue         string `json:"state_value"`
	CookieValue        string `json:"cookie_value"`
	SubjectCookieValue string `json:"subject_cookie_value"`
	RequiredApps       string `json:"required_apps"`
}

// TeleportNextAppRedirectUrlHeader is used to tell the browser which URL to navigate to
// next in the chain of required app authentication redirects
const TeleportNextAppRedirectUrlHeader = "X-Teleport-NextAppRedirectUrl"

// startAppAuthExchange will do two actions depending on the following:
//
//	1): On initiating auth exchange (indicated by an empty "state" query param)
//	    we create a crypto safe random token and send it back as part of a "state"
//	    query param in the redirection URL, as well as in a cookie with attributes
//	    that makes the cookie unaccesible and hard to tamper with. We use this
//	    "double submit cookie" method to protect the entire auth exchange flow
//	    from CSRF.
//
//	2): If the "state" query param is present, we will serve a blank HTML page
//	    that has inline JS that contains logic to complete the auth exchange.
func (h *Handler) startAppAuthExchange(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	q := r.URL.Query()

	requiredApps := strings.Split(q.Get("required-apps"), ",")
	reqAddr, err := utils.ParseAddr(r.Host)
	if err != nil {
		return trace.Wrap(err)
	}

	// if required apps is length 1, it _should_ be itself as the final stop of the redirect chain.
	// If that is the case, skip any further required app redirects and complete the auth flow.
	// This value gets shifted off the front after a complete auth exchange
	if len(requiredApps) > 1 && requiredApps[0] != reqAddr.Host() {
		nextRequiredApp := requiredApps[0]

		webLauncherURLParams := launcherURLParams{
			clusterName:         q.Get("cluster"),
			publicAddr:          nextRequiredApp,
			arn:                 q.Get("arn"),
			path:                q.Get("path"),
			requiredAppFQDNs:    strings.Join(requiredApps, ","),
			requiresAppRedirect: true,
		}
		return h.redirectToLauncher(w, r, webLauncherURLParams)
	}
	// Initiate auth exchange.
	if q.Get("state") == "" {
		// secretToken is the token we will look for in both the cookie
		// and in the request "state" query param.
		secretToken, err := utils.CryptoRandomHex(defaults.TokenLenBytes)
		if err != nil {
			h.logger.DebugContext(r.Context(), "Failed to generate token required for app auth exchange", "error", err)
			return trace.AccessDenied("access denied")
		}

		// cookieIdentifier is used to uniquely identify this cookie
		// that will be used to store this secret token.
		//
		// This prevents a race condition (state token mismatch error)
		// where we can overwrite existing cookie (with the same name) with a
		// different token value eg: launch app in multiple tabs in quick succession
		cookieIdentifier, err := utils.CryptoRandomHex(defaults.TokenLenBytes)
		if err != nil {
			h.logger.DebugContext(r.Context(), "Failed to generate an UID for the app auth state cookie", "error", err)
			return trace.AccessDenied("access denied")
		}

		h.setAuthStateCookie(w, secretToken, cookieIdentifier)

		webLauncherURLParams := launcherURLParams{
			clusterName: q.Get("cluster"),
			publicAddr:  q.Get("addr"),
			arn:         q.Get("arn"),
			path:        q.Get("path"),
			// The state token concats both the secret token and the cookie ID.
			// The server will break this token to its individual parts:
			//   - secretToken to compare against the one stored in cookie
			//   - cookieIdentifier to look up cookie sent by browser.
			stateToken: fmt.Sprintf("%s_%s", secretToken, cookieIdentifier),
			// required apps
			requiredAppFQDNs: q.Get("required-apps"),
		}
		return h.redirectToLauncher(w, r, webLauncherURLParams)
	}

	// Continue the auth exchange.

	nonce, err := utils.CryptoRandomHex(defaults.TokenLenBytes)
	if err != nil {
		h.logger.DebugContext(r.Context(), "Failed to create a random nonce for the app redirection HTML inline script", "error", err)
		return trace.AccessDenied("access denied")
	}
	SetRedirectPageHeaders(w.Header(), nonce)

	// Serving the HTML page.
	if err := appRedirectTemplate.Execute(w, nonce); err != nil {
		h.logger.DebugContext(r.Context(), "Failed executing appRedirect template", "error", err)
		return trace.AccessDenied("access denied")
	}
	return nil
}

// completeAppAuthExchange completes the auth exchange flow started by "startAppAuthExchange" handler
// by validating the values passed in the request body, and upon success sets cookies required
// for the current app session.
func (h *Handler) completeAppAuthExchange(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	httplib.SetNoCacheHeaders(w.Header())
	var req fragmentRequest
	if err := httplib.ReadJSON(r, &req); err != nil {
		h.logger.DebugContext(r.Context(), "Failed to decode JSON from request body", "error", err)
		return trace.AccessDenied("access denied")
	}

	secretToken, cookieID, ok := strings.Cut(req.StateValue, "_")
	if !ok {
		h.logger.WarnContext(r.Context(), "Request failed: request state token is not in the expected format")
		h.emitErrorEventAndDeleteAppSession(r, emitErrorEventFields{
			sessionID: req.CookieValue,
			err:       "state token was not in the expected format",
		})
		return trace.AccessDenied("access denied")
	}

	// Validate that the caller-provided state token matches the stored state token (CSRF check)
	stateCookie, err := r.Cookie(getAuthStateCookieName(cookieID))
	if err != nil || stateCookie.Value == "" {
		h.logger.WarnContext(r.Context(), "Request failed: state cookie is not set")
		h.emitErrorEventAndDeleteAppSession(r, emitErrorEventFields{
			sessionID: req.CookieValue,
			err:       "auth state cookie was not set",
		})
		return trace.AccessDenied("access denied")
	}
	if subtle.ConstantTimeCompare([]byte(secretToken), []byte(stateCookie.Value)) != 1 {
		h.logger.WarnContext(r.Context(), "Request failed: state token does not match")
		h.emitErrorEventAndDeleteAppSession(r, emitErrorEventFields{
			sessionID: req.CookieValue,
			err:       "state token did not match",
		})
		return trace.AccessDenied("access denied")
	}

	// Prevent reuse of the same state token.
	clearAuthStateCookie(w, cookieID)

	// Validate that the caller is asking for a session that exists and that they have the secret
	// session token for.
	ws, err := h.getAppSessionFromAccessPoint(r.Context(), req.CookieValue)
	if err != nil {
		h.logger.WarnContext(r.Context(), "Request failed: session does not exist", "error", err)
		return trace.AccessDenied("access denied")
	}
	if err := checkSubjectToken(req.SubjectCookieValue, ws); err != nil {
		h.logger.WarnContext(r.Context(), "Request failed", "error", err)
		h.emitErrorEventAndDeleteAppSession(r, emitErrorEventFields{
			sessionID: req.CookieValue,
			err:       err.Error(),
			loginName: ws.GetUser(),
		})
		return trace.AccessDenied("access denied")
	}

	// Set the "Set-Cookie" header on the response.
	// Set Same-Site policy for the session cookies to None in order to
	// support redirects that identity providers do during SSO auth.
	// Otherwise the session cookie won't be sent and the user will
	// get redirected to the application launcher.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    req.CookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     SubjectCookieName,
		Value:    ws.GetBearerToken(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	requiredApps := strings.Split(req.RequiredApps, ",")
	if len(requiredApps) > 1 {
		requiredApps := requiredApps[1:]
		nextRequiredApp := requiredApps[0]

		webLauncherURLParams := launcherURLParams{
			publicAddr:          nextRequiredApp,
			requiredAppFQDNs:    strings.Join(requiredApps, ","),
			requiresAppRedirect: true,
		}
		addr, err := utils.ParseAddr(webLauncherURLParams.publicAddr)
		if err != nil {
			return trace.Wrap(err)
		}

		var publicAddrs []string
		for _, addr := range h.c.ProxyPublicAddrs {
			publicAddrs = append(publicAddrs, addr.Host())
		}
		proxyDNSName := utils.FindMatchingProxyDNS(addr.String(), publicAddrs)
		urlString := makeAppRedirectURL(r, proxyDNSName, addr.Host(), webLauncherURLParams)
		// this request does not return a response, so we can pass this value through a custom header instead
		w.Header().Set(TeleportNextAppRedirectUrlHeader, urlString)
	}

	return nil
}

func checkSubjectToken(subjectCookieValue string, ws types.WebSession) error {
	if subjectCookieValue == "" {
		return trace.AccessDenied("subject session token is not set")
	}
	if subtle.ConstantTimeCompare([]byte(subjectCookieValue), []byte(ws.GetBearerToken())) != 1 {
		return trace.AccessDenied("subject session token does not match")
	}
	return nil
}

func (h *Handler) setAuthStateCookie(w http.ResponseWriter, cookieValue string, cookieID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     getAuthStateCookieName(cookieID),
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   60, // Expire in 1 minute.
	})
}

func clearAuthStateCookie(w http.ResponseWriter, cookieID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     getAuthStateCookieName(cookieID),
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	})
}

func getAuthStateCookieName(cookieID string) string {
	return fmt.Sprintf("%s_%s", AuthStateCookieName, cookieID)
}

type emitErrorEventFields struct {
	loginName string
	err       string
	sessionID string
}

func (h *Handler) emitErrorEventAndDeleteAppSession(r *http.Request, f emitErrorEventFields) {
	// Attempt to delete app session.
	if f.sessionID != "" {
		if err := h.c.AuthClient.DeleteAppSession(r.Context(), types.DeleteAppSessionRequest{
			SessionID: f.sessionID,
		}); err != nil {
			h.logger.WarnContext(r.Context(), "Failed to delete app session after failing authentication", "error", err)
		}
	}

	event := &apievents.AuthAttempt{
		Metadata: apievents.Metadata{
			Type: events.AuthAttemptEvent,
			Code: events.AuthAttemptFailureCode,
		},
		UserMetadata: apievents.UserMetadata{
			User:  "unknown",
			Login: f.loginName,
		},
		ConnectionMetadata: apievents.ConnectionMetadata{
			LocalAddr:  r.Host,
			RemoteAddr: r.RemoteAddr,
		},
		Status: apievents.Status{
			Success: false,
			Error:   fmt.Sprintf("Failed app access authentication: %s", f.err),
		},
	}

	h.c.AuthClient.EmitAuditEvent(h.closeContext, event)
}
