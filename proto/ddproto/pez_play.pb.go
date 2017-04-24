// Code generated by protoc-gen-go.
// source: pez_play.proto
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

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// 链接类型
type PEZ_RECONNECT_TYPE int32

const (
	PEZ_RECONNECT_TYPE_PEZ_NORMAL    PEZ_RECONNECT_TYPE = 1
	PEZ_RECONNECT_TYPE_PEZ_RECONNECT PEZ_RECONNECT_TYPE = 2
)

var PEZ_RECONNECT_TYPE_name = map[int32]string{
	1: "PEZ_NORMAL",
	2: "PEZ_RECONNECT",
}
var PEZ_RECONNECT_TYPE_value = map[string]int32{
	"PEZ_NORMAL":    1,
	"PEZ_RECONNECT": 2,
}

func (x PEZ_RECONNECT_TYPE) Enum() *PEZ_RECONNECT_TYPE {
	p := new(PEZ_RECONNECT_TYPE)
	*p = x
	return p
}
func (x PEZ_RECONNECT_TYPE) String() string {
	return proto.EnumName(PEZ_RECONNECT_TYPE_name, int32(x))
}
func (x *PEZ_RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PEZ_RECONNECT_TYPE_value, data, "PEZ_RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = PEZ_RECONNECT_TYPE(value)
	return nil
}
func (PEZ_RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor36, []int{0} }

// 积分
type PezUserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezUserCoinBean) Reset()                    { *m = PezUserCoinBean{} }
func (m *PezUserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*PezUserCoinBean) ProtoMessage()               {}
func (*PezUserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{0} }

func (m *PezUserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezUserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 开局（接收服务端消息）
type Pez_Opening struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrPlayCount    *int32             `protobuf:"varint,2,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	Dice1            *int32             `protobuf:"varint,3,opt,name=dice1" json:"dice1,omitempty"`
	Dice2            *int32             `protobuf:"varint,4,opt,name=dice2" json:"dice2,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,5,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_Opening) Reset()                    { *m = Pez_Opening{} }
func (m *Pez_Opening) String() string            { return proto.CompactTextString(m) }
func (*Pez_Opening) ProtoMessage()               {}
func (*Pez_Opening) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{1} }

func (m *Pez_Opening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Opening) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *Pez_Opening) GetDice1() int32 {
	if m != nil && m.Dice1 != nil {
		return *m.Dice1
	}
	return 0
}

func (m *Pez_Opening) GetDice2() int32 {
	if m != nil && m.Dice2 != nil {
		return *m.Dice2
	}
	return 0
}

func (m *Pez_Opening) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发牌
type Pez_DealCards struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerCard       []*PezBase_PlayerCard `protobuf:"bytes,2,rep,name=playerCard" json:"playerCard,omitempty"`
	PaiCount         *int32                `protobuf:"varint,3,opt,name=paiCount" json:"paiCount,omitempty"`
	DealerUserId     *uint32               `protobuf:"varint,4,opt,name=dealerUserId" json:"dealerUserId,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Pez_DealCards) Reset()                    { *m = Pez_DealCards{} }
func (m *Pez_DealCards) String() string            { return proto.CompactTextString(m) }
func (*Pez_DealCards) ProtoMessage()               {}
func (*Pez_DealCards) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{2} }

func (m *Pez_DealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DealCards) GetPlayerCard() []*PezBase_PlayerCard {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *Pez_DealCards) GetPaiCount() int32 {
	if m != nil && m.PaiCount != nil {
		return *m.PaiCount
	}
	return 0
}

func (m *Pez_DealCards) GetDealerUserId() uint32 {
	if m != nil && m.DealerUserId != nil {
		return *m.DealerUserId
	}
	return 0
}

// 押注
type Pez_Bet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetNum           *int32  `protobuf:"varint,2,opt,name=betNum" json:"betNum,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_Bet) Reset()                    { *m = Pez_Bet{} }
func (m *Pez_Bet) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bet) ProtoMessage()               {}
func (*Pez_Bet) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{3} }

