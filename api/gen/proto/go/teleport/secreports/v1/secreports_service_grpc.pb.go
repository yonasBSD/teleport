// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: teleport/secreports/v1/secreports_service.proto

package secreportsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SecReportsService_UpsertAuditQuery_FullMethodName    = "/teleport.secreports.v1.SecReportsService/UpsertAuditQuery"
	SecReportsService_GetAuditQuery_FullMethodName       = "/teleport.secreports.v1.SecReportsService/GetAuditQuery"
	SecReportsService_ListAuditQueries_FullMethodName    = "/teleport.secreports.v1.SecReportsService/ListAuditQueries"
	SecReportsService_DeleteAuditQuery_FullMethodName    = "/teleport.secreports.v1.SecReportsService/DeleteAuditQuery"
	SecReportsService_UpsertReport_FullMethodName        = "/teleport.secreports.v1.SecReportsService/UpsertReport"
	SecReportsService_GetReport_FullMethodName           = "/teleport.secreports.v1.SecReportsService/GetReport"
	SecReportsService_ListReports_FullMethodName         = "/teleport.secreports.v1.SecReportsService/ListReports"
	SecReportsService_DeleteReport_FullMethodName        = "/teleport.secreports.v1.SecReportsService/DeleteReport"
	SecReportsService_RunAuditQuery_FullMethodName       = "/teleport.secreports.v1.SecReportsService/RunAuditQuery"
	SecReportsService_GetAuditQueryResult_FullMethodName = "/teleport.secreports.v1.SecReportsService/GetAuditQueryResult"
	SecReportsService_RunReport_FullMethodName           = "/teleport.secreports.v1.SecReportsService/RunReport"
	SecReportsService_GetReportResult_FullMethodName     = "/teleport.secreports.v1.SecReportsService/GetReportResult"
	SecReportsService_GetReportState_FullMethodName      = "/teleport.secreports.v1.SecReportsService/GetReportState"
	SecReportsService_ListReportStates_FullMethodName    = "/teleport.secreports.v1.SecReportsService/ListReportStates"
	SecReportsService_GetSchema_FullMethodName           = "/teleport.secreports.v1.SecReportsService/GetSchema"
)

