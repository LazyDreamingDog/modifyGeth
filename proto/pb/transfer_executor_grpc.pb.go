// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.0
// source: transfer_executor.proto

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
	TransferExecutorGRPC_CommitWithdrawTx_FullMethodName = "/pb.TransferExecutorGRPC/CommitWithdrawTx"
	TransferExecutorGRPC_VerifyChangeTx_FullMethodName   = "/pb.TransferExecutorGRPC/VerifyChangeTx"
)

// TransferExecutorGRPCClient is the client API for TransferExecutorGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferExecutorGRPCClient interface {
	CommitWithdrawTx(ctx context.Context, in *CommitWithdrawTxRequest, opts ...grpc.CallOption) (*Empty, error)
	VerifyChangeTx(ctx context.Context, in *VerifyChangeTxRequest, opts ...grpc.CallOption) (*VerifyChangeTxReply, error)
}

type transferExecutorGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferExecutorGRPCClient(cc grpc.ClientConnInterface) TransferExecutorGRPCClient {
	return &transferExecutorGRPCClient{cc}
}

func (c *transferExecutorGRPCClient) CommitWithdrawTx(ctx context.Context, in *CommitWithdrawTxRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, TransferExecutorGRPC_CommitWithdrawTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferExecutorGRPCClient) VerifyChangeTx(ctx context.Context, in *VerifyChangeTxRequest, opts ...grpc.CallOption) (*VerifyChangeTxReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyChangeTxReply)
	err := c.cc.Invoke(ctx, TransferExecutorGRPC_VerifyChangeTx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferExecutorGRPCServer is the server API for TransferExecutorGRPC service.
// All implementations must embed UnimplementedTransferExecutorGRPCServer
// for forward compatibility.
type TransferExecutorGRPCServer interface {
	CommitWithdrawTx(context.Context, *CommitWithdrawTxRequest) (*Empty, error)
	VerifyChangeTx(context.Context, *VerifyChangeTxRequest) (*VerifyChangeTxReply, error)
	mustEmbedUnimplementedTransferExecutorGRPCServer()
}

// UnimplementedTransferExecutorGRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferExecutorGRPCServer struct{}

func (UnimplementedTransferExecutorGRPCServer) CommitWithdrawTx(context.Context, *CommitWithdrawTxRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitWithdrawTx not implemented")
}
func (UnimplementedTransferExecutorGRPCServer) VerifyChangeTx(context.Context, *VerifyChangeTxRequest) (*VerifyChangeTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyChangeTx not implemented")
}
func (UnimplementedTransferExecutorGRPCServer) mustEmbedUnimplementedTransferExecutorGRPCServer() {}
func (UnimplementedTransferExecutorGRPCServer) testEmbeddedByValue()                              {}

// UnsafeTransferExecutorGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferExecutorGRPCServer will
// result in compilation errors.
type UnsafeTransferExecutorGRPCServer interface {
	mustEmbedUnimplementedTransferExecutorGRPCServer()
}

func RegisterTransferExecutorGRPCServer(s grpc.ServiceRegistrar, srv TransferExecutorGRPCServer) {
	// If the following call pancis, it indicates UnimplementedTransferExecutorGRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransferExecutorGRPC_ServiceDesc, srv)
}

func _TransferExecutorGRPC_CommitWithdrawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitWithdrawTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferExecutorGRPCServer).CommitWithdrawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferExecutorGRPC_CommitWithdrawTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferExecutorGRPCServer).CommitWithdrawTx(ctx, req.(*CommitWithdrawTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferExecutorGRPC_VerifyChangeTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyChangeTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferExecutorGRPCServer).VerifyChangeTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferExecutorGRPC_VerifyChangeTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferExecutorGRPCServer).VerifyChangeTx(ctx, req.(*VerifyChangeTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferExecutorGRPC_ServiceDesc is the grpc.ServiceDesc for TransferExecutorGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferExecutorGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TransferExecutorGRPC",
	HandlerType: (*TransferExecutorGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommitWithdrawTx",
			Handler:    _TransferExecutorGRPC_CommitWithdrawTx_Handler,
		},
		{
			MethodName: "VerifyChangeTx",
			Handler:    _TransferExecutorGRPC_VerifyChangeTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer_executor.proto",
}
