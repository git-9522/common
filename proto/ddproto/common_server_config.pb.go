// Code generated by protoc-gen-go.
// source: common_server_config.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 对应mgo 中的表单,但是这里主要是做一些配置
type TConfigSys struct {
	Id               *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NewUserCoin      *int64 `protobuf:"varint,2,opt,name=newUserCoin" json:"newUserCoin,omitempty"`
	NewUserRoomcard  *int64 `protobuf:"varint,3,opt,name=newUserRoomcard" json:"newUserRoomcard,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TConfigSys) Reset()                    { *m = TConfigSys{} }
func (m *TConfigSys) String() string            { return proto.CompactTextString(m) }
func (*TConfigSys) ProtoMessage()               {}
func (*TConfigSys) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *TConfigSys) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TConfigSys) GetNewUserCoin() int64 {
	if m != nil && m.NewUserCoin != nil {
		return *m.NewUserCoin
	}
	return 0
}

func (m *TConfigSys) GetNewUserRoomcard() int64 {
	if m != nil && m.NewUserRoomcard != nil {
		return *m.NewUserRoomcard
	}
	return 0
}

// 转盘抽奖的配置
type TConfigDrawLottery struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             *int32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TConfigDrawLottery) Reset()                    { *m = TConfigDrawLottery{} }
func (m *TConfigDrawLottery) String() string            { return proto.CompactTextString(m) }
func (*TConfigDrawLottery) ProtoMessage()               {}
func (*TConfigDrawLottery) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *TConfigDrawLottery) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TConfigDrawLottery) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TConfigDrawLottery) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*TConfigSys)(nil), "ddproto.TConfigSys")
	proto.RegisterType((*TConfigDrawLottery)(nil), "ddproto.TConfigDrawLottery")
}

var fileDescriptor12 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0x94, 0xdc, 0xb8, 0xb8,
	0x42, 0x9c, 0xc1, 0x32, 0xc1, 0x95, 0xc5, 0x42, 0x5c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xac, 0x42, 0xc2, 0x5c, 0xdc, 0x79, 0xa9, 0xe5, 0xa1, 0x40, 0xdd, 0xce, 0xf9, 0x99,
	0x79, 0x12, 0x4c, 0x40, 0x41, 0x66, 0x21, 0x71, 0x2e, 0x7e, 0xa8, 0x60, 0x50, 0x7e, 0x7e, 0x6e,
	0x72, 0x62, 0x51, 0x8a, 0x04, 0x33, 0x48, 0x42, 0xc9, 0x86, 0x4b, 0x08, 0x6a, 0x8e, 0x4b, 0x51,
	0x62, 0xb9, 0x4f, 0x7e, 0x49, 0x49, 0x6a, 0x51, 0x25, 0x8a, 0x79, 0x3c, 0x5c, 0x2c, 0x79, 0x89,
	0xb9, 0xa9, 0x60, 0x83, 0x38, 0x41, 0xbc, 0x92, 0xca, 0x82, 0x54, 0xb0, 0x6e, 0x56, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x78, 0xc7, 0xc3, 0xab, 0x00, 0x00, 0x00,
}
