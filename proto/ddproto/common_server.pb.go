// Code generated by protoc-gen-go.
// source: common_server.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户的seesion 主要是游戏的
type GameSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	GameId           *int32  `protobuf:"varint,5,opt,name=GameId" json:"GameId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=GameNumber" json:"GameNumber,omitempty"`
	GameCustomStatus *int32  `protobuf:"varint,7,opt,name=GameCustomStatus" json:"GameCustomStatus,omitempty"`
	IsBreak          *bool   `protobuf:"varint,8,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,9,opt,name=isLeave" json:"isLeave,omitempty"`
	RoomType         *int32  `protobuf:"varint,10,opt,name=roomType" json:"roomType,omitempty"`
	RoomPassword     *string `protobuf:"bytes,11,opt,name=roomPassword" json:"roomPassword,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,12,opt,name=RoomLevel" json:"RoomLevel,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameSession) Reset()                    { *m = GameSession{} }
func (m *GameSession) String() string            { return proto.CompactTextString(m) }
func (*GameSession) ProtoMessage()               {}
func (*GameSession) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GameSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *GameSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *GameSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *GameSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *GameSession) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameSession) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *GameSession) GetGameCustomStatus() int32 {
	if m != nil && m.GameCustomStatus != nil {
		return *m.GameCustomStatus
	}
	return 0
}

func (m *GameSession) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *GameSession) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *GameSession) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *GameSession) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *GameSession) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

// 服务器游戏玩家通过的协议
type CommonSrvGameUser struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Coin             *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	RoomId           *int32  `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,5,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,6,opt,name=gameNumber" json:"gameNumber,omitempty"`
	IsBreak          *bool   `protobuf:"varint,7,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool   `protobuf:"varint,8,opt,name=isLeave" json:"isLeave,omitempty"`
	Status           *int32  `protobuf:"varint,9,opt,name=status" json:"status,omitempty"`
	WaitTime         *int64  `protobuf:"varint,10,opt,name=waitTime" json:"waitTime,omitempty"`
	RoomType         *int32  `protobuf:"varint,11,opt,name=roomType" json:"roomType,omitempty"`
	RoomPassword     *string `protobuf:"bytes,12,opt,name=roomPassword" json:"roomPassword,omitempty"`
	IsRobot          *bool   `protobuf:"varint,13,opt,name=isRobot" json:"isRobot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameUser) Reset()                    { *m = CommonSrvGameUser{} }
func (m *CommonSrvGameUser) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameUser) ProtoMessage()               {}
func (*CommonSrvGameUser) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CommonSrvGameUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CommonSrvGameUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *CommonSrvGameUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *CommonSrvGameUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *CommonSrvGameUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvGameUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *CommonSrvGameUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *CommonSrvGameUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *CommonSrvGameUser) GetWaitTime() int64 {
	if m != nil && m.WaitTime != nil {
		return *m.WaitTime
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *CommonSrvGameUser) GetRoomPassword() string {
	if m != nil && m.RoomPassword != nil {
		return *m.RoomPassword
	}
	return ""
}

func (m *CommonSrvGameUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

// 通用的desk 的协议
type CommonSrvGameDesk struct {
	DeskId           *int32  `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32  `protobuf:"varint,2,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32  `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	Round            *int32  `protobuf:"varint,4,opt,name=round" json:"round,omitempty"`
	Password         *string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	Owner            *uint32 `protobuf:"varint,6,opt,name=owner" json:"owner,omitempty"`
	Banker           *uint32 `protobuf:"varint,7,opt,name=banker" json:"banker,omitempty"`
	CreateFee        *int64  `protobuf:"varint,8,opt,name=createFee" json:"createFee,omitempty"`
	Status           *int32  `protobuf:"varint,9,opt,name=status" json:"status,omitempty"`
	BoardsCout       *int32  `protobuf:"varint,10,opt,name=boardsCout" json:"boardsCout,omitempty"`
	BaseValue        *int64  `protobuf:"varint,11,opt,name=baseValue" json:"baseValue,omitempty"`
	BeginTime        *string `protobuf:"bytes,12,opt,name=beginTime" json:"beginTime,omitempty"`
	EndTime          *string `protobuf:"bytes,13,opt,name=endTime" json:"endTime,omitempty"`
	NInitActionTime  *int32  `protobuf:"varint,14,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	UserCountLimit   *int32  `protobuf:"varint,15,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommonSrvGameDesk) Reset()                    { *m = CommonSrvGameDesk{} }
func (m *CommonSrvGameDesk) String() string            { return proto.CompactTextString(m) }
func (*CommonSrvGameDesk) ProtoMessage()               {}
func (*CommonSrvGameDesk) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *CommonSrvGameDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *CommonSrvGameDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *CommonSrvGameDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *CommonSrvGameDesk) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *CommonSrvGameDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CommonSrvGameDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *CommonSrvGameDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *CommonSrvGameDesk) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *CommonSrvGameDesk) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *CommonSrvGameDesk) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *CommonSrvGameDesk) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *CommonSrvGameDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*GameSession)(nil), "ddproto.GameSession")
	proto.RegisterType((*CommonSrvGameUser)(nil), "ddproto.common_srv_game_user")
	proto.RegisterType((*CommonSrvGameDesk)(nil), "ddproto.common_srv_game_desk")
}

