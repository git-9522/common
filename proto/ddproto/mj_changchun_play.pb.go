// Code generated by protoc-gen-go.
// source: mj_changchun_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 长春麻将的gameId = 16
// 长春麻将 创建房间
type P16ReqCreateDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqCreateDesk) Reset()                    { *m = P16ReqCreateDesk{} }
func (m *P16ReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*P16ReqCreateDesk) ProtoMessage()               {}
func (*P16ReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *P16ReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckCreateDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckCreateDesk) Reset()                    { *m = P16AckCreateDesk{} }
func (m *P16AckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*P16AckCreateDesk) ProtoMessage()               {}
func (*P16AckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *P16AckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 进入房间
type P16ReqEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqEnterDesk) Reset()                    { *m = P16ReqEnterDesk{} }
func (m *P16ReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*P16ReqEnterDesk) ProtoMessage()               {}
func (*P16ReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{2} }

func (m *P16ReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckEnterDesk) Reset()                    { *m = P16AckEnterDesk{} }
func (m *P16AckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*P16AckEnterDesk) ProtoMessage()               {}
func (*P16AckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{3} }

func (m *P16AckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 准备
type P16ReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqReady) Reset()                    { *m = P16ReqReady{} }
func (m *P16ReqReady) String() string            { return proto.CompactTextString(m) }
func (*P16ReqReady) ProtoMessage()               {}
func (*P16ReqReady) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{4} }

func (m *P16ReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckReady) Reset()                    { *m = P16AckReady{} }
func (m *P16AckReady) String() string            { return proto.CompactTextString(m) }
func (*P16AckReady) ProtoMessage()               {}
func (*P16AckReady) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{5} }

func (m *P16AckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 开始
type P16ReqOpening struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqOpening) Reset()                    { *m = P16ReqOpening{} }
func (m *P16ReqOpening) String() string            { return proto.CompactTextString(m) }
func (*P16ReqOpening) ProtoMessage()               {}
func (*P16ReqOpening) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{6} }

func (m *P16ReqOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckOpening struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckOpening) Reset()                    { *m = P16AckOpening{} }
func (m *P16AckOpening) String() string            { return proto.CompactTextString(m) }
func (*P16AckOpening) ProtoMessage()               {}
func (*P16AckOpening) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{7} }

func (m *P16AckOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 动作广播
type P16BcOverTurn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16BcOverTurn) Reset()                    { *m = P16BcOverTurn{} }
func (m *P16BcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*P16BcOverTurn) ProtoMessage()               {}
func (*P16BcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{8} }

func (m *P16BcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 碰
type P16ReqPeng struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqPeng) Reset()                    { *m = P16ReqPeng{} }
func (m *P16ReqPeng) String() string            { return proto.CompactTextString(m) }
func (*P16ReqPeng) ProtoMessage()               {}
func (*P16ReqPeng) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{9} }

func (m *P16ReqPeng) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckPeng struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckPeng) Reset()                    { *m = P16AckPeng{} }
func (m *P16AckPeng) String() string            { return proto.CompactTextString(m) }
func (*P16AckPeng) ProtoMessage()               {}
func (*P16AckPeng) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{10} }

func (m *P16AckPeng) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 缸
type P16ReqGang struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqGang) Reset()                    { *m = P16ReqGang{} }
func (m *P16ReqGang) String() string            { return proto.CompactTextString(m) }
func (*P16ReqGang) ProtoMessage()               {}
func (*P16ReqGang) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{11} }

func (m *P16ReqGang) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckGang struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckGang) Reset()                    { *m = P16AckGang{} }
func (m *P16AckGang) String() string            { return proto.CompactTextString(m) }
func (*P16AckGang) ProtoMessage()               {}
func (*P16AckGang) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{12} }

func (m *P16AckGang) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 吃
type P16ReqChi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqChi) Reset()                    { *m = P16ReqChi{} }
func (m *P16ReqChi) String() string            { return proto.CompactTextString(m) }
func (*P16ReqChi) ProtoMessage()               {}
func (*P16ReqChi) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{13} }

func (m *P16ReqChi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckChi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckChi) Reset()                    { *m = P16AckChi{} }
func (m *P16AckChi) String() string            { return proto.CompactTextString(m) }
func (*P16AckChi) ProtoMessage()               {}
func (*P16AckChi) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{14} }

func (m *P16AckChi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 胡
type P16ReqHu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16ReqHu) Reset()                    { *m = P16ReqHu{} }
func (m *P16ReqHu) String() string            { return proto.CompactTextString(m) }
func (*P16ReqHu) ProtoMessage()               {}
func (*P16ReqHu) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{15} }

func (m *P16ReqHu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type P16AckHu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *P16AckHu) Reset()                    { *m = P16AckHu{} }
func (m *P16AckHu) String() string            { return proto.CompactTextString(m) }
func (*P16AckHu) ProtoMessage()               {}
func (*P16AckHu) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{16} }

