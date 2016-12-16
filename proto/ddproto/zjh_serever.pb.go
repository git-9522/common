// Code generated by protoc-gen-go.
// source: zjh_serever.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// 三张牌的类型
type ZjhEnum_ZJHTYPE int32

const (
	ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI   ZjhEnum_ZJHTYPE = 1
	ZjhEnum_ZJHTYPE_ZJHTYPE_DUIZI    ZjhEnum_ZJHTYPE = 2
	ZjhEnum_ZJHTYPE_ZJHTYPE_LIANZI   ZjhEnum_ZJHTYPE = 3
	ZjhEnum_ZJHTYPE_ZJHTYPE_QING     ZjhEnum_ZJHTYPE = 4
	ZjhEnum_ZJHTYPE_ZJHTYPE_QINGLIAN ZjhEnum_ZJHTYPE = 5
	ZjhEnum_ZJHTYPE_ZJHTYPE_BAOZI    ZjhEnum_ZJHTYPE = 6
)

var ZjhEnum_ZJHTYPE_name = map[int32]string{
	1: "ZJHTYPE_GAOPAI",
	2: "ZJHTYPE_DUIZI",
	3: "ZJHTYPE_LIANZI",
	4: "ZJHTYPE_QING",
	5: "ZJHTYPE_QINGLIAN",
	6: "ZJHTYPE_BAOZI",
}
var ZjhEnum_ZJHTYPE_value = map[string]int32{
	"ZJHTYPE_GAOPAI":   1,
	"ZJHTYPE_DUIZI":    2,
	"ZJHTYPE_LIANZI":   3,
	"ZJHTYPE_QING":     4,
	"ZJHTYPE_QINGLIAN": 5,
	"ZJHTYPE_BAOZI":    6,
}

