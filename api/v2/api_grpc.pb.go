// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/v2/api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Dex_GetClient_FullMethodName      = "/api.Dex/GetClient"
	Dex_CreateClient_FullMethodName   = "/api.Dex/CreateClient"
	Dex_UpdateClient_FullMethodName   = "/api.Dex/UpdateClient"
	Dex_DeleteClient_FullMethodName   = "/api.Dex/DeleteClient"
	Dex_CreatePassword_FullMethodName = "/api.Dex/CreatePassword"
	Dex_UpdatePassword_FullMethodName = "/api.Dex/UpdatePassword"
	Dex_DeletePassword_FullMethodName = "/api.Dex/DeletePassword"
	Dex_ListPasswords_FullMethodName  = "/api.Dex/ListPasswords"
	Dex_GetVersion_FullMethodName     = "/api.Dex/GetVersion"
	Dex_ListRefresh_FullMethodName    = "/api.Dex/ListRefresh"
	Dex_RevokeRefresh_FullMethodName  = "/api.Dex/RevokeRefresh"
	Dex_VerifyPassword_FullMethodName = "/api.Dex/VerifyPassword"
)

// DexClient is the client API for Dex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DexClient interface {
	// GetClient gets a client.
	GetClient(ctx context.Context, in *GetClientReq, opts ...grpc.CallOption) (*GetClientResp, error)
	// CreateClient creates a client.
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error)
	// UpdateClient updates an existing client
	UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error)
	// RevokeRefresh revokes the refresh token for the provided user-client pair.
	//
	// Note that each user-client pair can have only one refresh token at a time.
	RevokeRefresh(ctx context.Context, in *RevokeRefreshReq, opts ...grpc.CallOption) (*RevokeRefreshResp, error)
	// VerifyPassword returns whether a password matches a hash for a specific email or not.
	VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordResp, error)
	// Create Connector on the fly
	CreateConnector(ctx context.Context, in *Connector, opts ...grpc.CallOption) (*Connector, error)
	// Update Connector on the fly
	UpdateConnector(ctx context.Context, in *Connector, opts ...grpc.CallOption) (*Connector, error)
	// DeleteConnector deletes the corresponding connector
	DeleteConnector(ctx context.Context, in *DeleteConnectorReq, opts ...grpc.CallOption) (*DeleteConnectorResp, error)
	// ListConnector lists the available connectors
	ListConnector(ctx context.Context, in *ListConnectorReq, opts ...grpc.CallOption) (*ListConnectorResp, error)
}

type dexClient struct {
	cc grpc.ClientConnInterface
}

func NewDexClient(cc grpc.ClientConnInterface) DexClient {
	return &dexClient{cc}
}

