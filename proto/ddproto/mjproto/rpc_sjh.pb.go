// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_sjh.proto

package mjproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ServerInfo from mahjong_hall.proto

// Ignoring public import of game_QuickConn from mahjong_hall.proto

// Ignoring public import of game_AckQuickConn from mahjong_hall.proto

// Ignoring public import of game_Login from mahjong_hall.proto

// Ignoring public import of game_AckLogin from mahjong_hall.proto

// Ignoring public import of game_Notice from mahjong_hall.proto

// Ignoring public import of game_AckNotice from mahjong_hall.proto

// Ignoring public import of game_GameRecord from mahjong_hall.proto

// Ignoring public import of BeanUserRecord from mahjong_hall.proto

// Ignoring public import of BeanGameRecord from mahjong_hall.proto

// Ignoring public import of game_AckGameRecord from mahjong_hall.proto

// Ignoring public import of game_Feedback from mahjong_hall.proto

// Ignoring public import of game_CreateRoom from mahjong_hall.proto

// Ignoring public import of game_AckCreateRoom from mahjong_hall.proto

// Ignoring public import of game_EnterRoom from mahjong_hall.proto

// Ignoring public import of game_AckEnterRoom from mahjong_hall.proto

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SjhMjRpc service

type SjhMjRpcClient interface {
	// 创建房间
	CreateRoom(ctx context.Context, in *Game_CreateRoom, opts ...grpc.CallOption) (*Game_AckCreateRoom, error)
}

type sjhMjRpcClient struct {
	cc *grpc.ClientConn
}

func NewSjhMjRpcClient(cc *grpc.ClientConn) SjhMjRpcClient {
	return &sjhMjRpcClient{cc}
}

func (c *sjhMjRpcClient) CreateRoom(ctx context.Context, in *Game_CreateRoom, opts ...grpc.CallOption) (*Game_AckCreateRoom, error) {
	out := new(Game_AckCreateRoom)
	err := grpc.Invoke(ctx, "/mjproto.SjhMjRpc/CreateRoom", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SjhMjRpc service

type SjhMjRpcServer interface {
	// 创建房间
	CreateRoom(context.Context, *Game_CreateRoom) (*Game_AckCreateRoom, error)
}

func RegisterSjhMjRpcServer(s *grpc.Server, srv SjhMjRpcServer) {
	s.RegisterService(&_SjhMjRpc_serviceDesc, srv)
}

func _SjhMjRpc_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game_CreateRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SjhMjRpcServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mjproto.SjhMjRpc/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SjhMjRpcServer).CreateRoom(ctx, req.(*Game_CreateRoom))
	}
	return interceptor(ctx, in, info, handler)
}

var _SjhMjRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mjproto.SjhMjRpc",
	HandlerType: (*SjhMjRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _SjhMjRpc_CreateRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_sjh.proto",
}

func init() { proto.RegisterFile("rpc_sjh.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2a, 0x48, 0x8e,
	0x2f, 0xce, 0xca, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xcd, 0x02, 0x33, 0xa4,
	0x84, 0x72, 0x13, 0x33, 0xb2, 0xf2, 0xf3, 0xd2, 0xe3, 0x33, 0x12, 0x73, 0x72, 0x20, 0x92, 0x46,
	0x81, 0x5c, 0x1c, 0xc1, 0x59, 0x19, 0xbe, 0x59, 0x41, 0x05, 0xc9, 0x42, 0xae, 0x5c, 0x5c, 0xce,
	0x45, 0xa9, 0x89, 0x25, 0xa9, 0x41, 0xf9, 0xf9, 0xb9, 0x42, 0x12, 0x7a, 0x50, 0x7d, 0x7a, 0xe9,
	0x89, 0xb9, 0xa9, 0xf1, 0x08, 0x19, 0x29, 0x69, 0x54, 0x19, 0xc7, 0xe4, 0x6c, 0x84, 0xa4, 0x12,
	0x43, 0x00, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x97, 0xbd, 0x75, 0x81, 0x00, 0x00, 0x00,
}