func (x ZjhEnum_ZJHTYPE) Enum() *ZjhEnum_ZJHTYPE {
	p := new(ZjhEnum_ZJHTYPE)
	*p = x
	return p
}
func (x ZjhEnum_ZJHTYPE) String() string {
	return proto.EnumName(ZjhEnum_ZJHTYPE_name, int32(x))
}
func (x *ZjhEnum_ZJHTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnum_ZJHTYPE_value, data, "ZjhEnum_ZJHTYPE")
	if err != nil {
		return err
	}
	*x = ZjhEnum_ZJHTYPE(value)
	return nil
}
func (ZjhEnum_ZJHTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

// 用户游戏状态
type ZjhEnumUserStatus int32

const (
	ZjhEnumUserStatus_zjh_S_GAMING   ZjhEnumUserStatus = 1
	ZjhEnumUserStatus_zjh_S_STANDING ZjhEnumUserStatus = 2
	ZjhEnumUserStatus_zjh_S_SITED    ZjhEnumUserStatus = 3
	ZjhEnumUserStatus_zjh_S_OFFLINE  ZjhEnumUserStatus = 4
)

var ZjhEnumUserStatus_name = map[int32]string{
	1: "zjh_S_GAMING",
	2: "zjh_S_STANDING",
	3: "zjh_S_SITED",
	4: "zjh_S_OFFLINE",
}
var ZjhEnumUserStatus_value = map[string]int32{
	"zjh_S_GAMING":   1,
	"zjh_S_STANDING": 2,
	"zjh_S_SITED":    3,
	"zjh_S_OFFLINE":  4,
}

func (x ZjhEnumUserStatus) Enum() *ZjhEnumUserStatus {
	p := new(ZjhEnumUserStatus)
	*p = x
	return p
}
func (x ZjhEnumUserStatus) String() string {
	return proto.EnumName(ZjhEnumUserStatus_name, int32(x))
}
func (x *ZjhEnumUserStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserStatus_value, data, "ZjhEnumUserStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserStatus(value)
	return nil
}
func (ZjhEnumUserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

// 打出去的牌
type ZjhSrvPoker struct {
	KeyValue         []int32              `protobuf:"varint,1,rep,name=keyValue" json:"keyValue,omitempty"`
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *ZjhEnum_ZJHTYPE     `protobuf:"varint,3,opt,name=type,enum=ddproto.ZjhEnum_ZJHTYPE" json:"type,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvPoker) Reset()                    { *m = ZjhSrvPoker{} }
func (m *ZjhSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvPoker) ProtoMessage()               {}
func (*ZjhSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *ZjhSrvPoker) GetKeyValue() []int32 {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *ZjhSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZjhSrvPoker) GetType() ZjhEnum_ZJHTYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI
}

func (m *ZjhSrvPoker) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 账单
type ZjhSrvBill struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ZjhSrvBill) Reset()                    { *m = ZjhSrvBill{} }
func (m *ZjhSrvBill) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvBill) ProtoMessage()               {}
func (*ZjhSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

// 用户的游戏数据
type ZjhSrv_GameData struct {
	Pai              *ZjhSrvPoker `protobuf:"bytes,1,opt,name=pai" json:"pai,omitempty"`
	Bill             *int64       `protobuf:"varint,2,opt,name=bill" json:"bill,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhSrv_GameData) Reset()                    { *m = ZjhSrv_GameData{} }
func (m *ZjhSrv_GameData) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrv_GameData) ProtoMessage()               {}
func (*ZjhSrv_GameData) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *ZjhSrv_GameData) GetPai() *ZjhSrvPoker {
	if m != nil {
		return m.Pai
	}
	return nil
}

func (m *ZjhSrv_GameData) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

// 用户属性
type ZjhSrvUser struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64             `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	RoomId           *int32             `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	DeskId           *int32             `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32             `protobuf:"varint,5,opt,name=gameNumber" json:"gameNumber,omitempty"`
	Data             *ZjhSrv_GameData   `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	Status           *ZjhEnumUserStatus `protobuf:"varint,7,opt,name=status,enum=ddproto.ZjhEnumUserStatus" json:"status,omitempty"`
	IsDiscard        *bool              `protobuf:"varint,8,opt,name=isDiscard" json:"isDiscard,omitempty"`
	IsWatch          *bool              `protobuf:"varint,9,opt,name=isWatch" json:"isWatch,omitempty"`
	IsOnline         *bool              `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	IsLost           *bool              `protobuf:"varint,11,opt,name=isLost" json:"isLost,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhSrvUser) Reset()                    { *m = ZjhSrvUser{} }
func (m *ZjhSrvUser) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvUser) ProtoMessage()               {}
func (*ZjhSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

func (m *ZjhSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvUser) GetData() *ZjhSrv_GameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ZjhSrvUser) GetStatus() ZjhEnumUserStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZjhEnumUserStatus_zjh_S_GAMING
}

func (m *ZjhSrvUser) GetIsDiscard() bool {
	if m != nil && m.IsDiscard != nil {
		return *m.IsDiscard
	}
	return false
}

func (m *ZjhSrvUser) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *ZjhSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *ZjhSrvUser) GetIsLost() bool {
	if m != nil && m.IsLost != nil {
		return *m.IsLost
	}
	return false
}

// desk 的信息
type ZjhSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32               `protobuf:"varint,2,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	AllPokers        []*CommonSrvPokerPai `protobuf:"bytes,4,rep,name=allPokers" json:"allPokers,omitempty"`
	LastUser         *uint32              `protobuf:"varint,5,opt,name=lastUser" json:"lastUser,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	IsGamming        *bool                `protobuf:"varint,7,opt,name=isGamming" json:"isGamming,omitempty"`
	CuurBaseValue    *int64               `protobuf:"varint,8,opt,name=cuurBaseValue" json:"cuurBaseValue,omitempty"`
	MinUser          *int32               `protobuf:"varint,9,opt,name=minUser" json:"minUser,omitempty"`
	BaseValue        *int64               `protobuf:"varint,10,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64               `protobuf:"varint,11,opt,name=maxValue" json:"maxValue,omitempty"`
	GameStatus       *int32               `protobuf:"varint,12,opt,name=gameStatus" json:"gameStatus,omitempty"`
	GamerNum         *int32               `protobuf:"varint,13,opt,name=gamerNum" json:"gamerNum,omitempty"`
	CircleNo         *int32               `protobuf:"varint,14,opt,name=circleNo" json:"circleNo,omitempty"`
	XuepinBaseValue  *int64               `protobuf:"varint,15,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvDesk) Reset()                    { *m = ZjhSrvDesk{} }
