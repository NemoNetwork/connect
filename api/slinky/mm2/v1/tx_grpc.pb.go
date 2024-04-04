// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: slinky/mm2/v1/tx.proto

package mm2v1

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
	Msg_CreateMarkets_FullMethodName           = "/slinky.mm2.v1.Msg/CreateMarkets"
	Msg_UpdateMarkets_FullMethodName           = "/slinky.mm2.v1.Msg/UpdateMarkets"
	Msg_Params_FullMethodName                  = "/slinky.mm2.v1.Msg/Params"
	Msg_RemoveMarketAuthorities_FullMethodName = "/slinky.mm2.v1.Msg/RemoveMarketAuthorities"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateMarkets creates markets from the given message.
	CreateMarkets(ctx context.Context, in *MsgCreateMarkets, opts ...grpc.CallOption) (*MsgCreateMarketsResponse, error)
	// UpdateMarkets updates markets from the given message.
	UpdateMarkets(ctx context.Context, in *MsgUpdateMarkets, opts ...grpc.CallOption) (*MsgUpdateMarketsResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
	// RemoveMarketAuthorities defines a method for removing market authorities
	// from the x/marketmap module. the signer must be the admin.
	RemoveMarketAuthorities(ctx context.Context, in *MsgRemoveMarketAuthorities, opts ...grpc.CallOption) (*MsgRemoveMarketAuthoritiesResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMarkets(ctx context.Context, in *MsgCreateMarkets, opts ...grpc.CallOption) (*MsgCreateMarketsResponse, error) {
	out := new(MsgCreateMarketsResponse)
	err := c.cc.Invoke(ctx, Msg_CreateMarkets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateMarkets(ctx context.Context, in *MsgUpdateMarkets, opts ...grpc.CallOption) (*MsgUpdateMarketsResponse, error) {
	out := new(MsgUpdateMarketsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateMarkets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, Msg_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveMarketAuthorities(ctx context.Context, in *MsgRemoveMarketAuthorities, opts ...grpc.CallOption) (*MsgRemoveMarketAuthoritiesResponse, error) {
	out := new(MsgRemoveMarketAuthoritiesResponse)
	err := c.cc.Invoke(ctx, Msg_RemoveMarketAuthorities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateMarkets creates markets from the given message.
	CreateMarkets(context.Context, *MsgCreateMarkets) (*MsgCreateMarketsResponse, error)
	// UpdateMarkets updates markets from the given message.
	UpdateMarkets(context.Context, *MsgUpdateMarkets) (*MsgUpdateMarketsResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
	// RemoveMarketAuthorities defines a method for removing market authorities
	// from the x/marketmap module. the signer must be the admin.
	RemoveMarketAuthorities(context.Context, *MsgRemoveMarketAuthorities) (*MsgRemoveMarketAuthoritiesResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateMarkets(context.Context, *MsgCreateMarkets) (*MsgCreateMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarkets not implemented")
}
func (UnimplementedMsgServer) UpdateMarkets(context.Context, *MsgUpdateMarkets) (*MsgUpdateMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarkets not implemented")
}
func (UnimplementedMsgServer) Params(context.Context, *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedMsgServer) RemoveMarketAuthorities(context.Context, *MsgRemoveMarketAuthorities) (*MsgRemoveMarketAuthoritiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMarketAuthorities not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMarkets)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateMarkets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMarkets(ctx, req.(*MsgCreateMarkets))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMarkets)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateMarkets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMarkets(ctx, req.(*MsgUpdateMarkets))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveMarketAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveMarketAuthorities)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveMarketAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveMarketAuthorities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveMarketAuthorities(ctx, req.(*MsgRemoveMarketAuthorities))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.mm2.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMarkets",
			Handler:    _Msg_CreateMarkets_Handler,
		},
		{
			MethodName: "UpdateMarkets",
			Handler:    _Msg_UpdateMarkets_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
		{
			MethodName: "RemoveMarketAuthorities",
			Handler:    _Msg_RemoveMarketAuthorities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/mm2/v1/tx.proto",
}