func (m *Pez_Bet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_Bet) GetBetNum() int32 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *Pez_Bet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_AckBet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetCount         *int32  `protobuf:"varint,2,opt,name=betCount" json:"betCount,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_AckBet) Reset()                    { *m = Pez_AckBet{} }
func (m *Pez_AckBet) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckBet) ProtoMessage()               {}
func (*Pez_AckBet) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{4} }

func (m *Pez_AckBet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckBet) GetBetCount() int32 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *Pez_AckBet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

// 开牌
type Pez_OpenCard struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CardId           *int32       `protobuf:"varint,2,opt,name=cardId" json:"cardId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_OpenCard) Reset()                    { *m = Pez_OpenCard{} }
func (m *Pez_OpenCard) String() string            { return proto.CompactTextString(m) }
func (*Pez_OpenCard) ProtoMessage()               {}
func (*Pez_OpenCard) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{5} }

func (m *Pez_OpenCard) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_OpenCard) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

type Pez_AckOpenCard struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Result           *int32             `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	PaiType          *int32             `protobuf:"varint,3,opt,name=paiType" json:"paiType,omitempty"`
	UserId           *uint32            `protobuf:"varint,4,opt,name=userId" json:"userId,omitempty"`
	Card             *PezBase_PaiInfo   `protobuf:"bytes,5,opt,name=card" json:"card,omitempty"`
	Card2            *PezBase_PaiInfo   `protobuf:"bytes,6,opt,name=card2" json:"card2,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,7,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_AckOpenCard) Reset()                    { *m = Pez_AckOpenCard{} }
func (m *Pez_AckOpenCard) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckOpenCard) ProtoMessage()               {}
func (*Pez_AckOpenCard) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{6} }

func (m *Pez_AckOpenCard) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckOpenCard) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *Pez_AckOpenCard) GetPaiType() int32 {
	if m != nil && m.PaiType != nil {
		return *m.PaiType
	}
	return 0
}

func (m *Pez_AckOpenCard) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckOpenCard) GetCard() *PezBase_PaiInfo {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *Pez_AckOpenCard) GetCard2() *PezBase_PaiInfo {
	if m != nil {
		return m.Card2
	}
	return nil
}

func (m *Pez_AckOpenCard) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

