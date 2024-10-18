// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: coupon/coupon_service.proto

package coupon_service

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
	CouponService_Echo_FullMethodName = "/coupon_service.CouponService/Echo"
)

// CouponServiceClient is the client API for CouponService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponServiceClient interface {
	Echo(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type couponServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponServiceClient(cc grpc.ClientConnInterface) CouponServiceClient {
	return &couponServiceClient{cc}
}

func (c *couponServiceClient) Echo(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, CouponService_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServiceServer is the server API for CouponService service.
// All implementations must embed UnimplementedCouponServiceServer
// for forward compatibility
type CouponServiceServer interface {
	Echo(context.Context, *CommonRequest) (*CommonResponse, error)
	mustEmbedUnimplementedCouponServiceServer()
}

// UnimplementedCouponServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServiceServer struct {
}

func (UnimplementedCouponServiceServer) Echo(context.Context, *CommonRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedCouponServiceServer) mustEmbedUnimplementedCouponServiceServer() {}

// UnsafeCouponServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServiceServer will
// result in compilation errors.
type UnsafeCouponServiceServer interface {
	mustEmbedUnimplementedCouponServiceServer()
}

func RegisterCouponServiceServer(s grpc.ServiceRegistrar, srv CouponServiceServer) {
	s.RegisterService(&CouponService_ServiceDesc, srv)
}

func _CouponService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CouponService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).Echo(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CouponService_ServiceDesc is the grpc.ServiceDesc for CouponService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CouponService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupon_service.CouponService",
	HandlerType: (*CouponServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _CouponService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coupon/coupon_service.proto",
}
