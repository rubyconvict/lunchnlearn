// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/lunchnlearn/lunchnlearn.proto

package lunchnlearn

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

// LunchnlearnServiceClient is the client API for LunchnlearnService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LunchnlearnServiceClient interface {
	PlaceOrder(ctx context.Context, in *LunchRequest, opts ...grpc.CallOption) (*LunchResponse, error)
}

type lunchnlearnServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLunchnlearnServiceClient(cc grpc.ClientConnInterface) LunchnlearnServiceClient {
	return &lunchnlearnServiceClient{cc}
}

func (c *lunchnlearnServiceClient) PlaceOrder(ctx context.Context, in *LunchRequest, opts ...grpc.CallOption) (*LunchResponse, error) {
	out := new(LunchResponse)
	err := c.cc.Invoke(ctx, "/lunchnlearn.LunchnlearnService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LunchnlearnServiceServer is the server API for LunchnlearnService service.
// All implementations must embed UnimplementedLunchnlearnServiceServer
// for forward compatibility
type LunchnlearnServiceServer interface {
	PlaceOrder(context.Context, *LunchRequest) (*LunchResponse, error)
	mustEmbedUnimplementedLunchnlearnServiceServer()
}

// UnimplementedLunchnlearnServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLunchnlearnServiceServer struct {
}

func (UnimplementedLunchnlearnServiceServer) PlaceOrder(context.Context, *LunchRequest) (*LunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedLunchnlearnServiceServer) mustEmbedUnimplementedLunchnlearnServiceServer() {}

// UnsafeLunchnlearnServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LunchnlearnServiceServer will
// result in compilation errors.
type UnsafeLunchnlearnServiceServer interface {
	mustEmbedUnimplementedLunchnlearnServiceServer()
}

func RegisterLunchnlearnServiceServer(s grpc.ServiceRegistrar, srv LunchnlearnServiceServer) {
	s.RegisterService(&LunchnlearnService_ServiceDesc, srv)
}

func _LunchnlearnService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LunchnlearnServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lunchnlearn.LunchnlearnService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LunchnlearnServiceServer).PlaceOrder(ctx, req.(*LunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LunchnlearnService_ServiceDesc is the grpc.ServiceDesc for LunchnlearnService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LunchnlearnService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lunchnlearn.LunchnlearnService",
	HandlerType: (*LunchnlearnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _LunchnlearnService_PlaceOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lunchnlearn/lunchnlearn.proto",
}
