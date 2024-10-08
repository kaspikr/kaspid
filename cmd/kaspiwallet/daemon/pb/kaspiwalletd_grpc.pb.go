// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: kaspiwalletd.proto

package pb

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

// KaspiwalletdClient is the client API for Kaspiwalletd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KaspiwalletdClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type kaspiwalletdClient struct {
	cc grpc.ClientConnInterface
}

func NewKaspiwalletdClient(cc grpc.ClientConnInterface) KaspiwalletdClient {
	return &kaspiwalletdClient{cc}
}

func (c *kaspiwalletdClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) GetExternalSpendableUTXOs(ctx context.Context, in *GetExternalSpendableUTXOsRequest, opts ...grpc.CallOption) (*GetExternalSpendableUTXOsResponse, error) {
	out := new(GetExternalSpendableUTXOsResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/GetExternalSpendableUTXOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) CreateUnsignedTransactions(ctx context.Context, in *CreateUnsignedTransactionsRequest, opts ...grpc.CallOption) (*CreateUnsignedTransactionsResponse, error) {
	out := new(CreateUnsignedTransactionsResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/CreateUnsignedTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) ShowAddresses(ctx context.Context, in *ShowAddressesRequest, opts ...grpc.CallOption) (*ShowAddressesResponse, error) {
	out := new(ShowAddressesResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/ShowAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/NewAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kaspiwalletdClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/kaspiwalletd.kaspiwalletd/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KaspiwalletdServer is the server API for Kaspiwalletd service.
// All implementations must embed UnimplementedKaspiwalletdServer
// for forward compatibility
type KaspiwalletdServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error)
	CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error)
	ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	// Since SendRequest contains a password - this command should only be used on a trusted or secure connection
	Send(context.Context, *SendRequest) (*SendResponse, error)
	// Since SignRequest contains a password - this command should only be used on a trusted or secure connection
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedKaspiwalletdServer()
}

// UnimplementedKaspiwalletdServer must be embedded to have forward compatible implementations.
type UnimplementedKaspiwalletdServer struct {
}

func (UnimplementedKaspiwalletdServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedKaspiwalletdServer) GetExternalSpendableUTXOs(context.Context, *GetExternalSpendableUTXOsRequest) (*GetExternalSpendableUTXOsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExternalSpendableUTXOs not implemented")
}
func (UnimplementedKaspiwalletdServer) CreateUnsignedTransactions(context.Context, *CreateUnsignedTransactionsRequest) (*CreateUnsignedTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnsignedTransactions not implemented")
}
func (UnimplementedKaspiwalletdServer) ShowAddresses(context.Context, *ShowAddressesRequest) (*ShowAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAddresses not implemented")
}
func (UnimplementedKaspiwalletdServer) NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAddress not implemented")
}
func (UnimplementedKaspiwalletdServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedKaspiwalletdServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedKaspiwalletdServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedKaspiwalletdServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedKaspiwalletdServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedKaspiwalletdServer) mustEmbedUnimplementedKaspiwalletdServer() {}

// UnsafeKaspiwalletdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KaspiwalletdServer will
// result in compilation errors.
type UnsafeKaspiwalletdServer interface {
	mustEmbedUnimplementedKaspiwalletdServer()
}

func RegisterKaspiwalletdServer(s grpc.ServiceRegistrar, srv KaspiwalletdServer) {
	s.RegisterService(&Kaspiwalletd_ServiceDesc, srv)
}

func _Kaspiwalletd_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_GetExternalSpendableUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExternalSpendableUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).GetExternalSpendableUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/GetExternalSpendableUTXOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).GetExternalSpendableUTXOs(ctx, req.(*GetExternalSpendableUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_CreateUnsignedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUnsignedTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).CreateUnsignedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/CreateUnsignedTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).CreateUnsignedTransactions(ctx, req.(*CreateUnsignedTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_ShowAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).ShowAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/ShowAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).ShowAddresses(ctx, req.(*ShowAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kaspiwalletd_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KaspiwalletdServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaspiwalletd.kaspiwalletd/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KaspiwalletdServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kaspiwalletd_ServiceDesc is the grpc.ServiceDesc for Kaspiwalletd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kaspiwalletd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kaspiwalletd.kaspiwalletd",
	HandlerType: (*KaspiwalletdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Kaspiwalletd_GetBalance_Handler,
		},
		{
			MethodName: "GetExternalSpendableUTXOs",
			Handler:    _Kaspiwalletd_GetExternalSpendableUTXOs_Handler,
		},
		{
			MethodName: "CreateUnsignedTransactions",
			Handler:    _Kaspiwalletd_CreateUnsignedTransactions_Handler,
		},
		{
			MethodName: "ShowAddresses",
			Handler:    _Kaspiwalletd_ShowAddresses_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Kaspiwalletd_NewAddress_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Kaspiwalletd_Shutdown_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Kaspiwalletd_Broadcast_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Kaspiwalletd_Send_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Kaspiwalletd_Sign_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Kaspiwalletd_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kaspiwalletd.proto",
}
