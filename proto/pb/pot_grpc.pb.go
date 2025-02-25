// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.0
// source: pot.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PoTExecutor_GetTxs_FullMethodName             = "/pb.PoTExecutor/GetTxs"
	PoTExecutor_VerifyTxs_FullMethodName          = "/pb.PoTExecutor/VerifyTxs"
	PoTExecutor_ExecuteTxs_FullMethodName         = "/pb.PoTExecutor/ExecuteTxs"
	PoTExecutor_VerifyIncensentive_FullMethodName = "/pb.PoTExecutor/VerifyIncensentive"
	PoTExecutor_GetIncentive_FullMethodName       = "/pb.PoTExecutor/GetIncentive"
)

// PoTExecutorClient is the client API for PoTExecutor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoTExecutorClient interface {
	GetTxs(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error)
	VerifyTxs(ctx context.Context, in *VerifyTxRequest, opts ...grpc.CallOption) (*VerifyTxResponse, error)
	ExecuteTxs(ctx context.Context, in *ExecuteTxRequest, opts ...grpc.CallOption) (*ExecuteTxResponse, error)
	VerifyIncensentive(ctx context.Context, in *IncensentiveVerifyRequest, opts ...grpc.CallOption) (*IncensentiveVerifyResponse, error)
	GetIncentive(ctx context.Context, in *GetIncentiveRequest, opts ...grpc.CallOption) (*GetIncentiveResponse, error)
}

type poTExecutorClient struct {
	cc grpc.ClientConnInterface
}

func NewPoTExecutorClient(cc grpc.ClientConnInterface) PoTExecutorClient {
	return &poTExecutorClient{cc}
}