<<<<<<< HEAD
var fileDescriptor5 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xba, 0xa9, 0xe3, 0xc9, 0xbf, 0x62, 0x2a, 0xd8, 0x23, 0xca, 0x89, 0x13, 0xef,
	0x00, 0x45, 0xa0, 0x4a, 0x51, 0x85, 0x4a, 0xe1, 0x1a, 0xad, 0xed, 0x51, 0xb4, 0x4a, 0xbd, 0x13,
	0xed, 0xae, 0x13, 0xf1, 0x48, 0xbc, 0x1b, 0x77, 0xae, 0xcc, 0xce, 0x26, 0x51, 0x68, 0x73, 0xf3,
	0x7c, 0xfe, 0x66, 0xf7, 0xfb, 0xcd, 0x2c, 0xbc, 0x6e, 0xa8, 0xeb, 0xc8, 0x2e, 0x3d, 0xba, 0x2d,
	0xba, 0x0f, 0x1b, 0x47, 0x81, 0xaa, 0xa2, 0x6d, 0xe5, 0x63, 0xfe, 0x27, 0x83, 0xd1, 0x57, 0xdd,
	0xe1, 0x77, 0xf4, 0xde, 0x90, 0xad, 0xa6, 0x70, 0xf5, 0x83, 0x9d, 0x77, 0xad, 0xca, 0xde, 0x65,
	0xef, 0x27, 0xb1, 0x7e, 0x20, 0xea, 0xb8, 0xbe, 0xe0, 0x7a, 0x10, 0xeb, 0xcf, 0xe8, 0xd7, 0x5c,
	0xe7, 0x52, 0x57, 0x00, 0xd2, 0x1e, 0x74, 0xe8, 0xbd, 0xba, 0x3c, 0x78, 0xa2, 0xc6, 0x9e, 0xc1,
	0xa9, 0xe7, 0xbe, 0xef, 0x6a, 0x74, 0xea, 0x4a, 0x34, 0x05, 0xd7, 0x51, 0xbb, 0xed, 0x7d, 0xa0,
	0x6e, 0xdf, 0x5d, 0xc8, 0x9f, 0x19, 0x14, 0xc6, 0x7f, 0x72, 0xa8, 0xd7, 0x6a, 0xc8, 0xc2, 0x30,
	0x09, 0x0b, 0xd4, 0x5b, 0x54, 0xa5, 0x08, 0xd7, 0x30, 0x74, 0x9c, 0xe9, 0xf1, 0xd7, 0x06, 0x15,
	0x48, 0xcf, 0x0d, 0x8c, 0xa3, 0xf2, 0x4d, 0x7b, 0xbf, 0x23, 0xd7, 0xaa, 0x11, 0xab, 0x65, 0xf5,
	0x0a, 0xca, 0x98, 0x7d, 0x81, 0x5b, 0x7c, 0x52, 0xe3, 0x68, 0x9c, 0xff, 0xcd, 0xe0, 0xe6, 0x30,
	0x0f, 0xb7, 0x5d, 0xae, 0x38, 0xc2, 0xb2, 0x67, 0xde, 0x98, 0xb9, 0x3f, 0xe5, 0xe6, 0x3b, 0xac,
	0x69, 0xd6, 0xf7, 0x6c, 0x10, 0xf2, 0xb2, 0x1a, 0xc3, 0x65, 0x43, 0xc6, 0x0a, 0x77, 0x1e, 0xfd,
	0x2e, 0xcd, 0xe5, 0xc8, 0xdc, 0xa6, 0xb9, 0x1c, 0x99, 0x57, 0xcf, 0x99, 0x4f, 0xc8, 0x8a, 0xe7,
	0x64, 0x09, 0x95, 0x4f, 0xf1, 0x69, 0x16, 0xa5, 0x74, 0x70, 0x8a, 0x9d, 0x36, 0xe1, 0xd1, 0x74,
	0x89, 0x34, 0xff, 0x8f, 0x7d, 0x74, 0x96, 0x7d, 0x2c, 0x69, 0xe5, 0xe8, 0x07, 0xaa, 0x29, 0xa8,
	0x49, 0x3c, 0x7a, 0xfe, 0xfb, 0xe2, 0x25, 0x79, 0x4c, 0x7c, 0x92, 0x3c, 0x3b, 0x93, 0xfc, 0xb8,
	0xf5, 0x3d, 0x6d, 0xda, 0xfa, 0x04, 0x06, 0x8e, 0x7a, 0x7b, 0x80, 0xe7, 0x50, 0x9b, 0xc3, 0xf5,
	0x03, 0xb9, 0x9e, 0x0d, 0xb4, 0xb3, 0x7b, 0x72, 0x79, 0x45, 0xb5, 0xb6, 0x6b, 0xae, 0x0b, 0xa9,
	0x79, 0x33, 0x0d, 0xcf, 0x21, 0xe0, 0x17, 0x4c, 0xe8, 0xf9, 0x0b, 0x74, 0x8e, 0x51, 0x93, 0x76,
	0xad, 0xbf, 0xa5, 0x3e, 0xec, 0xd7, 0xcc, 0x6d, 0xb5, 0xf6, 0xf8, 0x53, 0x3f, 0xf5, 0x89, 0x3e,
	0x17, 0x09, 0x57, 0xc6, 0xca, 0x88, 0x8e, 0xe8, 0x68, 0x5b, 0x11, 0x26, 0x22, 0xbc, 0x85, 0x99,
	0xbd, 0xb3, 0x26, 0x7c, 0x6c, 0x02, 0x3f, 0x71, 0xf9, 0x31, 0x95, 0xf3, 0xde, 0xc0, 0x34, 0x2e,
	0x9d, 0x6f, 0xb0, 0x61, 0x61, 0x3a, 0x13, 0xd4, 0x2c, 0xea, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0xfc, 0xe9, 0x48, 0x33, 0x03, 0x00, 0x00,