type Pez_ActCompare struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ComparedUserId   *uint32      `protobuf:"varint,2,opt,name=comparedUserId" json:"comparedUserId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_ActCompare) Reset()                    { *m = Pez_ActCompare{} }
func (m *Pez_ActCompare) String() string            { return proto.CompactTextString(m) }
func (*Pez_ActCompare) ProtoMessage()               {}
func (*Pez_ActCompare) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{7} }

func (m *Pez_ActCompare) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_ActCompare) GetComparedUserId() uint32 {
	if m != nil && m.ComparedUserId != nil {
		return *m.ComparedUserId
	}
	return 0
}

type Pez_AckActCompare struct {
	Header           *ProtoHeader            `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ComparedUserId   *uint32                 `protobuf:"varint,2,opt,name=comparedUserId" json:"comparedUserId,omitempty"`
	IsCpareWin       *bool                   `protobuf:"varint,3,opt,name=isCpareWin" json:"isCpareWin,omitempty"`
	PlayerInfo       *PezBase_PlayerInfo     `protobuf:"bytes,4,opt,name=playerInfo" json:"playerInfo,omitempty"`
	Deskstate        *PezEnum_DeskGameStatus `protobuf:"varint,5,opt,name=deskstate,enum=ddproto.PezEnum_DeskGameStatus" json:"deskstate,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Pez_AckActCompare) Reset()                    { *m = Pez_AckActCompare{} }
func (m *Pez_AckActCompare) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckActCompare) ProtoMessage()               {}
func (*Pez_AckActCompare) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{8} }

func (m *Pez_AckActCompare) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckActCompare) GetComparedUserId() uint32 {
	if m != nil && m.ComparedUserId != nil {
		return *m.ComparedUserId
	}
	return 0
}

func (m *Pez_AckActCompare) GetIsCpareWin() bool {
	if m != nil && m.IsCpareWin != nil {
		return *m.IsCpareWin
	}
	return false
}

func (m *Pez_AckActCompare) GetPlayerInfo() *PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Pez_AckActCompare) GetDeskstate() PezEnum_DeskGameStatus {
	if m != nil && m.Deskstate != nil {
		return *m.Deskstate
	}
	return PezEnum_DeskGameStatus_PEZ_INIT
}

// 发送游戏信息(广播)
type Pez_SendGameInfo struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 1. 首先是牌桌的玩家数据（玩家数据包括其id昵称筹码头像等基本信息，其手牌数据，以及自己打开牌的数据，还有状态是否已经押注了，玩家在整局的总输赢）
	PlayerInfo []*PezBase_PlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	// 2. 桌面信息（包括：游戏是否结束，当前哪个玩家未押注，倒计时剩余时间）
	DeskGameInfo *PezBase_DeskGameInfo `protobuf:"bytes,3,opt,name=deskGameInfo" json:"deskGameInfo,omitempty"`
	//
	SenderUserId     *uint32             `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *PEZ_RECONNECT_TYPE `protobuf:"varint,5,opt,name=isReconnect,enum=ddproto.PEZ_RECONNECT_TYPE" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Pez_SendGameInfo) Reset()                    { *m = Pez_SendGameInfo{} }
func (m *Pez_SendGameInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendGameInfo) ProtoMessage()               {}
func (*Pez_SendGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor36, []int{9} }

func (m *Pez_SendGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_SendGameInfo) GetPlayerInfo() []*PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetDeskGameInfo() *PezBase_DeskGameInfo {
	if m != nil {
		return m.DeskGameInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *Pez_SendGameInfo) GetIsReconnect() PEZ_RECONNECT_TYPE {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return PEZ_RECONNECT_TYPE_PEZ_NORMAL
}

func init() {
	proto.RegisterType((*PezUserCoinBean)(nil), "ddproto.pez_user_coin_bean")
	proto.RegisterType((*Pez_Opening)(nil), "ddproto.pez_Opening")
	proto.RegisterType((*Pez_DealCards)(nil), "ddproto.pez_DealCards")
	proto.RegisterType((*Pez_Bet)(nil), "ddproto.pez_Bet")
	proto.RegisterType((*Pez_AckBet)(nil), "ddproto.pez_AckBet")
	proto.RegisterType((*Pez_OpenCard)(nil), "ddproto.pez_OpenCard")
	proto.RegisterType((*Pez_AckOpenCard)(nil), "ddproto.pez_AckOpenCard")
	proto.RegisterType((*Pez_ActCompare)(nil), "ddproto.pez_ActCompare")
	proto.RegisterType((*Pez_AckActCompare)(nil), "ddproto.pez_AckActCompare")
	proto.RegisterType((*Pez_SendGameInfo)(nil), "ddproto.pez_SendGameInfo")
	proto.RegisterEnum("ddproto.PEZ_RECONNECT_TYPE", PEZ_RECONNECT_TYPE_name, PEZ_RECONNECT_TYPE_value)
}

<<<<<<< HEAD
var fileDescriptor36 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0x96, 0x69, 0x9d, 0xa6, 0x4b, 0x41, 0xa6, 0x20, 0x54, 0x59, 0x48, 0x54,
	0x1c, 0x2a, 0x30, 0x48, 0xbd, 0x70, 0x49, 0x9c, 0x88, 0x46, 0x82, 0xc4, 0x4a, 0x83, 0x10, 0x5c,
	0xac, 0x8d, 0xbd, 0x80, 0x95, 0x78, 0x6d, 0x79, 0xed, 0x43, 0xf9, 0x09, 0x7e, 0x80, 0x4f, 0xe2,
	0x43, 0xb8, 0xf1, 0x0b, 0xcc, 0xac, 0xdd, 0x36, 0x51, 0x2a, 0x92, 0x4a, 0x5c, 0xa2, 0xec, 0xec,
	0xbc, 0x79, 0x6f, 0xde, 0xce, 0x18, 0x9a, 0xa9, 0xf8, 0xee, 0xa7, 0x33, 0x7e, 0x71, 0x92, 0x66,
	0x49, 0x9e, 0xb0, 0xad, 0x30, 0xd4, 0x7f, 0x0e, 0xef, 0x05, 0x49, 0x1c, 0x27, 0xd2, 0x0f, 0x66,
	0x91, 0x90, 0x79, 0x79, 0x7b, 0xa8, 0xb3, 0x27, 0x5c, 0x89, 0xea, 0xbc, 0x5f, 0x25, 0x09, 0x59,
	0xc4, 0x65, 0xc8, 0x76, 0x80, 0x51, 0x52, 0xa1, 0x44, 0xe6, 0x07, 0x49, 0x24, 0xfd, 0x89, 0xe0,
	0x92, 0x35, 0x61, 0x93, 0x22, 0xfd, 0xd0, 0x32, 0x8e, 0x8c, 0x63, 0x93, 0xed, 0xc2, 0x06, 0x5d,
	0x5a, 0x35, 0x3c, 0xd5, 0xed, 0x9f, 0x06, 0xec, 0x10, 0x68, 0x98, 0x0a, 0x19, 0xc9, 0xaf, 0xec,
	0x29, 0x6c, 0x7e, 0x13, 0x3c, 0x14, 0x99, 0xce, 0xde, 0x71, 0x0e, 0x4e, 0x2a, 0x55, 0x27, 0x1e,
	0xfd, 0x9e, 0xe9, 0x3b, 0x76, 0x1f, 0x4c, 0xb7, 0xc8, 0x32, 0x0f, 0xc5, 0xbb, 0x49, 0x21, 0x73,
	0x5d, 0xac, 0xc1, 0x4c, 0x68, 0x84, 0x51, 0x20, 0x5e, 0x5a, 0xf5, 0xf9, 0xa3, 0x63, 0x6d, 0xe8,
	0xa3, 0x03, 0x26, 0x09, 0x71, 0x91, 0xbc, 0x83, 0xc2, 0x94, 0xd5, 0x38, 0xaa, 0x23, 0xc3, 0xa3,
	0x2b, 0x86, 0x65, 0xf1, 0xf6, 0x0f, 0x03, 0x4c, 0x0a, 0x77, 0x05, 0x9f, 0xb9, 0x3c, 0x0b, 0xd5,
	0x9a, 0x02, 0x5f, 0x00, 0x90, 0xb3, 0xc8, 0x86, 0x20, 0x54, 0x47, 0x44, 0x8f, 0x17, 0x88, 0xc8,
	0x4a, 0xdf, 0xbb, 0xca, 0x61, 0x2d, 0xd8, 0x4e, 0x79, 0x54, 0x76, 0x53, 0xca, 0x3f, 0x80, 0xdd,
	0x10, 0x69, 0x45, 0xf6, 0xa1, 0xb4, 0x8f, 0xba, 0x30, 0xed, 0x53, 0xd8, 0x22, 0x78, 0x47, 0xe4,
	0x4b, 0xce, 0xe2, 0x79, 0x22, 0xf2, 0x41, 0x11, 0x57, 0x76, 0xa0, 0xd3, 0x79, 0x14, 0x8b, 0xb2,
	0x9c, 0xfd, 0x06, 0x25, 0x21, 0xb0, 0x1d, 0x4c, 0x6f, 0xc2, 0x22, 0x3d, 0x62, 0xe7, 0xcd, 0x5c,
	0x44, 0x77, 0x61, 0xf7, 0xf2, 0x99, 0xb4, 0xdc, 0xf5, 0x6c, 0x40, 0x96, 0x00, 0xb3, 0xfb, 0x61,
	0x59, 0xd3, 0xfe, 0x63, 0xc0, 0x5e, 0x25, 0xe2, 0xf6, 0x95, 0x32, 0xa1, 0x8a, 0xd9, 0xa5, 0xba,
	0x3d, 0xb4, 0x81, 0x47, 0xe3, 0x8b, 0xb4, 0x12, 0x38, 0xd7, 0x90, 0xf6, 0x89, 0x3d, 0xc3, 0x31,
	0x23, 0xef, 0x1b, 0xba, 0xe8, 0xc3, 0x1b, 0xbc, 0xe7, 0x51, 0x5f, 0x7e, 0x49, 0xd8, 0x31, 0x34,
	0x28, 0xd1, 0xb1, 0x36, 0x57, 0x65, 0x2e, 0x0d, 0xd0, 0xd6, 0xea, 0x01, 0x1a, 0x94, 0x6b, 0xd6,
	0x0e, 0xd0, 0xdb, 0x38, 0xe5, 0x99, 0x58, 0xb3, 0xdf, 0x07, 0xd0, 0x0c, 0x4a, 0x40, 0x58, 0x3d,
	0x7f, 0x4d, 0x3f, 0xff, 0x2f, 0x03, 0xf6, 0x2b, 0x07, 0xff, 0x57, 0x4d, 0xc6, 0x00, 0x22, 0xe5,
	0x52, 0xfc, 0x23, 0xee, 0x25, 0xd9, 0xb9, 0x7d, 0x3d, 0xc0, 0xd4, 0xb9, 0xb6, 0xf4, 0x1f, 0x03,
	0xac, 0xdd, 0x79, 0x05, 0x77, 0x43, 0xa1, 0xa6, 0x2a, 0xe7, 0xb9, 0xd0, 0xae, 0x37, 0x9d, 0xa3,
	0x05, 0x00, 0x7d, 0x29, 0x70, 0x91, 0xd4, 0xf4, 0x2d, 0x8f, 0xc5, 0x39, 0xa6, 0x15, 0xca, 0xfe,
	0x6d, 0x40, 0x8b, 0xee, 0xce, 0x85, 0x0c, 0x29, 0xac, 0x2b, 0xdd, 0x72, 0xc5, 0xb4, 0xc2, 0x15,
	0x2b, 0xa6, 0xeb, 0xbe, 0xa6, 0x85, 0x2a, 0xe9, 0x35, 0xa6, 0xae, 0xab, 0x3f, 0x59, 0xc6, 0x74,
	0xe7, 0xb2, 0x68, 0x0d, 0x15, 0xaa, 0x5b, 0x5c, 0x43, 0x64, 0xdf, 0x89, 0xd4, 0x48, 0x04, 0x89,
	0x94, 0x22, 0xc8, 0xab, 0x7e, 0xaf, 0x27, 0xc1, 0xeb, 0x7d, 0xf6, 0x47, 0x3d, 0x77, 0x38, 0x18,
	0xf4, 0xdc, 0xb1, 0x3f, 0xfe, 0xe4, 0xf5, 0x9e, 0x9f, 0x02, 0x5b, 0x8e, 0xe2, 0xd8, 0x02, 0x45,
	0x07, 0xc3, 0xd1, 0xfb, 0xf6, 0xbb, 0x96, 0xc1, 0xf6, 0xc1, 0x5c, 0xc8, 0x6a, 0xd5, 0x3a, 0xb5,
	0xb3, 0xba, 0x77, 0xc7, 0x33, 0xbc, 0xda, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x91, 0x28,
	0xd0, 0xb2, 0x05, 0x00, 0x00,
=======
var fileDescriptor39 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0xce, 0xab, 0xdc, 0x34, 0xa1, 0x1d, 0xaa, 0xca, 0x04, 0x54, 0x45, 0x16, 0x42, 0x15,
	0x82, 0x22, 0xb2, 0x61, 0x03, 0x88, 0xc4, 0x8d, 0x68, 0x25, 0x9a, 0x5a, 0xd3, 0x20, 0x04, 0x1b,
	0x6b, 0x62, 0xdf, 0x82, 0xd5, 0x78, 0x6c, 0xf9, 0xb1, 0x08, 0x3f, 0xc0, 0x97, 0xb0, 0xe3, 0x0f,
	0xf8, 0x09, 0x3e, 0x09, 0xcd, 0x78, 0x92, 0xd8, 0x0d, 0x6d, 0x95, 0x4a, 0x6c, 0x92, 0xb9, 0x67,
	0xce, 0x7d, 0x9e, 0xf1, 0x85, 0x76, 0x84, 0xdf, 0x9d, 0x68, 0xca, 0x66, 0x07, 0x51, 0x1c, 0xa6,
	0x21, 0x69, 0x78, 0x9e, 0x3c, 0x74, 0xee, 0xbb, 0x61, 0x10, 0x84, 0xdc, 0x71, 0xa7, 0x3e, 0xf2,
	0x34, 0xbf, 0xed, 0x48, 0xf6, 0x84, 0x25, 0xa8, 0xec, 0x6d, 0x45, 0x42, 0x9e, 0x05, 0x39, 0x64,
	0xbe, 0x03, 0x22, 0x48, 0x59, 0x82, 0xb1, 0xe3, 0x86, 0x3e, 0x77, 0x26, 0xc8, 0x38, 0xd9, 0x85,
	0xba, 0x40, 0x8e, 0x3d, 0x43, 0xeb, 0x6a, 0xfb, 0x2d, 0xaa, 0x2c, 0x42, 0xa0, 0x2a, 0x48, 0x86,
	0xde, 0xd5, 0xf6, 0x2b, 0x54, 0x9e, 0xcd, 0x3f, 0x1a, 0x34, 0x45, 0x88, 0xd3, 0x08, 0xb9, 0xcf,
	0xbf, 0x92, 0x67, 0x50, 0xff, 0x86, 0xcc, 0xc3, 0x58, 0xfa, 0x36, 0x7b, 0x3b, 0x07, 0xaa, 0xc6,
	0x03, 0x5b, 0xfc, 0x1e, 0xc9, 0x3b, 0xaa, 0x38, 0xe4, 0x31, 0xb4, 0xac, 0x2c, 0x8e, 0xed, 0x29,
	0x9b, 0x59, 0x61, 0xc6, 0x53, 0x19, 0xba, 0x46, 0xcb, 0x20, 0xd9, 0x81, 0x9a, 0xe7, 0xbb, 0xf8,
	0xd2, 0xa8, 0xc8, 0xdb, 0xdc, 0x98, 0xa3, 0x3d, 0xa3, 0xba, 0x44, 0x7b, 0xa4, 0x0f, 0x2d, 0x51,
	0xad, 0x15, 0xfa, 0x7c, 0x80, 0x8c, 0x27, 0x46, 0xad, 0x5b, 0xd9, 0x6f, 0xf6, 0x1e, 0x2e, 0xca,
	0x58, 0xed, 0x97, 0x96, 0x3d, 0xcc, 0xdf, 0x1a, 0xb4, 0x04, 0xeb, 0x10, 0xd9, 0xd4, 0x62, 0xb1,
	0x97, 0xac, 0xd9, 0xd4, 0x6b, 0x00, 0xa1, 0x11, 0xc6, 0xc2, 0xd9, 0xd0, 0x65, 0xfe, 0x47, 0xa5,
	0xfc, 0x42, 0x14, 0xc7, 0x5e, 0x70, 0x68, 0x81, 0x4f, 0x3a, 0xb0, 0x11, 0x31, 0x3f, 0x9f, 0x46,
	0xde, 0xef, 0xc2, 0x26, 0x26, 0x6c, 0x7a, 0xc8, 0xa6, 0x18, 0x7f, 0xcc, 0xe5, 0xa9, 0x4a, 0x79,
	0x4a, 0x98, 0x79, 0x02, 0x0d, 0x91, 0x62, 0x80, 0xe9, 0x95, 0x3a, 0xee, 0x42, 0x7d, 0x82, 0xe9,
	0x28, 0x0b, 0xd4, 0xb8, 0x95, 0x25, 0xf4, 0x4d, 0xfd, 0x00, 0x55, 0x5a, 0x79, 0x36, 0xc7, 0x00,
	0x22, 0x5c, 0xdf, 0xbd, 0xb8, 0x2e, 0x62, 0x07, 0x36, 0x26, 0x98, 0x16, 0x25, 0x5c, 0xd8, 0x57,
	0x44, 0xdd, 0x9c, 0x3f, 0x1a, 0xd9, 0xf4, 0x7a, 0x03, 0xde, 0x85, 0xba, 0xcb, 0x62, 0xef, 0xd8,
	0x9b, 0xd7, 0x9f, 0x5b, 0xe6, 0x2f, 0x1d, 0xee, 0xa9, 0x62, 0x6f, 0x1f, 0x39, 0xc6, 0x24, 0x9b,
	0xce, 0xbb, 0x50, 0x16, 0x31, 0xa0, 0x11, 0x31, 0x7f, 0x3c, 0x8b, 0xe6, 0x6d, 0xcc, 0xcd, 0xc2,
	0x44, 0xaa, 0xa5, 0x89, 0x3c, 0x87, 0xaa, 0xa8, 0xca, 0xa8, 0xc9, 0xac, 0x0f, 0xfe, 0x21, 0x3f,
	0xf3, 0x8f, 0xf9, 0x79, 0x48, 0x25, 0x8d, 0xbc, 0x80, 0x9a, 0xf8, 0xef, 0x19, 0xf5, 0x9b, 0xf8,
	0x39, 0x6f, 0xf5, 0x9d, 0x37, 0xd6, 0x7e, 0xe7, 0xe7, 0xf9, 0x3e, 0xe9, 0xbb, 0xa9, 0x15, 0x06,
	0x11, 0x8b, 0x71, 0xcd, 0x61, 0x3d, 0x81, 0xb6, 0x9b, 0x3b, 0x7a, 0xea, 0x3d, 0xea, 0x72, 0x04,
	0x97, 0x50, 0xf3, 0x87, 0x0e, 0xdb, 0x4a, 0x96, 0xff, 0x9d, 0x8b, 0xec, 0x01, 0xf8, 0x89, 0x25,
	0x90, 0x4f, 0x3e, 0x97, 0x5a, 0x6d, 0xd0, 0x02, 0xb2, 0xfc, 0x36, 0xc5, 0x2c, 0xa5, 0x64, 0xd7,
	0x7c, 0x9b, 0x72, 0xde, 0x05, 0x3e, 0x79, 0x0b, 0x77, 0x3d, 0x4c, 0x2e, 0x92, 0x94, 0xa5, 0x28,
	0x95, 0x6d, 0xf7, 0xba, 0x25, 0x67, 0xb1, 0x5a, 0x9d, 0x43, 0x4c, 0x2e, 0xde, 0xb3, 0x00, 0xcf,
	0x52, 0x96, 0x66, 0x09, 0x5d, 0xba, 0x98, 0x3f, 0x75, 0xd8, 0x12, 0xb4, 0x33, 0xe4, 0x9e, 0x60,
	0xc8, 0xa0, 0xb7, 0x5c, 0x2e, 0xb2, 0x81, 0x1b, 0x96, 0xcb, 0x4a, 0x03, 0x03, 0xb1, 0x40, 0xf2,
	0xea, 0xa4, 0x7f, 0x45, 0x66, 0xdc, 0x5b, 0xf5, 0x3f, 0x2c, 0xb0, 0x68, 0xc9, 0x47, 0x2c, 0xa1,
	0x04, 0xb9, 0x77, 0x79, 0x09, 0x15, 0x31, 0xf2, 0x06, 0x9a, 0x7e, 0x42, 0xd1, 0x0d, 0x39, 0x47,
	0x37, 0x55, 0xa3, 0x5a, 0xbe, 0x4d, 0x7b, 0xf8, 0xc5, 0xa1, 0x43, 0xeb, 0x74, 0x34, 0x1a, 0x5a,
	0x63, 0x67, 0xfc, 0xd9, 0x1e, 0xd2, 0x22, 0xff, 0xe9, 0x2b, 0x20, 0xab, 0x14, 0xd2, 0x06, 0x10,
	0xe8, 0xe8, 0x94, 0x9e, 0xf4, 0x3f, 0x6c, 0x69, 0x64, 0x1b, 0x5a, 0x25, 0xd6, 0x96, 0x3e, 0xd0,
	0x8f, 0x2a, 0xf6, 0x1d, 0x5b, 0xb3, 0xf5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x6e, 0x66,
	0xbe, 0x2b, 0x07, 0x00, 0x00,
>>>>>>> 909db0f792448fe2ab44876c4ce012340a336faf
}