func (c *poTExecutorClient) GetTxs(ctx context.Context, in *GetTxRequest, opts ...grpc.CallOption) (*GetTxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTxResponse)
	err := c.cc.Invoke(ctx, PoTExecutor_GetTxs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poTExecutorClient) VerifyTxs(ctx context.Context, in *VerifyTxRequest, opts ...grpc.CallOption) (*VerifyTxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyTxResponse)
	err := c.cc.Invoke(ctx, PoTExecutor_VerifyTxs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poTExecutorClient) ExecuteTxs(ctx context.Context, in *ExecuteTxRequest, opts ...grpc.CallOption) (*ExecuteTxResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteTxResponse)
	err := c.cc.Invoke(ctx, PoTExecutor_ExecuteTxs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poTExecutorClient) VerifyIncensentive(ctx context.Context, in *IncensentiveVerifyRequest, opts ...grpc.CallOption) (*IncensentiveVerifyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncensentiveVerifyResponse)
	err := c.cc.Invoke(ctx, PoTExecutor_VerifyIncensentive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poTExecutorClient) GetIncentive(ctx context.Context, in *GetIncentiveRequest, opts ...grpc.CallOption) (*GetIncentiveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIncentiveResponse)
	err := c.cc.Invoke(ctx, PoTExecutor_GetIncentive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoTExecutorServer is the server API for PoTExecutor service.
// All implementations should embed UnimplementedPoTExecutorServer
// for forward compatibility.
type PoTExecutorServer interface {
	GetTxs(context.Context, *GetTxRequest) (*GetTxResponse, error)
	VerifyTxs(context.Context, *VerifyTxRequest) (*VerifyTxResponse, error)
	ExecuteTxs(context.Context, *ExecuteTxRequest) (*ExecuteTxResponse, error)
	VerifyIncensentive(context.Context, *IncensentiveVerifyRequest) (*IncensentiveVerifyResponse, error)
	GetIncentive(context.Context, *GetIncentiveRequest) (*GetIncentiveResponse, error)
}

// UnimplementedPoTExecutorServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPoTExecutorServer struct{}

func (UnimplementedPoTExecutorServer) GetTxs(context.Context, *GetTxRequest) (*GetTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxs not implemented")
}
func (UnimplementedPoTExecutorServer) VerifyTxs(context.Context, *VerifyTxRequest) (*VerifyTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTxs not implemented")
}
func (UnimplementedPoTExecutorServer) ExecuteTxs(context.Context, *ExecuteTxRequest) (*ExecuteTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteTxs not implemented")
}
func (UnimplementedPoTExecutorServer) VerifyIncensentive(context.Context, *IncensentiveVerifyRequest) (*IncensentiveVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyIncensentive not implemented")
}
func (UnimplementedPoTExecutorServer) GetIncentive(context.Context, *GetIncentiveRequest) (*GetIncentiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncentive not implemented")
}
func (UnimplementedPoTExecutorServer) testEmbeddedByValue() {}

// UnsafePoTExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoTExecutorServer will
// result in compilation errors.
type UnsafePoTExecutorServer interface {
	mustEmbedUnimplementedPoTExecutorServer()
}

func RegisterPoTExecutorServer(s grpc.ServiceRegistrar, srv PoTExecutorServer) {
	// If the following call pancis, it indicates UnimplementedPoTExecutorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PoTExecutor_ServiceDesc, srv)
}

func _PoTExecutor_GetTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoTExecutorServer).GetTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoTExecutor_GetTxs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoTExecutorServer).GetTxs(ctx, req.(*GetTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoTExecutor_VerifyTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoTExecutorServer).VerifyTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoTExecutor_VerifyTxs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoTExecutorServer).VerifyTxs(ctx, req.(*VerifyTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoTExecutor_ExecuteTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoTExecutorServer).ExecuteTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoTExecutor_ExecuteTxs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoTExecutorServer).ExecuteTxs(ctx, req.(*ExecuteTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoTExecutor_VerifyIncensentive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncensentiveVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoTExecutorServer).VerifyIncensentive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoTExecutor_VerifyIncensentive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoTExecutorServer).VerifyIncensentive(ctx, req.(*IncensentiveVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoTExecutor_GetIncentive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncentiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoTExecutorServer).GetIncentive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoTExecutor_GetIncentive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoTExecutorServer).GetIncentive(ctx, req.(*GetIncentiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoTExecutor_ServiceDesc is the grpc.ServiceDesc for PoTExecutor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoTExecutor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PoTExecutor",
	HandlerType: (*PoTExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxs",
			Handler:    _PoTExecutor_GetTxs_Handler,
		},
		{
			MethodName: "VerifyTxs",
			Handler:    _PoTExecutor_VerifyTxs_Handler,
		},
		{
			MethodName: "ExecuteTxs",
			Handler:    _PoTExecutor_ExecuteTxs_Handler,
		},
		{
			MethodName: "VerifyIncensentive",
			Handler:    _PoTExecutor_VerifyIncensentive_Handler,
		},
		{
			MethodName: "GetIncentive",
			Handler:    _PoTExecutor_GetIncentive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pot.proto",
}

const (
	BciExector_SendBci_FullMethodName                          = "/pb.BciExector/SendBci"
	BciExector_GetBalance_FullMethodName                       = "/pb.BciExector/GetBalance"
	BciExector_VerifyUTXO_FullMethodName                       = "/pb.BciExector/VerifyUTXO"
	BciExector_CreateLockTransaction_FullMethodName            = "/pb.BciExector/CreateLockTransaction"
	BciExector_CreateLockTransferTransaction_FullMethodName    = "/pb.BciExector/CreateLockTransferTransaction"
	BciExector_CreateDevastateTransaction_FullMethodName       = "/pb.BciExector/CreateDevastateTransaction"
	BciExector_CreateNonLockTransferTransaction_FullMethodName = "/pb.BciExector/CreateNonLockTransferTransaction"
	BciExector_CreateBciToVsiTransaction_FullMethodName        = "/pb.BciExector/CreateBciToVsiTransaction"
	BciExector_GetPqcKey_FullMethodName                        = "/pb.BciExector/GetPqcKey"
)

// BciExectorClient is the client API for BciExector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BciExectorClient interface {
	SendBci(ctx context.Context, in *SendBciRequest, opts ...grpc.CallOption) (*SendBciResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	VerifyUTXO(ctx context.Context, in *VerifyUTXORequest, opts ...grpc.CallOption) (*VerifyUTXOResponse, error)
	CreateLockTransaction(ctx context.Context, in *CreateLockTransactionRequest, opts ...grpc.CallOption) (*CreateLockTransactionResponse, error)
	CreateLockTransferTransaction(ctx context.Context, in *CreateLockTransferTransactionRequest, opts ...grpc.CallOption) (*CreateLockTransferTransactionResponse, error)
	CreateDevastateTransaction(ctx context.Context, in *CreateDevastateTransactionRequest, opts ...grpc.CallOption) (*CreateDevastateTransactionResponse, error)
	CreateNonLockTransferTransaction(ctx context.Context, in *CreateNonLockTransferTransactionRequest, opts ...grpc.CallOption) (*CreateNonLockTransferTransactionResponse, error)
	CreateBciToVsiTransaction(ctx context.Context, in *CreateBciToVsiRequest, opts ...grpc.CallOption) (*CreateBciToVsiResponse, error)
	GetPqcKey(ctx context.Context, in *GetPqcKeyRequest, opts ...grpc.CallOption) (*GetPqcKeyResponse, error)
}

type bciExectorClient struct {
	cc grpc.ClientConnInterface
}

func NewBciExectorClient(cc grpc.ClientConnInterface) BciExectorClient {
	return &bciExectorClient{cc}
}

func (c *bciExectorClient) SendBci(ctx context.Context, in *SendBciRequest, opts ...grpc.CallOption) (*SendBciResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendBciResponse)
	err := c.cc.Invoke(ctx, BciExector_SendBci_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, BciExector_GetBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) VerifyUTXO(ctx context.Context, in *VerifyUTXORequest, opts ...grpc.CallOption) (*VerifyUTXOResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyUTXOResponse)
	err := c.cc.Invoke(ctx, BciExector_VerifyUTXO_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) CreateLockTransaction(ctx context.Context, in *CreateLockTransactionRequest, opts ...grpc.CallOption) (*CreateLockTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLockTransactionResponse)
	err := c.cc.Invoke(ctx, BciExector_CreateLockTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) CreateLockTransferTransaction(ctx context.Context, in *CreateLockTransferTransactionRequest, opts ...grpc.CallOption) (*CreateLockTransferTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLockTransferTransactionResponse)
	err := c.cc.Invoke(ctx, BciExector_CreateLockTransferTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) CreateDevastateTransaction(ctx context.Context, in *CreateDevastateTransactionRequest, opts ...grpc.CallOption) (*CreateDevastateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDevastateTransactionResponse)
	err := c.cc.Invoke(ctx, BciExector_CreateDevastateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) CreateNonLockTransferTransaction(ctx context.Context, in *CreateNonLockTransferTransactionRequest, opts ...grpc.CallOption) (*CreateNonLockTransferTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNonLockTransferTransactionResponse)
	err := c.cc.Invoke(ctx, BciExector_CreateNonLockTransferTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) CreateBciToVsiTransaction(ctx context.Context, in *CreateBciToVsiRequest, opts ...grpc.CallOption) (*CreateBciToVsiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBciToVsiResponse)
	err := c.cc.Invoke(ctx, BciExector_CreateBciToVsiTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bciExectorClient) GetPqcKey(ctx context.Context, in *GetPqcKeyRequest, opts ...grpc.CallOption) (*GetPqcKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPqcKeyResponse)
	err := c.cc.Invoke(ctx, BciExector_GetPqcKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BciExectorServer is the server API for BciExector service.
// All implementations should embed UnimplementedBciExectorServer
// for forward compatibility.
type BciExectorServer interface {
	SendBci(context.Context, *SendBciRequest) (*SendBciResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	VerifyUTXO(context.Context, *VerifyUTXORequest) (*VerifyUTXOResponse, error)
	CreateLockTransaction(context.Context, *CreateLockTransactionRequest) (*CreateLockTransactionResponse, error)
	CreateLockTransferTransaction(context.Context, *CreateLockTransferTransactionRequest) (*CreateLockTransferTransactionResponse, error)
	CreateDevastateTransaction(context.Context, *CreateDevastateTransactionRequest) (*CreateDevastateTransactionResponse, error)
	CreateNonLockTransferTransaction(context.Context, *CreateNonLockTransferTransactionRequest) (*CreateNonLockTransferTransactionResponse, error)
	CreateBciToVsiTransaction(context.Context, *CreateBciToVsiRequest) (*CreateBciToVsiResponse, error)
	GetPqcKey(context.Context, *GetPqcKeyRequest) (*GetPqcKeyResponse, error)
}

// UnimplementedBciExectorServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBciExectorServer struct{}

func (UnimplementedBciExectorServer) SendBci(context.Context, *SendBciRequest) (*SendBciResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBci not implemented")
}
func (UnimplementedBciExectorServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBciExectorServer) VerifyUTXO(context.Context, *VerifyUTXORequest) (*VerifyUTXOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUTXO not implemented")
}
func (UnimplementedBciExectorServer) CreateLockTransaction(context.Context, *CreateLockTransactionRequest) (*CreateLockTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLockTransaction not implemented")
}
func (UnimplementedBciExectorServer) CreateLockTransferTransaction(context.Context, *CreateLockTransferTransactionRequest) (*CreateLockTransferTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLockTransferTransaction not implemented")
}
func (UnimplementedBciExectorServer) CreateDevastateTransaction(context.Context, *CreateDevastateTransactionRequest) (*CreateDevastateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevastateTransaction not implemented")
}
func (UnimplementedBciExectorServer) CreateNonLockTransferTransaction(context.Context, *CreateNonLockTransferTransactionRequest) (*CreateNonLockTransferTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNonLockTransferTransaction not implemented")
}
func (UnimplementedBciExectorServer) CreateBciToVsiTransaction(context.Context, *CreateBciToVsiRequest) (*CreateBciToVsiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBciToVsiTransaction not implemented")
}
func (UnimplementedBciExectorServer) GetPqcKey(context.Context, *GetPqcKeyRequest) (*GetPqcKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPqcKey not implemented")
}
func (UnimplementedBciExectorServer) testEmbeddedByValue() {}

// UnsafeBciExectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BciExectorServer will
// result in compilation errors.
type UnsafeBciExectorServer interface {
	mustEmbedUnimplementedBciExectorServer()
}

func RegisterBciExectorServer(s grpc.ServiceRegistrar, srv BciExectorServer) {
	// If the following call pancis, it indicates UnimplementedBciExectorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BciExector_ServiceDesc, srv)
}

func _BciExector_SendBci_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBciRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).SendBci(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_SendBci_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).SendBci(ctx, req.(*SendBciRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_VerifyUTXO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUTXORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).VerifyUTXO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_VerifyUTXO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).VerifyUTXO(ctx, req.(*VerifyUTXORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_CreateLockTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLockTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).CreateLockTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_CreateLockTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).CreateLockTransaction(ctx, req.(*CreateLockTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_CreateLockTransferTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLockTransferTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).CreateLockTransferTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_CreateLockTransferTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).CreateLockTransferTransaction(ctx, req.(*CreateLockTransferTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_CreateDevastateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDevastateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).CreateDevastateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_CreateDevastateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).CreateDevastateTransaction(ctx, req.(*CreateDevastateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_CreateNonLockTransferTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNonLockTransferTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).CreateNonLockTransferTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_CreateNonLockTransferTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).CreateNonLockTransferTransaction(ctx, req.(*CreateNonLockTransferTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_CreateBciToVsiTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBciToVsiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).CreateBciToVsiTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_CreateBciToVsiTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).CreateBciToVsiTransaction(ctx, req.(*CreateBciToVsiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BciExector_GetPqcKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPqcKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BciExectorServer).GetPqcKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BciExector_GetPqcKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BciExectorServer).GetPqcKey(ctx, req.(*GetPqcKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BciExector_ServiceDesc is the grpc.ServiceDesc for BciExector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BciExector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BciExector",
	HandlerType: (*BciExectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBci",
			Handler:    _BciExector_SendBci_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _BciExector_GetBalance_Handler,
		},
		{
			MethodName: "VerifyUTXO",
			Handler:    _BciExector_VerifyUTXO_Handler,
		},
		{
			MethodName: "CreateLockTransaction",
			Handler:    _BciExector_CreateLockTransaction_Handler,
		},
		{
			MethodName: "CreateLockTransferTransaction",
			Handler:    _BciExector_CreateLockTransferTransaction_Handler,
		},
		{
			MethodName: "CreateDevastateTransaction",
			Handler:    _BciExector_CreateDevastateTransaction_Handler,
		},
		{
			MethodName: "CreateNonLockTransferTransaction",
			Handler:    _BciExector_CreateNonLockTransferTransaction_Handler,
		},
		{
			MethodName: "CreateBciToVsiTransaction",
			Handler:    _BciExector_CreateBciToVsiTransaction_Handler,
		},
		{
			MethodName: "GetPqcKey",
			Handler:    _BciExector_GetPqcKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pot.proto",
}
