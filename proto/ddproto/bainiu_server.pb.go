// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bainiu_server.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of bainiu_client_poker from bainiu_base.proto

// Ignoring public import of bainiu_enum_protoid from bainiu_base.proto

// Ignoring public import of bainiu_area_name from bainiu_base.proto

// Ignoring public import of bainiu_enum_PokerType from bainiu_base.proto

// 打出去的牌
type BainiuSrvPoker struct {
	Pais             []*CommonSrvPokerPai  `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *BainiuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.BainiuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *BainiuSrvPoker) Reset()                    { *m = BainiuSrvPoker{} }
func (m *BainiuSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*BainiuSrvPoker) ProtoMessage()               {}
func (*BainiuSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BainiuSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *BainiuSrvPoker) GetType() BainiuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return BainiuEnum_PokerType_BAINIU_NO_NIU
}

func init() {
	proto.RegisterType((*BainiuSrvPoker)(nil), "ddproto.bainiu_srv_poker")
}

func init() { proto.RegisterFile("bainiu_server.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4a, 0xcc, 0xcc,
	0xcb, 0x2c, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4, 0x24, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xa0, 0xb2, 0xf1,
	0x05, 0xf9, 0xd9, 0x30, 0x35, 0x52, 0x82, 0x50, 0x8d, 0x49, 0x89, 0xc5, 0xa9, 0x10, 0x21, 0xa5,
	0x0a, 0x2e, 0x01, 0x98, 0x69, 0x45, 0x65, 0x10, 0xc5, 0x42, 0x06, 0x5c, 0x2c, 0x05, 0x89, 0x99,
	0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x32, 0x7a, 0x50, 0x93, 0xf5, 0x60, 0x06, 0xc3,
	0x14, 0x06, 0x24, 0x66, 0x06, 0x81, 0x55, 0x0a, 0x19, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a,
	0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xc1, 0x75, 0x40, 0x8d, 0x4e, 0xcd, 0x2b, 0xcd, 0x8d,
	0x0f, 0x00, 0x69, 0x09, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0xab, 0x0d, 0x60, 0x08, 0x60, 0x04, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x58, 0x73, 0xf9, 0xca, 0x00, 0x00, 0x00,
}