func (c *dexClient) GetClient(ctx context.Context, in *GetClientReq, opts ...grpc.CallOption) (*GetClientResp, error) {
	out := new(GetClientResp)
	err := c.cc.Invoke(ctx, Dex_GetClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error) {
	out := new(CreateClientResp)
	err := c.cc.Invoke(ctx, Dex_CreateClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error) {
	out := new(UpdateClientResp)
	err := c.cc.Invoke(ctx, Dex_UpdateClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	out := new(DeleteClientResp)
	err := c.cc.Invoke(ctx, Dex_DeleteClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error) {
	out := new(CreatePasswordResp)
	err := c.cc.Invoke(ctx, Dex_CreatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := c.cc.Invoke(ctx, Dex_UpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error) {
	out := new(DeletePasswordResp)
	err := c.cc.Invoke(ctx, Dex_DeletePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error) {
	out := new(ListPasswordResp)
	err := c.cc.Invoke(ctx, Dex_ListPasswords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	out := new(VersionResp)
	err := c.cc.Invoke(ctx, Dex_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error) {
	out := new(ListRefreshResp)
	err := c.cc.Invoke(ctx, Dex_ListRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) RevokeRefresh(ctx context.Context, in *RevokeRefreshReq, opts ...grpc.CallOption) (*RevokeRefreshResp, error) {
	out := new(RevokeRefreshResp)
	err := c.cc.Invoke(ctx, Dex_RevokeRefresh_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordResp, error) {
	out := new(VerifyPasswordResp)
	err := c.cc.Invoke(ctx, Dex_VerifyPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreateConnector(ctx context.Context, in *Connector, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/api.Dex/CreateConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdateConnector(ctx context.Context, in *Connector, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/api.Dex/UpdateConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeleteConnector(ctx context.Context, in *DeleteConnectorReq, opts ...grpc.CallOption) (*DeleteConnectorResp, error) {
	out := new(DeleteConnectorResp)
	err := c.cc.Invoke(ctx, "/api.Dex/DeleteConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) ListConnector(ctx context.Context, in *ListConnectorReq, opts ...grpc.CallOption) (*ListConnectorResp, error) {
	out := new(ListConnectorResp)
	err := c.cc.Invoke(ctx, "/api.Dex/ListConnector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DexServer is the server API for Dex service.
// All implementations must embed UnimplementedDexServer
// for forward compatibility
type DexServer interface {
	// GetClient gets a client.
	GetClient(context.Context, *GetClientReq) (*GetClientResp, error)
	// CreateClient creates a client.
	CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error)
	// UpdateClient updates an existing client
	UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(context.Context, *ListPasswordReq) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(context.Context, *VersionReq) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(context.Context, *ListRefreshReq) (*ListRefreshResp, error)
	// RevokeRefresh revokes the refresh token for the provided user-client pair.
	//
	// Note that each user-client pair can have only one refresh token at a time.
	RevokeRefresh(context.Context, *RevokeRefreshReq) (*RevokeRefreshResp, error)
	// VerifyPassword returns whether a password matches a hash for a specific email or not.
	VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordResp, error)
	// Create Connector on the fly
	CreateConnector(context.Context, *Connector) (*Connector, error)
	// Update Connector on the fly
	UpdateConnector(context.Context, *Connector) (*Connector, error)
	// DeleteConnector deletes the corresponding connector
	DeleteConnector(context.Context, *DeleteConnectorReq) (*DeleteConnectorResp, error)
	// ListConnector lists the available connectors
	ListConnector(context.Context, *ListConnectorReq) (*ListConnectorResp, error)
	mustEmbedUnimplementedDexServer()
}

// UnimplementedDexServer must be embedded to have forward compatible implementations.
type UnimplementedDexServer struct {
}

func (UnimplementedDexServer) GetClient(context.Context, *GetClientReq) (*GetClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedDexServer) CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedDexServer) UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedDexServer) DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedDexServer) CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePassword not implemented")
}
func (UnimplementedDexServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedDexServer) DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassword not implemented")
}
func (UnimplementedDexServer) ListPasswords(context.Context, *ListPasswordReq) (*ListPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPasswords not implemented")
}
func (UnimplementedDexServer) GetVersion(context.Context, *VersionReq) (*VersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedDexServer) ListRefresh(context.Context, *ListRefreshReq) (*ListRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefresh not implemented")
}
func (UnimplementedDexServer) RevokeRefresh(context.Context, *RevokeRefreshReq) (*RevokeRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefresh not implemented")
}
func (UnimplementedDexServer) VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPassword not implemented")
}
func (UnimplementedDexServer) CreateConnector(context.Context, *Connector) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnector not implemented")
}
func (UnimplementedDexServer) UpdateConnector(context.Context, *Connector) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnector not implemented")
}
func (UnimplementedDexServer) DeleteConnector(context.Context, *DeleteConnectorReq) (*DeleteConnectorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnector not implemented")
}
func (UnimplementedDexServer) ListConnector(context.Context, *ListConnectorReq) (*ListConnectorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnector not implemented")
}
func (UnimplementedDexServer) mustEmbedUnimplementedDexServer() {}

// UnsafeDexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DexServer will
// result in compilation errors.
type UnsafeDexServer interface {
	mustEmbedUnimplementedDexServer()
}

func RegisterDexServer(s grpc.ServiceRegistrar, srv DexServer) {
	s.RegisterService(&Dex_ServiceDesc, srv)
}

func _Dex_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_GetClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).GetClient(ctx, req.(*GetClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_CreateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_UpdateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdateClient(ctx, req.(*UpdateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_DeleteClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_CreatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreatePassword(ctx, req.(*CreatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_DeletePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeletePassword(ctx, req.(*DeletePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_ListPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).ListPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_ListPasswords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).ListPasswords(ctx, req.(*ListPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).GetVersion(ctx, req.(*VersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_ListRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).ListRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_ListRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).ListRefresh(ctx, req.(*ListRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_RevokeRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).RevokeRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_RevokeRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).RevokeRefresh(ctx, req.(*RevokeRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_VerifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).VerifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dex_VerifyPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).VerifyPassword(ctx, req.(*VerifyPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreateConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreateConnector(ctx, req.(*Connector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/UpdateConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdateConnector(ctx, req.(*Connector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeleteConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeleteConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeleteConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeleteConnector(ctx, req.(*DeleteConnectorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_ListConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).ListConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/ListConnector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).ListConnector(ctx, req.(*ListConnectorReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Dex_ServiceDesc is the grpc.ServiceDesc for Dex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Dex",
	HandlerType: (*DexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClient",
			Handler:    _Dex_GetClient_Handler,
		},
		{
			MethodName: "CreateClient",
			Handler:    _Dex_CreateClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _Dex_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Dex_DeleteClient_Handler,
		},
		{
			MethodName: "CreatePassword",
			Handler:    _Dex_CreatePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Dex_UpdatePassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _Dex_DeletePassword_Handler,
		},
		{
			MethodName: "ListPasswords",
			Handler:    _Dex_ListPasswords_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Dex_GetVersion_Handler,
		},
		{
			MethodName: "ListRefresh",
			Handler:    _Dex_ListRefresh_Handler,
		},
		{
			MethodName: "RevokeRefresh",
			Handler:    _Dex_RevokeRefresh_Handler,
		},
		{
			MethodName: "VerifyPassword",
			Handler:    _Dex_VerifyPassword_Handler,
		},
		{
			MethodName: "CreateConnector",
			Handler:    _Dex_CreateConnector_Handler,
		},
		{
			MethodName: "UpdateConnector",
			Handler:    _Dex_UpdateConnector_Handler,
		},
		{
			MethodName: "DeleteConnector",
			Handler:    _Dex_DeleteConnector_Handler,
		},
		{
			MethodName: "ListConnector",
			Handler:    _Dex_ListConnector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/api.proto",
}