=======
var fileDescriptor4 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0x99, 0xce, 0x76, 0xdb, 0xa6, 0xed, 0xae, 0xc4, 0x65, 0x09, 0x22, 0x4b, 0xe9, 0x85,
	0x14, 0x2f, 0x7c, 0x07, 0xad, 0x28, 0x85, 0xb2, 0x48, 0x5c, 0xbd, 0x2d, 0xe9, 0xe4, 0x50, 0x42,
	0x37, 0x49, 0x49, 0x32, 0x2d, 0xbe, 0x9f, 0x6f, 0xe3, 0x8d, 0x8f, 0x20, 0xc9, 0x99, 0xe9, 0x74,
	0xb6, 0xae, 0xde, 0xcd, 0xf7, 0x9f, 0x39, 0x67, 0x32, 0xdf, 0x09, 0x79, 0x59, 0x58, 0xad, 0xad,
	0x59, 0x79, 0x70, 0x7b, 0x70, 0xef, 0x76, 0xce, 0x06, 0x4b, 0x7b, 0x52, 0xa6, 0x87, 0xe9, 0xaf,
	0x0e, 0x19, 0x7e, 0x16, 0x1a, 0xbe, 0x82, 0xf7, 0xca, 0x1a, 0x7a, 0x4b, 0x2e, 0xbf, 0x79, 0x70,
	0x0b, 0xc9, 0xb2, 0x49, 0x36, 0x1b, 0xf3, 0x8a, 0x62, 0xce, 0xad, 0xd5, 0x0b, 0xc9, 0x3a, 0x93,
	0x6c, 0xd6, 0xe5, 0x15, 0xc5, 0xfc, 0x23, 0xf8, 0xed, 0x42, 0xb2, 0x1c, 0x73, 0x24, 0x7a, 0x47,
	0x48, 0x1a, 0x1b, 0x44, 0x28, 0x3d, 0xbb, 0x48, 0xb5, 0x93, 0x24, 0xf6, 0x45, 0x5a, 0x48, 0xd6,
	0xc5, 0x3e, 0xa4, 0xba, 0xef, 0xbe, 0xd4, 0x6b, 0x70, 0xec, 0xb2, 0xe9, 0xc3, 0x84, 0xbe, 0x25,
	0x2f, 0x22, 0xcd, 0x4b, 0x1f, 0xac, 0xae, 0xa6, 0xf7, 0xd2, 0x5b, 0x67, 0x39, 0x65, 0xa4, 0xa7,
	0xfc, 0x07, 0x07, 0x62, 0xcb, 0xfa, 0x93, 0x6c, 0xd6, 0xe7, 0x35, 0x62, 0x65, 0x09, 0x62, 0x0f,
	0x6c, 0x50, 0x57, 0x12, 0xd2, 0x57, 0xa4, 0xef, 0xac, 0xd5, 0x0f, 0x3f, 0x76, 0xc0, 0x48, 0x9a,
	0x7b, 0x64, 0x3a, 0x25, 0xa3, 0xf8, 0xfc, 0x45, 0x78, 0x7f, 0xb0, 0x4e, 0xb2, 0xe1, 0x24, 0x9b,
	0x0d, 0x78, 0x2b, 0xa3, 0xaf, 0xc9, 0x20, 0x9a, 0x59, 0xc2, 0x1e, 0x1e, 0xd9, 0x28, 0x0d, 0x68,
	0x82, 0xe9, 0xef, 0x0e, 0xb9, 0xa9, 0xd7, 0xe1, 0xf6, 0xab, 0x8d, 0xd0, 0xb0, 0x2a, 0x3d, 0xb8,
	0xa8, 0xa3, 0x6c, 0x69, 0x47, 0x8a, 0xc7, 0x31, 0xaa, 0xd8, 0xde, 0x0b, 0x0d, 0x49, 0xfc, 0x80,
	0x1f, 0x99, 0x52, 0x72, 0x51, 0x58, 0x65, 0x92, 0xf8, 0x9c, 0xa7, 0xe7, 0x38, 0xc7, 0xe1, 0x9a,
	0x50, 0x79, 0x45, 0x31, 0x97, 0xb8, 0xa6, 0x4a, 0xb7, 0x3c, 0xae, 0x69, 0x73, 0xa6, 0xbb, 0x49,
	0x4e, 0x15, 0xf6, 0x9e, 0x55, 0xd8, 0x6f, 0x2b, 0xbc, 0x25, 0x97, 0x1e, 0x17, 0x33, 0xc0, 0x6f,
	0x21, 0xc5, 0x7f, 0x39, 0x08, 0x15, 0x1e, 0x94, 0x46, 0xb5, 0x39, 0x3f, 0x72, 0x4b, 0xfb, 0xf0,
	0x3f, 0xda, 0x47, 0x7f, 0xd1, 0x9e, 0x4e, 0xc3, 0xed, 0xda, 0x06, 0x36, 0xae, 0x4f, 0x93, 0x70,
	0xfa, 0x33, 0x3f, 0x57, 0x1e, 0x7f, 0xfe, 0x44, 0x49, 0xf6, 0x0f, 0x25, 0x9d, 0x33, 0x25, 0x8d,
	0xe2, 0xbc, 0xa5, 0xf8, 0x86, 0x74, 0x9d, 0x2d, 0x4d, 0x6d, 0x1e, 0x21, 0xfe, 0xd8, 0xae, 0x3e,
	0x78, 0x17, 0x17, 0x58, 0x73, 0xec, 0xb0, 0x07, 0x53, 0x79, 0x1f, 0x73, 0x84, 0x38, 0x7f, 0x2d,
	0xcc, 0x16, 0x5c, 0x32, 0x3e, 0xe6, 0x15, 0xc5, 0x9b, 0x55, 0x38, 0x10, 0x01, 0x3e, 0x01, 0x2a,
	0xcf, 0x79, 0x13, 0x3c, 0x2b, 0xfd, 0x8e, 0x90, 0xb5, 0x15, 0x4e, 0xfa, 0xb9, 0x2d, 0x43, 0x75,
	0xa3, 0x4f, 0x92, 0x38, 0x75, 0x2d, 0x3c, 0x7c, 0x17, 0x8f, 0x25, 0x9a, 0xcf, 0x79, 0x13, 0xa4,
	0x2a, 0x6c, 0x94, 0x49, 0x3b, 0x43, 0xef, 0x4d, 0x10, 0xa5, 0x83, 0x91, 0xa9, 0x36, 0x4e, 0xb5,
	0x1a, 0xe9, 0x8c, 0x5c, 0x9b, 0x85, 0x51, 0xe1, 0x7d, 0x11, 0x94, 0xc5, 0xee, 0xab, 0xf4, 0xe9,
	0xa7, 0x31, 0x7d, 0x43, 0xae, 0xe2, 0x55, 0x9f, 0xdb, 0xd2, 0x84, 0xa5, 0xd2, 0x2a, 0xb0, 0xeb,
	0xf4, 0xe2, 0x93, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x5a, 0x21, 0xbd, 0xc6, 0x04,
	0x00, 0x00,
>>>>>>> 909db0f792448fe2ab44876c4ce012340a336faf
}