func (m *ZjhSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvDesk) ProtoMessage()               {}
func (*ZjhSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *ZjhSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvDesk) GetAllPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokers
	}
	return nil
}

func (m *ZjhSrvDesk) GetLastUser() uint32 {
	if m != nil && m.LastUser != nil {
		return *m.LastUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *ZjhSrvDesk) GetIsGamming() bool {
	if m != nil && m.IsGamming != nil {
		return *m.IsGamming
	}
	return false
}

func (m *ZjhSrvDesk) GetCuurBaseValue() int64 {
	if m != nil && m.CuurBaseValue != nil {
		return *m.CuurBaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ZjhSrvDesk) GetGamerNum() int32 {
	if m != nil && m.GamerNum != nil {
		return *m.GamerNum
	}
	return 0
}

func (m *ZjhSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

// room 的信息
type ZjhSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseValue        *int64  `protobuf:"varint,5,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64  `protobuf:"varint,6,opt,name=maxValue" json:"maxValue,omitempty"`
	XuepinBaseValue  *int64  `protobuf:"varint,7,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZjhSrvRoom) Reset()                    { *m = ZjhSrvRoom{} }
func (m *ZjhSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvRoom) ProtoMessage()               {}
func (*ZjhSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

func (m *ZjhSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *ZjhSrvRoom) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

func init() {
	proto.RegisterType((*ZjhSrvPoker)(nil), "ddproto.zjh_srv_poker")
	proto.RegisterType((*ZjhSrvBill)(nil), "ddproto.zjh_srv_bill")
	proto.RegisterType((*ZjhSrv_GameData)(nil), "ddproto.zjh_srv_GameData")
	proto.RegisterType((*ZjhSrvUser)(nil), "ddproto.zjh_srv_user")
	proto.RegisterType((*ZjhSrvDesk)(nil), "ddproto.zjh_srv_desk")
	proto.RegisterType((*ZjhSrvRoom)(nil), "ddproto.zjh_srv_room")
	proto.RegisterEnum("ddproto.ZjhEnum_ZJHTYPE", ZjhEnum_ZJHTYPE_name, ZjhEnum_ZJHTYPE_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserStatus", ZjhEnumUserStatus_name, ZjhEnumUserStatus_value)
}

var fileDescriptor16 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x4e, 0xdb, 0x40,
	0x10, 0xad, 0x73, 0x23, 0xd9, 0xc4, 0xc1, 0x6c, 0x6f, 0x4b, 0xc5, 0x03, 0x4a, 0x1f, 0x8a, 0x50,
	0x45, 0x25, 0xfe, 0x20, 0x28, 0x81, 0xba, 0x4a, 0x9d, 0x54, 0x09, 0x45, 0xe5, 0xc5, 0x5a, 0xec,
	0x15, 0x6c, 0xf1, 0x25, 0xf2, 0xda, 0x08, 0xfa, 0x01, 0x55, 0x3f, 0xa0, 0x3f, 0xd5, 0x8f, 0xaa,
	0xd4, 0x99, 0xb5, 0x4d, 0x1c, 0x11, 0xa9, 0x6f, 0x99, 0xb3, 0x67, 0x67, 0xce, 0x39, 0x9e, 0x0d,
	0xd9, 0xf9, 0xf1, 0xfd, 0xc6, 0x55, 0x22, 0x11, 0x77, 0x22, 0x39, 0x5a, 0x26, 0x71, 0x1a, 0xd3,
	0x2d, 0xdf, 0xd7, 0x3f, 0xde, 0xec, 0x7a, 0x71, 0x18, 0xc6, 0x11, 0x1e, 0xc3, 0xa9, 0xbb, 0x8c,
	0x6f, 0x4b, 0xce, 0xe0, 0xa7, 0x41, 0x4c, 0x7d, 0x33, 0xb9, 0xcb, 0x71, 0x6a, 0x91, 0xf6, 0xad,
	0x78, 0xf8, 0xca, 0x83, 0x4c, 0x30, 0x63, 0xbf, 0x7e, 0xd0, 0xa4, 0x87, 0xa4, 0xb1, 0xe4, 0x52,
	0xb1, 0x1a, 0x54, 0xdd, 0xe3, 0xbd, 0xa3, 0xa2, 0xed, 0x51, 0xd9, 0xb5, 0xbc, 0x3a, 0xe3, 0x92,
	0xbe, 0x23, 0x8d, 0xf4, 0x61, 0x29, 0x58, 0x7d, 0xdf, 0x38, 0xe8, 0x1f, 0xef, 0x3e, 0x72, 0x71,
	0x86, 0x88, 0xb2, 0xd0, 0xbd, 0xfc, 0xf4, 0x71, 0xf1, 0x6d, 0x36, 0xa6, 0x7d, 0xd2, 0xca, 0x40,
	0x8f, 0xed, 0xb3, 0x0e, 0x50, 0xcd, 0x41, 0x9f, 0xf4, 0x4a, 0x1d, 0x57, 0x32, 0x08, 0x06, 0x63,
	0x62, 0x95, 0xf5, 0x19, 0x0f, 0xc5, 0x88, 0xa7, 0x9c, 0xbe, 0x25, 0x75, 0x10, 0x02, 0xaa, 0x0c,
	0xd0, 0xf1, 0x6a, 0xad, 0xf7, 0x4a, 0x7f, 0x8f, 0x34, 0xb0, 0x01, 0xa8, 0x35, 0x0e, 0xea, 0x83,
	0xbf, 0xc6, 0xaa, 0x2f, 0xce, 0xab, 0xcc, 0xc5, 0x36, 0x26, 0xd2, 0xbd, 0x58, 0x46, 0x39, 0x1d,
	0x4f, 0x93, 0x38, 0x0e, 0xe1, 0x14, 0x0d, 0x34, 0xb1, 0xf6, 0x85, 0xba, 0x85, 0xba, 0xa1, 0x6b,
	0x4a, 0xc8, 0x35, 0xa8, 0x71, 0xb2, 0xf0, 0x4a, 0x24, 0xac, 0xa9, 0x31, 0xb0, 0xec, 0x83, 0x3a,
	0xd6, 0xd2, 0xb2, 0x76, 0x9f, 0xc8, 0x7a, 0x94, 0xff, 0x9e, 0xb4, 0x54, 0xca, 0xd3, 0x4c, 0xb1,
	0x2d, 0x9d, 0xce, 0xde, 0xd3, 0x74, 0x50, 0xda, 0x5c, 0x73, 0xe8, 0x0e, 0xe9, 0x48, 0x35, 0x92,
	0xca, 0xe3, 0x89, 0xcf, 0xda, 0x70, 0xa1, 0x4d, 0xb7, 0xc9, 0x96, 0x54, 0x17, 0x3c, 0xf5, 0x6e,
	0x74, 0x68, 0x6d, 0xfc, 0x56, 0x52, 0x4d, 0xa3, 0x40, 0x46, 0x82, 0x11, 0x8d, 0x80, 0x60, 0xa9,
	0x26, 0xb1, 0x4a, 0x59, 0x17, 0xeb, 0xc1, 0x9f, 0xda, 0xca, 0x3f, 0x3a, 0xa9, 0x38, 0x32, 0x36,
	0x38, 0xaa, 0x95, 0xae, 0xd7, 0x52, 0xf8, 0x40, 0x3a, 0x3c, 0x08, 0x66, 0x18, 0xaf, 0x82, 0x20,
	0xfe, 0xbf, 0x05, 0xa0, 0x2b, 0xe0, 0x2a, 0x3d, 0x57, 0x45, 0x48, 0x26, 0xba, 0x41, 0xe4, 0x02,
	0x94, 0x26, 0x3a, 0x29, 0x33, 0x37, 0x08, 0xe1, 0x84, 0x32, 0xba, 0xd6, 0x89, 0xb4, 0xe9, 0x4b,
	0x62, 0x7a, 0x59, 0x96, 0x9c, 0x70, 0x25, 0xf2, 0x05, 0x6c, 0xeb, 0xaf, 0x02, 0xbe, 0x81, 0xa4,
	0xbb, 0x75, 0xb4, 0x20, 0xb8, 0x7a, 0xf5, 0xc8, 0x21, 0x9a, 0x03, 0x23, 0x43, 0x7e, 0x9f, 0x23,
	0x5d, 0x8d, 0x14, 0xce, 0xf2, 0x38, 0x59, 0x4f, 0x5f, 0x04, 0x16, 0x62, 0x09, 0xd8, 0x65, 0x66,
	0x89, 0x78, 0x32, 0xf1, 0x02, 0xe1, 0xc4, 0xac, 0xaf, 0x91, 0xd7, 0x64, 0xfb, 0x3e, 0x13, 0x4b,
	0x19, 0xad, 0x64, 0x6c, 0xeb, 0x5d, 0xfa, 0x5d, 0xd9, 0x25, 0xcc, 0xa7, 0x92, 0x93, 0x51, 0xf6,
	0xc2, 0x7a, 0x81, 0x0f, 0xa0, 0x56, 0x0a, 0x45, 0x64, 0x02, 0xaf, 0x32, 0x28, 0xc2, 0x2c, 0xa0,
	0x85, 0x4c, 0x03, 0xa1, 0xb7, 0xaa, 0xb3, 0x6e, 0xa7, 0xf9, 0xc4, 0x4e, 0x4b, 0x23, 0x1b, 0x64,
	0x61, 0x68, 0xf5, 0xc3, 0x5f, 0x46, 0xfe, 0x54, 0xd6, 0x9e, 0x17, 0x25, 0xfd, 0xe2, 0xa7, 0x7b,
	0x36, 0x9c, 0xce, 0x86, 0xb6, 0x65, 0xc0, 0x18, 0xb3, 0xc4, 0x46, 0xe7, 0xf6, 0xa5, 0x6d, 0xd5,
	0xaa, 0xb4, 0x89, 0x3d, 0x74, 0x00, 0xc3, 0xd1, 0xbd, 0x12, 0xfb, 0x62, 0x3b, 0x67, 0x56, 0x83,
	0xbe, 0x20, 0x56, 0x15, 0x41, 0xa6, 0xd5, 0xac, 0xb6, 0x3b, 0x19, 0x4e, 0xe1, 0x6a, 0xeb, 0xd0,
	0x25, 0xcf, 0x37, 0xad, 0xb2, 0x95, 0xe7, 0x36, 0x07, 0x29, 0x9f, 0xb1, 0xa3, 0x81, 0x73, 0x73,
	0x64, 0xbe, 0x18, 0x3a, 0x23, 0xc4, 0x6a, 0xf0, 0x95, 0xbb, 0x05, 0x66, 0x2f, 0xc6, 0x23, 0x10,
	0xb2, 0x93, 0xff, 0x35, 0xcd, 0xdd, 0xe9, 0xe9, 0xe9, 0xc4, 0x76, 0xc6, 0x56, 0x63, 0xf6, 0xec,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xd8, 0x46, 0x6c, 0xe8, 0x04, 0x00, 0x00,
}