// SecReportsServiceClient is the client API for SecReportsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SecReportsService is a service that manages security reports.
type SecReportsServiceClient interface {
	// UpsertAuditQuery upsets an audit query.
	UpsertAuditQuery(ctx context.Context, in *UpsertAuditQueryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetAuditQuery returns an audit query.
	GetAuditQuery(ctx context.Context, in *GetAuditQueryRequest, opts ...grpc.CallOption) (*AuditQuery, error)
	// ListAuditQueries returns a paginated list of audit query resources.
	ListAuditQueries(ctx context.Context, in *ListAuditQueriesRequest, opts ...grpc.CallOption) (*ListAuditQueriesResponse, error)
	// DeleteAuditQuery deletes an audit query.
	DeleteAuditQuery(ctx context.Context, in *DeleteAuditQueryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UpsertReport upsets a report.
	UpsertReport(ctx context.Context, in *UpsertReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetReport returns a report.
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error)
	// ListReports returns a paginated list of security report resources.
	ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error)
	// DeleteReport deletes a security report.
	DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// RunAuditQuery runs an audit query.
	RunAuditQuery(ctx context.Context, in *RunAuditQueryRequest, opts ...grpc.CallOption) (*RunAuditQueryResponse, error)
	// GetAuditQueryResult returns an audit query result.
	GetAuditQueryResult(ctx context.Context, in *GetAuditQueryResultRequest, opts ...grpc.CallOption) (*GetAuditQueryResultResponse, error)
	// RunReport runs a security report.
	RunReport(ctx context.Context, in *RunReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetReportResult returns a security report result.
	GetReportResult(ctx context.Context, in *GetReportResultRequest, opts ...grpc.CallOption) (*GetReportResultResponse, error)
	// GetReportState returns a security report state.
	GetReportState(ctx context.Context, in *GetReportStateRequest, opts ...grpc.CallOption) (*ReportState, error)
	// ListReportStates returns a paginated list of security report state resources.
	ListReportStates(ctx context.Context, in *ListReportStatesRequest, opts ...grpc.CallOption) (*ListReportStatesResponse, error)
	// GetSchema returns a schema of audit query.
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
}

type secReportsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecReportsServiceClient(cc grpc.ClientConnInterface) SecReportsServiceClient {
	return &secReportsServiceClient{cc}
}

func (c *secReportsServiceClient) UpsertAuditQuery(ctx context.Context, in *UpsertAuditQueryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SecReportsService_UpsertAuditQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetAuditQuery(ctx context.Context, in *GetAuditQueryRequest, opts ...grpc.CallOption) (*AuditQuery, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditQuery)
	err := c.cc.Invoke(ctx, SecReportsService_GetAuditQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) ListAuditQueries(ctx context.Context, in *ListAuditQueriesRequest, opts ...grpc.CallOption) (*ListAuditQueriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAuditQueriesResponse)
	err := c.cc.Invoke(ctx, SecReportsService_ListAuditQueries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) DeleteAuditQuery(ctx context.Context, in *DeleteAuditQueryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SecReportsService_DeleteAuditQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) UpsertReport(ctx context.Context, in *UpsertReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SecReportsService_UpsertReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Report)
	err := c.cc.Invoke(ctx, SecReportsService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) ListReports(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReportsResponse)
	err := c.cc.Invoke(ctx, SecReportsService_ListReports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SecReportsService_DeleteReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) RunAuditQuery(ctx context.Context, in *RunAuditQueryRequest, opts ...grpc.CallOption) (*RunAuditQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RunAuditQueryResponse)
	err := c.cc.Invoke(ctx, SecReportsService_RunAuditQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetAuditQueryResult(ctx context.Context, in *GetAuditQueryResultRequest, opts ...grpc.CallOption) (*GetAuditQueryResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuditQueryResultResponse)
	err := c.cc.Invoke(ctx, SecReportsService_GetAuditQueryResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) RunReport(ctx context.Context, in *RunReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SecReportsService_RunReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetReportResult(ctx context.Context, in *GetReportResultRequest, opts ...grpc.CallOption) (*GetReportResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReportResultResponse)
	err := c.cc.Invoke(ctx, SecReportsService_GetReportResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetReportState(ctx context.Context, in *GetReportStateRequest, opts ...grpc.CallOption) (*ReportState, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportState)
	err := c.cc.Invoke(ctx, SecReportsService_GetReportState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) ListReportStates(ctx context.Context, in *ListReportStatesRequest, opts ...grpc.CallOption) (*ListReportStatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListReportStatesResponse)
	err := c.cc.Invoke(ctx, SecReportsService_ListReportStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secReportsServiceClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, SecReportsService_GetSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecReportsServiceServer is the server API for SecReportsService service.
// All implementations must embed UnimplementedSecReportsServiceServer
// for forward compatibility.
//
// SecReportsService is a service that manages security reports.
type SecReportsServiceServer interface {
	// UpsertAuditQuery upsets an audit query.
	UpsertAuditQuery(context.Context, *UpsertAuditQueryRequest) (*emptypb.Empty, error)
	// GetAuditQuery returns an audit query.
	GetAuditQuery(context.Context, *GetAuditQueryRequest) (*AuditQuery, error)
	// ListAuditQueries returns a paginated list of audit query resources.
	ListAuditQueries(context.Context, *ListAuditQueriesRequest) (*ListAuditQueriesResponse, error)
	// DeleteAuditQuery deletes an audit query.
	DeleteAuditQuery(context.Context, *DeleteAuditQueryRequest) (*emptypb.Empty, error)
	// UpsertReport upsets a report.
	UpsertReport(context.Context, *UpsertReportRequest) (*emptypb.Empty, error)
	// GetReport returns a report.
	GetReport(context.Context, *GetReportRequest) (*Report, error)
	// ListReports returns a paginated list of security report resources.
	ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error)
	// DeleteReport deletes a security report.
	DeleteReport(context.Context, *DeleteReportRequest) (*emptypb.Empty, error)
	// RunAuditQuery runs an audit query.
	RunAuditQuery(context.Context, *RunAuditQueryRequest) (*RunAuditQueryResponse, error)
	// GetAuditQueryResult returns an audit query result.
	GetAuditQueryResult(context.Context, *GetAuditQueryResultRequest) (*GetAuditQueryResultResponse, error)
	// RunReport runs a security report.
	RunReport(context.Context, *RunReportRequest) (*emptypb.Empty, error)
	// GetReportResult returns a security report result.
	GetReportResult(context.Context, *GetReportResultRequest) (*GetReportResultResponse, error)
	// GetReportState returns a security report state.
	GetReportState(context.Context, *GetReportStateRequest) (*ReportState, error)
	// ListReportStates returns a paginated list of security report state resources.
	ListReportStates(context.Context, *ListReportStatesRequest) (*ListReportStatesResponse, error)
	// GetSchema returns a schema of audit query.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	mustEmbedUnimplementedSecReportsServiceServer()
}

// UnimplementedSecReportsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSecReportsServiceServer struct{}

func (UnimplementedSecReportsServiceServer) UpsertAuditQuery(context.Context, *UpsertAuditQueryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAuditQuery not implemented")
}
func (UnimplementedSecReportsServiceServer) GetAuditQuery(context.Context, *GetAuditQueryRequest) (*AuditQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditQuery not implemented")
}
func (UnimplementedSecReportsServiceServer) ListAuditQueries(context.Context, *ListAuditQueriesRequest) (*ListAuditQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuditQueries not implemented")
}
func (UnimplementedSecReportsServiceServer) DeleteAuditQuery(context.Context, *DeleteAuditQueryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuditQuery not implemented")
}
func (UnimplementedSecReportsServiceServer) UpsertReport(context.Context, *UpsertReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertReport not implemented")
}
func (UnimplementedSecReportsServiceServer) GetReport(context.Context, *GetReportRequest) (*Report, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedSecReportsServiceServer) ListReports(context.Context, *ListReportsRequest) (*ListReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReports not implemented")
}
func (UnimplementedSecReportsServiceServer) DeleteReport(context.Context, *DeleteReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReport not implemented")
}
func (UnimplementedSecReportsServiceServer) RunAuditQuery(context.Context, *RunAuditQueryRequest) (*RunAuditQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunAuditQuery not implemented")
}
func (UnimplementedSecReportsServiceServer) GetAuditQueryResult(context.Context, *GetAuditQueryResultRequest) (*GetAuditQueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditQueryResult not implemented")
}
func (UnimplementedSecReportsServiceServer) RunReport(context.Context, *RunReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunReport not implemented")
}
func (UnimplementedSecReportsServiceServer) GetReportResult(context.Context, *GetReportResultRequest) (*GetReportResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportResult not implemented")
}
func (UnimplementedSecReportsServiceServer) GetReportState(context.Context, *GetReportStateRequest) (*ReportState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportState not implemented")
}
func (UnimplementedSecReportsServiceServer) ListReportStates(context.Context, *ListReportStatesRequest) (*ListReportStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReportStates not implemented")
}
func (UnimplementedSecReportsServiceServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedSecReportsServiceServer) mustEmbedUnimplementedSecReportsServiceServer() {}
func (UnimplementedSecReportsServiceServer) testEmbeddedByValue()                           {}

// UnsafeSecReportsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecReportsServiceServer will
// result in compilation errors.
type UnsafeSecReportsServiceServer interface {
	mustEmbedUnimplementedSecReportsServiceServer()
}

func RegisterSecReportsServiceServer(s grpc.ServiceRegistrar, srv SecReportsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSecReportsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SecReportsService_ServiceDesc, srv)
}

func _SecReportsService_UpsertAuditQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAuditQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).UpsertAuditQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_UpsertAuditQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).UpsertAuditQuery(ctx, req.(*UpsertAuditQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetAuditQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetAuditQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetAuditQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetAuditQuery(ctx, req.(*GetAuditQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_ListAuditQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuditQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).ListAuditQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_ListAuditQueries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).ListAuditQueries(ctx, req.(*ListAuditQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_DeleteAuditQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuditQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).DeleteAuditQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_DeleteAuditQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).DeleteAuditQuery(ctx, req.(*DeleteAuditQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_UpsertReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).UpsertReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_UpsertReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).UpsertReport(ctx, req.(*UpsertReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_ListReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).ListReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_ListReports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).ListReports(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_DeleteReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).DeleteReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_DeleteReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).DeleteReport(ctx, req.(*DeleteReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_RunAuditQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunAuditQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).RunAuditQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_RunAuditQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).RunAuditQuery(ctx, req.(*RunAuditQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetAuditQueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditQueryResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetAuditQueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetAuditQueryResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetAuditQueryResult(ctx, req.(*GetAuditQueryResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_RunReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).RunReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_RunReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).RunReport(ctx, req.(*RunReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetReportResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetReportResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetReportResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetReportResult(ctx, req.(*GetReportResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetReportState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetReportState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetReportState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetReportState(ctx, req.(*GetReportStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_ListReportStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).ListReportStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_ListReportStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).ListReportStates(ctx, req.(*ListReportStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecReportsService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecReportsServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecReportsService_GetSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecReportsServiceServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecReportsService_ServiceDesc is the grpc.ServiceDesc for SecReportsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecReportsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.secreports.v1.SecReportsService",
	HandlerType: (*SecReportsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertAuditQuery",
			Handler:    _SecReportsService_UpsertAuditQuery_Handler,
		},
		{
			MethodName: "GetAuditQuery",
			Handler:    _SecReportsService_GetAuditQuery_Handler,
		},
		{
			MethodName: "ListAuditQueries",
			Handler:    _SecReportsService_ListAuditQueries_Handler,
		},
		{
			MethodName: "DeleteAuditQuery",
			Handler:    _SecReportsService_DeleteAuditQuery_Handler,
		},
		{
			MethodName: "UpsertReport",
			Handler:    _SecReportsService_UpsertReport_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _SecReportsService_GetReport_Handler,
		},
		{
			MethodName: "ListReports",
			Handler:    _SecReportsService_ListReports_Handler,
		},
		{
			MethodName: "DeleteReport",
			Handler:    _SecReportsService_DeleteReport_Handler,
		},
		{
			MethodName: "RunAuditQuery",
			Handler:    _SecReportsService_RunAuditQuery_Handler,
		},
		{
			MethodName: "GetAuditQueryResult",
			Handler:    _SecReportsService_GetAuditQueryResult_Handler,
		},
		{
			MethodName: "RunReport",
			Handler:    _SecReportsService_RunReport_Handler,
		},
		{
			MethodName: "GetReportResult",
			Handler:    _SecReportsService_GetReportResult_Handler,
		},
		{
			MethodName: "GetReportState",
			Handler:    _SecReportsService_GetReportState_Handler,
		},
		{
			MethodName: "ListReportStates",
			Handler:    _SecReportsService_ListReportStates_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _SecReportsService_GetSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/secreports/v1/secreports_service.proto",
}
