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

package gcp

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/gravitational/trace"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/utils/keys"
)

// SQLAdminClient defines an interface providing access to the GCP Cloud SQL API.
type SQLAdminClient interface {
	// GetUser retrieves a resource containing information about a user.
	GetUser(ctx context.Context, db types.Database, dbUser string) (*sqladmin.User, error)
	// UpdateUser updates an existing user for the project/instance configured in a session.
	UpdateUser(ctx context.Context, db types.Database, dbUser string, user *sqladmin.User) error
	// GetDatabaseInstance returns database instance details for the project/instance
	// configured in a session.
	GetDatabaseInstance(ctx context.Context, db types.Database) (*sqladmin.DatabaseInstance, error)
	// GenerateEphemeralCert returns a new PEM-encoded client certificate for
	// the project/instance configured in a session.
	GenerateEphemeralCert(ctx context.Context, db types.Database, certExpiry time.Time, pubKey crypto.PublicKey) (string, error)
}

// NewGCPSQLAdminClient returns a GCPSQLAdminClient interface wrapping sqladmin.Service.
func NewSQLAdminClient(ctx context.Context) (SQLAdminClient, error) {
	service, err := sqladmin.NewService(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &gcpSQLAdminClient{service: service}, nil
}

// gcpSQLAdminClient implements the GCPSQLAdminClient interface by wrapping
// sqladmin.Service.
type gcpSQLAdminClient struct {
	service *sqladmin.Service
}

// GetUser retrieves a resource containing information about a user.
func (g *gcpSQLAdminClient) GetUser(ctx context.Context, db types.Database, dbUser string) (*sqladmin.User, error) {
	user, err := g.service.Users.Get(
		db.GetGCP().ProjectID,
		db.GetGCP().InstanceID,
		dbUser,
	).Host("%").Context(ctx).Do()
	return user, trace.Wrap(convertAPIError(err))
}

// UpdateUser updates an existing user in a Cloud SQL for the project/instance
// configured in a session.
func (g *gcpSQLAdminClient) UpdateUser(ctx context.Context, db types.Database, dbUser string, user *sqladmin.User) error {
	_, err := g.service.Users.Update(
		db.GetGCP().ProjectID,
		db.GetGCP().InstanceID,
		user).Name(dbUser).Host("%").Context(ctx).Do()
	if err != nil {
		return trace.Wrap(convertAPIError(err))
	}
	return nil
}

// GetDatabaseInstance returns database instance details from Cloud SQL for the
// project/instance configured in a session.
func (g *gcpSQLAdminClient) GetDatabaseInstance(ctx context.Context, db types.Database) (*sqladmin.DatabaseInstance, error) {
	gcp := db.GetGCP()
	dbi, err := g.service.Instances.Get(gcp.ProjectID, gcp.InstanceID).Context(ctx).Do()
	if err != nil {
		return nil, trace.Wrap(convertAPIError(err))
	}
	return dbi, nil
}

// GenerateEphemeralCert returns a new client certificate created using the
// GenerateEphemeralCertRequest Cloud SQL API. Client certificates are required
// when enabling SSL in Cloud SQL.
func (g *gcpSQLAdminClient) GenerateEphemeralCert(ctx context.Context, db types.Database, certExpiry time.Time, pubKey crypto.PublicKey) (string, error) {
	// TODO(jimbishopp): cache database certificates to avoid expensive generate
	// operation on each connection.

	var keyPEM []byte
	switch pubKey.(type) {
	case *rsa.PublicKey:
		// keys.MarshalPublicKey would use PKCS#1 format for an RSA public key,
		// we specifically want PKIX here.
		pkix, err := x509.MarshalPKIXPublicKey(pubKey)
		if err != nil {
			return "", trace.Wrap(err)
		}
		keyPEM = pem.EncodeToMemory(&pem.Block{Bytes: pkix, Type: "RSA PUBLIC KEY"})
	default:
		var err error
		keyPEM, err = keys.MarshalPublicKey(pubKey)
		if err != nil {
			return "", trace.Wrap(err)
		}
	}

	// Make API call.
	gcp := db.GetGCP()
	req := g.service.Connect.GenerateEphemeralCert(gcp.ProjectID, gcp.InstanceID, &sqladmin.GenerateEphemeralCertRequest{
		PublicKey:     string(keyPEM),
		ValidDuration: fmt.Sprintf("%ds", int(time.Until(certExpiry).Seconds())),
	})
	resp, err := req.Context(ctx).Do()
	if err != nil {
		return "", trace.Wrap(convertAPIError(err))
	}

	return resp.EphemeralCert.Cert, nil
}