func (m *P16AckHu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*P16ReqCreateDesk)(nil), "ddproto.p16_req_createDesk")
	proto.RegisterType((*P16AckCreateDesk)(nil), "ddproto.p16_ack_createDesk")
	proto.RegisterType((*P16ReqEnterDesk)(nil), "ddproto.p16_req_enterDesk")
	proto.RegisterType((*P16AckEnterDesk)(nil), "ddproto.p16_ack_enterDesk")
	proto.RegisterType((*P16ReqReady)(nil), "ddproto.p16_req_ready")
	proto.RegisterType((*P16AckReady)(nil), "ddproto.p16_ack_ready")
	proto.RegisterType((*P16ReqOpening)(nil), "ddproto.p16_req_opening")
	proto.RegisterType((*P16AckOpening)(nil), "ddproto.p16_ack_opening")
	proto.RegisterType((*P16BcOverTurn)(nil), "ddproto.p16_bc_overTurn")
	proto.RegisterType((*P16ReqPeng)(nil), "ddproto.p16_req_peng")
	proto.RegisterType((*P16AckPeng)(nil), "ddproto.p16_ack_peng")
	proto.RegisterType((*P16ReqGang)(nil), "ddproto.p16_req_gang")
	proto.RegisterType((*P16AckGang)(nil), "ddproto.p16_ack_gang")
	proto.RegisterType((*P16ReqChi)(nil), "ddproto.p16_req_chi")
	proto.RegisterType((*P16AckChi)(nil), "ddproto.p16_ack_chi")
	proto.RegisterType((*P16ReqHu)(nil), "ddproto.p16_req_hu")
	proto.RegisterType((*P16AckHu)(nil), "ddproto.p16_ack_hu")
}

var fileDescriptor30 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0xd2, 0x3f, 0x4b, 0xc4, 0x30,
	0x18, 0xc7, 0x71, 0xbb, 0x28, 0x3c, 0x55, 0xc4, 0x2a, 0x28, 0x4e, 0x22, 0x0e, 0x4e, 0x05, 0xeb,
	0x3f, 0x74, 0x76, 0x70, 0xec, 0xe0, 0x1e, 0x62, 0xf2, 0x90, 0xf4, 0xda, 0x26, 0xb9, 0x5c, 0x7b,
	0xd0, 0x77, 0x7f, 0x4d, 0xef, 0x72, 0x73, 0x9f, 0x2e, 0xa5, 0x10, 0x3e, 0xf9, 0xfe, 0x86, 0xc0,
	0x6d, 0xbb, 0x62, 0x42, 0x73, 0xa3, 0x84, 0xee, 0x0d, 0x73, 0x0d, 0x1f, 0x72, 0xe7, 0x6d, 0x67,
	0xb3, 0x33, 0x29, 0xa7, 0x9f, 0xfb, 0x6b, 0x61, 0xdb, 0xd6, 0x1a, 0x26, 0x9a, 0x0a, 0x4d, 0xb7,
	0x3f, 0x7d, 0xfc, 0x86, 0xcc, 0xbd, 0x7c, 0x30, 0x8f, 0x6b, 0x26, 0x3c, 0xf2, 0x0e, 0x7f, 0x70,
	0x53, 0x67, 0x4f, 0x70, 0xaa, 0x91, 0x4b, 0xf4, 0x77, 0xc9, 0x43, 0xf2, 0x9c, 0x16, 0x37, 0xf9,
	0xe1, 0x92, 0xbc, 0x0c, 0xdf, 0xdf, 0xe9, 0x2c, 0x5a, 0x2e, 0x6a, 0xba, 0xfd, 0x82, 0xab, 0xd8,
	0x1d, 0xc7, 0xa0, 0xa7, 0xd3, 0x90, 0xa5, 0xd2, 0x77, 0xb8, 0x88, 0xd5, 0x71, 0xb0, 0x1c, 0x68,
	0x2c, 0x14, 0x29, 0xec, 0x13, 0x2e, 0x63, 0xcd, 0x3a, 0x34, 0x95, 0x51, 0x34, 0x18, 0x7a, 0x8b,
	0xe0, 0xbf, 0x60, 0x76, 0x8b, 0xfe, 0xaf, 0xf7, 0x66, 0x26, 0x7c, 0x83, 0xf3, 0x38, 0x75, 0x0c,
	0x2a, 0x9a, 0x0a, 0x3b, 0xe9, 0x2a, 0xb4, 0x14, 0x5f, 0xd2, 0x22, 0xa8, 0x57, 0x48, 0x8f, 0xcf,
	0x5b, 0x57, 0x34, 0x34, 0xbd, 0xeb, 0xd9, 0xa8, 0x00, 0x88, 0x25, 0xdd, 0xd3, 0x4c, 0x08, 0xcd,
	0x35, 0xe5, 0xc9, 0x2e, 0x00, 0x00, 0xff, 0xff, 0x88, 0x5f, 0x90, 0x25, 0xea, 0x03, 0x00, 0x00,
}