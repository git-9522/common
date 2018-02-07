// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zhadan_base.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_req_list_coin_info from common_client.proto

// Ignoring public import of common_ack_list_coin_info from common_client.proto

// Ignoring public import of CommonCoinLevelInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type ZhadanEnumProtoid int32

const (
	// //////////////////////////////////
	ZhadanEnumProtoid_ZHADAN_PID_HEARTBEAT          ZhadanEnumProtoid = 0
	ZhadanEnumProtoid_ZHADAN_PID_CREATE_DESK_REQ    ZhadanEnumProtoid = 1
	ZhadanEnumProtoid_ZHADAN_PID_ENTER_DESK_REQ     ZhadanEnumProtoid = 2
	ZhadanEnumProtoid_ZHADAN_PID_ENTER_DESK_ACK     ZhadanEnumProtoid = 3
	ZhadanEnumProtoid_ZHADAN_PID_ENTER_DESK_BC      ZhadanEnumProtoid = 4
	ZhadanEnumProtoid_ZHADAN_PID_READY_REQ          ZhadanEnumProtoid = 5
	ZhadanEnumProtoid_ZHADAN_PID_READY_BC           ZhadanEnumProtoid = 6
	ZhadanEnumProtoid_ZHADAN_PID_FAPAI_OT           ZhadanEnumProtoid = 7
	ZhadanEnumProtoid_ZHADAN_PID_BAOZHUANG_OT       ZhadanEnumProtoid = 8
	ZhadanEnumProtoid_ZHADAN_PID_BAOZHUANG_REQ      ZhadanEnumProtoid = 9
	ZhadanEnumProtoid_ZHADAN_PID_BAOZHUANG_ACK      ZhadanEnumProtoid = 10
	ZhadanEnumProtoid_ZHADAN_PID_BAOZHUANG_RES      ZhadanEnumProtoid = 11
	ZhadanEnumProtoid_ZHADAN_PID_CHUPAI_OT          ZhadanEnumProtoid = 12
	ZhadanEnumProtoid_ZHADAN_PID_CHUPAI_REQ         ZhadanEnumProtoid = 13
	ZhadanEnumProtoid_ZHADAN_PID_CHUPAI_BC          ZhadanEnumProtoid = 14
	ZhadanEnumProtoid_ZHADAN_PID_PASS_REQ           ZhadanEnumProtoid = 15
	ZhadanEnumProtoid_ZHADAN_PID_PASS_BC            ZhadanEnumProtoid = 16
	ZhadanEnumProtoid_ZHADAN_PID_GAME_END_ONE_BC    ZhadanEnumProtoid = 17
	ZhadanEnumProtoid_ZHADAN_PID_GAME_END_ALL_BC    ZhadanEnumProtoid = 18
	ZhadanEnumProtoid_ZHADAN_PID_APPLY_DISSOLVE_REQ ZhadanEnumProtoid = 19
	ZhadanEnumProtoid_ZHADAN_PID_APPLY_DISSOLVE_ACK ZhadanEnumProtoid = 20
	ZhadanEnumProtoid_ZHADAN_PID_DISSOLVE_BACK_REQ  ZhadanEnumProtoid = 21
	ZhadanEnumProtoid_ZHADAN_PID_DISSOLVE_BACK_ACK  ZhadanEnumProtoid = 22
	ZhadanEnumProtoid_ZHADAN_PID_OWNER_DISSOLVE_REQ ZhadanEnumProtoid = 23
	ZhadanEnumProtoid_ZHADAN_PID_OWNER_DISSOLVE_ACK ZhadanEnumProtoid = 24
	ZhadanEnumProtoid_ZHADAN_PID_SEND_MESSAGE_REQ   ZhadanEnumProtoid = 25
	ZhadanEnumProtoid_ZHADAN_PID_MESSAGE_BC         ZhadanEnumProtoid = 26
	ZhadanEnumProtoid_ZHADAN_PID_LEAVE_DESK_REQ     ZhadanEnumProtoid = 27
	ZhadanEnumProtoid_ZHADAN_PID_LEAVE_DESK_ACK     ZhadanEnumProtoid = 28
	ZhadanEnumProtoid_ZHADAN_PID_OFFLINE_BC         ZhadanEnumProtoid = 29
	ZhadanEnumProtoid_ZHADAN_PID_COIN_ROOM_LIST_REQ ZhadanEnumProtoid = 30
	ZhadanEnumProtoid_ZHADAN_PID_COIN_ROOM_LIST_ACK ZhadanEnumProtoid = 31
)

var ZhadanEnumProtoid_name = map[int32]string{
	0:  "ZHADAN_PID_HEARTBEAT",
	1:  "ZHADAN_PID_CREATE_DESK_REQ",
	2:  "ZHADAN_PID_ENTER_DESK_REQ",
	3:  "ZHADAN_PID_ENTER_DESK_ACK",
	4:  "ZHADAN_PID_ENTER_DESK_BC",
	5:  "ZHADAN_PID_READY_REQ",
	6:  "ZHADAN_PID_READY_BC",
	7:  "ZHADAN_PID_FAPAI_OT",
	8:  "ZHADAN_PID_BAOZHUANG_OT",
	9:  "ZHADAN_PID_BAOZHUANG_REQ",
	10: "ZHADAN_PID_BAOZHUANG_ACK",
	11: "ZHADAN_PID_BAOZHUANG_RES",
	12: "ZHADAN_PID_CHUPAI_OT",
	13: "ZHADAN_PID_CHUPAI_REQ",
	14: "ZHADAN_PID_CHUPAI_BC",
	15: "ZHADAN_PID_PASS_REQ",
	16: "ZHADAN_PID_PASS_BC",
	17: "ZHADAN_PID_GAME_END_ONE_BC",
	18: "ZHADAN_PID_GAME_END_ALL_BC",
	19: "ZHADAN_PID_APPLY_DISSOLVE_REQ",
	20: "ZHADAN_PID_APPLY_DISSOLVE_ACK",
	21: "ZHADAN_PID_DISSOLVE_BACK_REQ",
	22: "ZHADAN_PID_DISSOLVE_BACK_ACK",
	23: "ZHADAN_PID_OWNER_DISSOLVE_REQ",
	24: "ZHADAN_PID_OWNER_DISSOLVE_ACK",
	25: "ZHADAN_PID_SEND_MESSAGE_REQ",
	26: "ZHADAN_PID_MESSAGE_BC",
	27: "ZHADAN_PID_LEAVE_DESK_REQ",
	28: "ZHADAN_PID_LEAVE_DESK_ACK",
	29: "ZHADAN_PID_OFFLINE_BC",
	30: "ZHADAN_PID_COIN_ROOM_LIST_REQ",
	31: "ZHADAN_PID_COIN_ROOM_LIST_ACK",
}
var ZhadanEnumProtoid_value = map[string]int32{
	"ZHADAN_PID_HEARTBEAT":          0,
	"ZHADAN_PID_CREATE_DESK_REQ":    1,
	"ZHADAN_PID_ENTER_DESK_REQ":     2,
	"ZHADAN_PID_ENTER_DESK_ACK":     3,
	"ZHADAN_PID_ENTER_DESK_BC":      4,
	"ZHADAN_PID_READY_REQ":          5,
	"ZHADAN_PID_READY_BC":           6,
	"ZHADAN_PID_FAPAI_OT":           7,
	"ZHADAN_PID_BAOZHUANG_OT":       8,
	"ZHADAN_PID_BAOZHUANG_REQ":      9,
	"ZHADAN_PID_BAOZHUANG_ACK":      10,
	"ZHADAN_PID_BAOZHUANG_RES":      11,
	"ZHADAN_PID_CHUPAI_OT":          12,
	"ZHADAN_PID_CHUPAI_REQ":         13,
	"ZHADAN_PID_CHUPAI_BC":          14,
	"ZHADAN_PID_PASS_REQ":           15,
	"ZHADAN_PID_PASS_BC":            16,
	"ZHADAN_PID_GAME_END_ONE_BC":    17,
	"ZHADAN_PID_GAME_END_ALL_BC":    18,
	"ZHADAN_PID_APPLY_DISSOLVE_REQ": 19,
	"ZHADAN_PID_APPLY_DISSOLVE_ACK": 20,
	"ZHADAN_PID_DISSOLVE_BACK_REQ":  21,
	"ZHADAN_PID_DISSOLVE_BACK_ACK":  22,
	"ZHADAN_PID_OWNER_DISSOLVE_REQ": 23,
	"ZHADAN_PID_OWNER_DISSOLVE_ACK": 24,
	"ZHADAN_PID_SEND_MESSAGE_REQ":   25,
	"ZHADAN_PID_MESSAGE_BC":         26,
	"ZHADAN_PID_LEAVE_DESK_REQ":     27,
	"ZHADAN_PID_LEAVE_DESK_ACK":     28,
	"ZHADAN_PID_OFFLINE_BC":         29,
	"ZHADAN_PID_COIN_ROOM_LIST_REQ": 30,
	"ZHADAN_PID_COIN_ROOM_LIST_ACK": 31,
}

func (x ZhadanEnumProtoid) Enum() *ZhadanEnumProtoid {
	p := new(ZhadanEnumProtoid)
	*p = x
	return p
}
func (x ZhadanEnumProtoid) String() string {
	return proto.EnumName(ZhadanEnumProtoid_name, int32(x))
}
func (x *ZhadanEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZhadanEnumProtoid_value, data, "ZhadanEnumProtoid")
	if err != nil {
		return err
	}
	*x = ZhadanEnumProtoid(value)
	return nil
}
func (ZhadanEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor57, []int{0} }

// =================================公共================================
// 牌型
type ZhadanEnumPokerType int32

const (
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_OTHER    ZhadanEnumPokerType = 0
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_DAN_PAI  ZhadanEnumPokerType = 1
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_DUI_PAI  ZhadanEnumPokerType = 2
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_SAN_TIAO ZhadanEnumPokerType = 3
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_SHUNZI   ZhadanEnumPokerType = 4
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_LIANDUI  ZhadanEnumPokerType = 5
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_FEIJI    ZhadanEnumPokerType = 6
	ZhadanEnumPokerType_ZHADAN_POKER_TYPE_ZHADAN   ZhadanEnumPokerType = 7
)

var ZhadanEnumPokerType_name = map[int32]string{
	0: "ZHADAN_POKER_TYPE_OTHER",
	1: "ZHADAN_POKER_TYPE_DAN_PAI",
	2: "ZHADAN_POKER_TYPE_DUI_PAI",
	3: "ZHADAN_POKER_TYPE_SAN_TIAO",
	4: "ZHADAN_POKER_TYPE_SHUNZI",
	5: "ZHADAN_POKER_TYPE_LIANDUI",
	6: "ZHADAN_POKER_TYPE_FEIJI",
	7: "ZHADAN_POKER_TYPE_ZHADAN",
}
var ZhadanEnumPokerType_value = map[string]int32{
	"ZHADAN_POKER_TYPE_OTHER":    0,
	"ZHADAN_POKER_TYPE_DAN_PAI":  1,
	"ZHADAN_POKER_TYPE_DUI_PAI":  2,
	"ZHADAN_POKER_TYPE_SAN_TIAO": 3,
	"ZHADAN_POKER_TYPE_SHUNZI":   4,
	"ZHADAN_POKER_TYPE_LIANDUI":  5,
	"ZHADAN_POKER_TYPE_FEIJI":    6,
	"ZHADAN_POKER_TYPE_ZHADAN":   7,
}

func (x ZhadanEnumPokerType) Enum() *ZhadanEnumPokerType {
	p := new(ZhadanEnumPokerType)
	*p = x
	return p
}
func (x ZhadanEnumPokerType) String() string {
	return proto.EnumName(ZhadanEnumPokerType_name, int32(x))
}
func (x *ZhadanEnumPokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZhadanEnumPokerType_value, data, "ZhadanEnumPokerType")
	if err != nil {
		return err
	}
	*x = ZhadanEnumPokerType(value)
	return nil
}
func (ZhadanEnumPokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor57, []int{1} }

// 房间状态
type ZhadanEnumDeskStatus int32

const (
	ZhadanEnumDeskStatus_ZHADAN_DESK_STATUS_WAIT_READY     ZhadanEnumDeskStatus = 1
	ZhadanEnumDeskStatus_ZHADAN_DESK_STATUS_WAIT_BAOZHUANG ZhadanEnumDeskStatus = 2
	ZhadanEnumDeskStatus_ZHADAN_DESK_STATUS_ON_GAMMING     ZhadanEnumDeskStatus = 3
)

var ZhadanEnumDeskStatus_name = map[int32]string{
	1: "ZHADAN_DESK_STATUS_WAIT_READY",
	2: "ZHADAN_DESK_STATUS_WAIT_BAOZHUANG",
	3: "ZHADAN_DESK_STATUS_ON_GAMMING",
}
var ZhadanEnumDeskStatus_value = map[string]int32{
	"ZHADAN_DESK_STATUS_WAIT_READY":     1,
	"ZHADAN_DESK_STATUS_WAIT_BAOZHUANG": 2,
	"ZHADAN_DESK_STATUS_ON_GAMMING":     3,
}

func (x ZhadanEnumDeskStatus) Enum() *ZhadanEnumDeskStatus {
	p := new(ZhadanEnumDeskStatus)
	*p = x
	return p
}
func (x ZhadanEnumDeskStatus) String() string {
	return proto.EnumName(ZhadanEnumDeskStatus_name, int32(x))
}
func (x *ZhadanEnumDeskStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZhadanEnumDeskStatus_value, data, "ZhadanEnumDeskStatus")
	if err != nil {
		return err
	}
	*x = ZhadanEnumDeskStatus(value)
	return nil
}
func (ZhadanEnumDeskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor57, []int{2} }

type ZhadanGanOpt int32

const (
	ZhadanGanOpt_ZHADAN_GAN_OPT_BANGAN  ZhadanGanOpt = 1
	ZhadanGanOpt_ZHADAN_GAN_OPT_QUANGAN ZhadanGanOpt = 2
)

var ZhadanGanOpt_name = map[int32]string{
	1: "ZHADAN_GAN_OPT_BANGAN",
	2: "ZHADAN_GAN_OPT_QUANGAN",
}
var ZhadanGanOpt_value = map[string]int32{
	"ZHADAN_GAN_OPT_BANGAN":  1,
	"ZHADAN_GAN_OPT_QUANGAN": 2,
}

func (x ZhadanGanOpt) Enum() *ZhadanGanOpt {
	p := new(ZhadanGanOpt)
	*p = x
	return p
}
func (x ZhadanGanOpt) String() string {
	return proto.EnumName(ZhadanGanOpt_name, int32(x))
}
func (x *ZhadanGanOpt) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZhadanGanOpt_value, data, "ZhadanGanOpt")
	if err != nil {
		return err
	}
	*x = ZhadanGanOpt(value)
	return nil
}
func (ZhadanGanOpt) EnumDescriptor() ([]byte, []int) { return fileDescriptor57, []int{3} }

// 打出去的牌
type ZhadanClientPoker struct {
	Pais             []int32              `protobuf:"varint,2,rep,name=pais" json:"pais,omitempty"`
	PokerType        *ZhadanEnumPokerType `protobuf:"varint,3,opt,name=pokerType,enum=ddproto.ZhadanEnumPokerType" json:"pokerType,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZhadanClientPoker) Reset()                    { *m = ZhadanClientPoker{} }
func (m *ZhadanClientPoker) String() string            { return proto.CompactTextString(m) }
func (*ZhadanClientPoker) ProtoMessage()               {}
func (*ZhadanClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor57, []int{0} }

func (m *ZhadanClientPoker) GetPais() []int32 {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZhadanClientPoker) GetPokerType() ZhadanEnumPokerType {
	if m != nil && m.PokerType != nil {
		return *m.PokerType
	}
	return ZhadanEnumPokerType_ZHADAN_POKER_TYPE_OTHER
}

// 对局账单信息
type ZhadanUserBill struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	AllBaseScore     *int32  `protobuf:"varint,3,opt,name=allBaseScore" json:"allBaseScore,omitempty"`
	AllExtScore      *int32  `protobuf:"varint,4,opt,name=allExtScore" json:"allExtScore,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZhadanUserBill) Reset()                    { *m = ZhadanUserBill{} }
func (m *ZhadanUserBill) String() string            { return proto.CompactTextString(m) }
func (*ZhadanUserBill) ProtoMessage()               {}
func (*ZhadanUserBill) Descriptor() ([]byte, []int) { return fileDescriptor57, []int{1} }

func (m *ZhadanUserBill) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZhadanUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *ZhadanUserBill) GetAllBaseScore() int32 {
	if m != nil && m.AllBaseScore != nil {
		return *m.AllBaseScore
	}
	return 0
}

func (m *ZhadanUserBill) GetAllExtScore() int32 {
	if m != nil && m.AllExtScore != nil {
		return *m.AllExtScore
	}
	return 0
}

// desk 配置选项
type ZhadanDeskOption struct {
	// 金币场
	GammerNum        *int32                          `protobuf:"varint,1,opt,name=gammerNum" json:"gammerNum,omitempty"`
	BoardsCout       *int32                          `protobuf:"varint,2,opt,name=boardsCout" json:"boardsCout,omitempty"`
	GanOpt           *ZhadanGanOpt                   `protobuf:"varint,3,opt,name=ganOpt,enum=ddproto.ZhadanGanOpt" json:"ganOpt,omitempty"`
	IsBaoZhuang      *bool                           `protobuf:"varint,4,opt,name=isBaoZhuang" json:"isBaoZhuang,omitempty"`
	HasAnimation     *bool                           `protobuf:"varint,8,opt,name=hasAnimation" json:"hasAnimation,omitempty"`
	BaseChip         *int32                          `protobuf:"varint,9,opt,name=baseChip" json:"baseChip,omitempty"`
	IsBuDaGan        *bool                           `protobuf:"varint,10,opt,name=isBuDaGan" json:"isBuDaGan,omitempty"`
	ChanelId         *int32                          `protobuf:"varint,11,opt,name=chanelId" json:"chanelId,omitempty"`
	RoomCardBillType *COMMON_ENUM_ROOMCARD_BILL_TYPE `protobuf:"varint,12,opt,name=roomCardBillType,enum=ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE" json:"roomCardBillType,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ZhadanDeskOption) Reset()                    { *m = ZhadanDeskOption{} }
func (m *ZhadanDeskOption) String() string            { return proto.CompactTextString(m) }
func (*ZhadanDeskOption) ProtoMessage()               {}
func (*ZhadanDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor57, []int{2} }

func (m *ZhadanDeskOption) GetGammerNum() int32 {
	if m != nil && m.GammerNum != nil {
		return *m.GammerNum
	}
	return 0
}

func (m *ZhadanDeskOption) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *ZhadanDeskOption) GetGanOpt() ZhadanGanOpt {
	if m != nil && m.GanOpt != nil {
		return *m.GanOpt
	}
	return ZhadanGanOpt_ZHADAN_GAN_OPT_BANGAN
}

func (m *ZhadanDeskOption) GetIsBaoZhuang() bool {
	if m != nil && m.IsBaoZhuang != nil {
		return *m.IsBaoZhuang
	}
	return false
}

func (m *ZhadanDeskOption) GetHasAnimation() bool {
	if m != nil && m.HasAnimation != nil {
		return *m.HasAnimation
	}
	return false
}

func (m *ZhadanDeskOption) GetBaseChip() int32 {
	if m != nil && m.BaseChip != nil {
		return *m.BaseChip
	}
	return 0
}

func (m *ZhadanDeskOption) GetIsBuDaGan() bool {
	if m != nil && m.IsBuDaGan != nil {
		return *m.IsBuDaGan
	}
	return false
}

func (m *ZhadanDeskOption) GetChanelId() int32 {
	if m != nil && m.ChanelId != nil {
		return *m.ChanelId
	}
	return 0
}

func (m *ZhadanDeskOption) GetRoomCardBillType() COMMON_ENUM_ROOMCARD_BILL_TYPE {
	if m != nil && m.RoomCardBillType != nil {
		return *m.RoomCardBillType
	}
	return COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY
}

// room 的信息
type ZhadanSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseChip         *int32  `protobuf:"varint,5,opt,name=baseChip" json:"baseChip,omitempty"`
	EnterCoin        *int32  `protobuf:"varint,6,opt,name=enterCoin" json:"enterCoin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZhadanSrvRoom) Reset()                    { *m = ZhadanSrvRoom{} }
func (m *ZhadanSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvRoom) ProtoMessage()               {}
func (*ZhadanSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor57, []int{3} }

func (m *ZhadanSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZhadanSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *ZhadanSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *ZhadanSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *ZhadanSrvRoom) GetBaseChip() int32 {
	if m != nil && m.BaseChip != nil {
		return *m.BaseChip
	}
	return 0
}

func (m *ZhadanSrvRoom) GetEnterCoin() int32 {
	if m != nil && m.EnterCoin != nil {
		return *m.EnterCoin
	}
	return 0
}

func init() {
	proto.RegisterType((*ZhadanClientPoker)(nil), "ddproto.zhadan_client_poker")
	proto.RegisterType((*ZhadanUserBill)(nil), "ddproto.zhadan_user_bill")
	proto.RegisterType((*ZhadanDeskOption)(nil), "ddproto.zhadan_desk_option")
	proto.RegisterType((*ZhadanSrvRoom)(nil), "ddproto.zhadan_srv_room")
	proto.RegisterEnum("ddproto.ZhadanEnumProtoid", ZhadanEnumProtoid_name, ZhadanEnumProtoid_value)
	proto.RegisterEnum("ddproto.ZhadanEnumPokerType", ZhadanEnumPokerType_name, ZhadanEnumPokerType_value)
	proto.RegisterEnum("ddproto.ZhadanEnumDeskStatus", ZhadanEnumDeskStatus_name, ZhadanEnumDeskStatus_value)
	proto.RegisterEnum("ddproto.ZhadanGanOpt", ZhadanGanOpt_name, ZhadanGanOpt_value)
}

func init() { proto.RegisterFile("zhadan_base.proto", fileDescriptor57) }

var fileDescriptor57 = []byte{
	// 966 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x95, 0xdd, 0x72, 0xe2, 0x36,
	0x14, 0x80, 0x63, 0x12, 0x92, 0xa0, 0x64, 0xb3, 0x5a, 0xe5, 0x07, 0x87, 0xfc, 0x2c, 0xcb, 0x4c,
	0xa7, 0x99, 0x5c, 0xa4, 0x33, 0xbd, 0xee, 0x8d, 0x6c, 0x14, 0x70, 0x63, 0x6c, 0xd6, 0x36, 0xbb,
	0x93, 0xdc, 0x68, 0x94, 0xe0, 0x01, 0xcf, 0x1a, 0x9b, 0xb1, 0xcd, 0x4e, 0xdb, 0x07, 0xd8, 0x8b,
	0xbe, 0x52, 0x9f, 0xa4, 0xef, 0xd1, 0x07, 0xe8, 0x48, 0x76, 0x40, 0x38, 0xc0, 0x15, 0xd2, 0xf9,
	0xce, 0xff, 0x39, 0xc2, 0xe0, 0xc3, 0x5f, 0x63, 0x36, 0x64, 0x11, 0x7d, 0x66, 0xa9, 0x7f, 0x37,
	0x4d, 0xe2, 0x2c, 0x46, 0x7b, 0xc3, 0xa1, 0x38, 0x34, 0x8e, 0x5f, 0xe2, 0xc9, 0x24, 0x8e, 0xe8,
	0x4b, 0x18, 0xf8, 0x51, 0x96, 0xd3, 0xd6, 0x08, 0x1c, 0x17, 0x26, 0xb9, 0x98, 0x4e, 0xe3, 0x6f,
	0x7e, 0x82, 0x10, 0xd8, 0x99, 0xb2, 0x20, 0x55, 0x2b, 0xcd, 0xed, 0x9b, 0xaa, 0x23, 0xce, 0xe8,
	0x37, 0x50, 0x13, 0xd0, 0xfb, 0x73, 0xea, 0xab, 0xdb, 0x4d, 0xe5, 0xe6, 0xe8, 0xd7, 0xeb, 0xbb,
	0xc2, 0xf9, 0x5d, 0xe1, 0xc4, 0x8f, 0x66, 0x13, 0x3a, 0xd7, 0x72, 0x16, 0x06, 0xad, 0x1f, 0x0a,
	0x80, 0x85, 0xd2, 0x2c, 0xf5, 0x13, 0xfa, 0x1c, 0x84, 0x21, 0x3a, 0x03, 0xbb, 0xfc, 0x62, 0x0c,
	0x55, 0xa5, 0xa9, 0xdc, 0xbc, 0x73, 0x8a, 0x1b, 0x3a, 0x01, 0xd5, 0xf4, 0x25, 0x4e, 0x7c, 0xb5,
	0xd2, 0x54, 0x6e, 0xaa, 0x4e, 0x7e, 0x41, 0x2d, 0x70, 0xc8, 0xc2, 0x50, 0x63, 0xa9, 0xef, 0x0a,
	0xb8, 0x2d, 0xe0, 0x92, 0x0c, 0x35, 0xc1, 0x01, 0x0b, 0x43, 0xf2, 0x47, 0x96, 0xab, 0xec, 0x08,
	0x15, 0x59, 0xd4, 0xfa, 0xaf, 0x02, 0x50, 0x91, 0xc8, 0xd0, 0x4f, 0xbf, 0xd1, 0x78, 0x9a, 0x05,
	0x71, 0x84, 0x2e, 0x41, 0x6d, 0xc4, 0x26, 0x13, 0x3f, 0xb1, 0x66, 0x13, 0x91, 0x4d, 0xd5, 0x59,
	0x08, 0xd0, 0x35, 0x00, 0xcf, 0x31, 0x4b, 0x86, 0xa9, 0x1e, 0xcf, 0xb2, 0x22, 0x2b, 0x49, 0x82,
	0x7e, 0x01, 0xbb, 0x23, 0x16, 0xd9, 0xd3, 0xac, 0x68, 0x4c, 0xbd, 0xdc, 0x98, 0x11, 0x8b, 0x78,
	0x24, 0xa7, 0x50, 0xe3, 0x79, 0x06, 0xa9, 0xc6, 0xe2, 0xa7, 0xf1, 0x8c, 0x45, 0x23, 0x91, 0xe7,
	0xbe, 0x23, 0x8b, 0x78, 0xb5, 0x63, 0x96, 0xe2, 0x28, 0x98, 0x30, 0x9e, 0xa0, 0xba, 0x2f, 0x54,
	0x96, 0x64, 0xa8, 0x01, 0xf6, 0xf9, 0xa4, 0xf5, 0x71, 0x30, 0x55, 0x6b, 0x22, 0xa9, 0xf9, 0x9d,
	0x17, 0x14, 0xa4, 0xda, 0xac, 0xcd, 0x3a, 0x2c, 0x52, 0x81, 0x30, 0x5e, 0x08, 0xb8, 0xe5, 0xcb,
	0x98, 0x45, 0x7e, 0x68, 0x0c, 0xd5, 0x83, 0xdc, 0xf2, 0xf5, 0x8e, 0x5c, 0x00, 0x93, 0x38, 0x9e,
	0xe8, 0x2c, 0x19, 0x6a, 0x41, 0x18, 0x8a, 0x79, 0x1f, 0x8a, 0xb2, 0x7e, 0x9e, 0x97, 0xa5, 0xdb,
	0xbd, 0x9e, 0x6d, 0x51, 0x62, 0x0d, 0x7a, 0xd4, 0xb1, 0xed, 0x9e, 0x8e, 0x9d, 0x36, 0xd5, 0x0c,
	0xd3, 0xa4, 0xde, 0x63, 0x9f, 0x38, 0x6f, 0x1c, 0xb4, 0xfe, 0x51, 0xc0, 0xfb, 0xa2, 0x17, 0x69,
	0xf2, 0x9d, 0x72, 0xce, 0xc7, 0xcf, 0x7f, 0x8b, 0xf1, 0x57, 0x9d, 0xe2, 0xc6, 0x93, 0xe3, 0x27,
	0x11, 0x38, 0xef, 0xf5, 0xfc, 0xce, 0xcb, 0xe2, 0x67, 0xd3, 0xff, 0xee, 0x87, 0xc5, 0x06, 0x2c,
	0x04, 0xaf, 0xd4, 0x0b, 0xb2, 0x30, 0x1f, 0x7e, 0xcd, 0x59, 0x08, 0x96, 0xda, 0x55, 0x7d, 0xdb,
	0x2e, 0x3f, 0xca, 0xfc, 0x44, 0x8f, 0x83, 0x48, 0xdd, 0xcd, 0xfd, 0xce, 0x05, 0xb7, 0xff, 0xee,
	0xcd, 0xdf, 0x49, 0xbe, 0xe2, 0xbc, 0x09, 0xc1, 0x10, 0xa9, 0xe0, 0xe4, 0xa9, 0x8b, 0xdb, 0xd8,
	0xa2, 0x7d, 0xa3, 0x4d, 0xbb, 0x04, 0x3b, 0x9e, 0x46, 0xb0, 0x07, 0xb7, 0xd0, 0x35, 0x68, 0x48,
	0x44, 0x77, 0x08, 0xf6, 0x08, 0x6d, 0x13, 0xf7, 0x81, 0x3a, 0xe4, 0x33, 0x54, 0xd0, 0x15, 0x38,
	0x97, 0x38, 0xb1, 0x3c, 0xe2, 0x2c, 0x70, 0x65, 0x3d, 0xc6, 0xfa, 0x03, 0xdc, 0x46, 0x97, 0x40,
	0x5d, 0x8d, 0x35, 0x1d, 0xee, 0x94, 0xb2, 0x72, 0x08, 0x6e, 0x3f, 0x0a, 0xb7, 0x55, 0x54, 0x07,
	0xc7, 0x6f, 0x88, 0xa6, 0xc3, 0xdd, 0x12, 0xb8, 0xc7, 0x7d, 0x6c, 0x50, 0xdb, 0x83, 0x7b, 0xe8,
	0x02, 0xd4, 0x25, 0xa0, 0x61, 0xfb, 0xa9, 0x3b, 0xc0, 0x56, 0x87, 0xc3, 0xfd, 0x52, 0x1a, 0x0b,
	0xc8, 0x83, 0xd5, 0xd6, 0x52, 0x5e, 0x02, 0xd8, 0x60, 0xeb, 0xc2, 0x83, 0x52, 0x09, 0x7a, 0x77,
	0x50, 0x24, 0x74, 0x88, 0xce, 0xc1, 0xe9, 0x5b, 0xc2, 0x03, 0xbe, 0x5b, 0x6d, 0xa4, 0xe9, 0xf0,
	0xa8, 0x54, 0x5e, 0x1f, 0xbb, 0xae, 0x30, 0x79, 0x8f, 0xce, 0x00, 0x2a, 0x03, 0x4d, 0x87, 0xb0,
	0x34, 0xbe, 0x0e, 0xee, 0x11, 0x4a, 0xac, 0x36, 0xb5, 0x2d, 0xc2, 0xf9, 0x87, 0x75, 0x1c, 0x9b,
	0x26, 0xe7, 0x08, 0x7d, 0x02, 0x57, 0x12, 0xc7, 0xfd, 0xbe, 0xf9, 0x48, 0xdb, 0x86, 0xeb, 0xda,
	0xe6, 0x17, 0x22, 0x42, 0x1f, 0x6f, 0x56, 0xe1, 0x3d, 0x3a, 0x41, 0x4d, 0x70, 0x29, 0xa9, 0xcc,
	0xa1, 0x86, 0xf5, 0x7c, 0x4f, 0x4e, 0x37, 0x6a, 0x70, 0x1f, 0x67, 0xa5, 0x30, 0xf6, 0x57, 0x8b,
	0xaf, 0x8a, 0x9c, 0x49, 0x7d, 0xb3, 0x0a, 0xf7, 0xa2, 0xa2, 0x8f, 0xe0, 0x42, 0x52, 0x71, 0x79,
	0xad, 0x3d, 0xe2, 0xba, 0xb8, 0x93, 0xfb, 0x38, 0x2f, 0x8d, 0xe5, 0x95, 0x69, 0x3a, 0x6c, 0x94,
	0x76, 0xd9, 0x24, 0xf8, 0x8b, 0xf4, 0x12, 0x2e, 0xd6, 0x63, 0x1e, 0xf9, 0xb2, 0xe4, 0xd8, 0xbe,
	0xbf, 0x37, 0x8d, 0x7c, 0x08, 0x57, 0xa5, 0xbc, 0x75, 0xdb, 0xb0, 0xc4, 0x7f, 0x11, 0x35, 0x0d,
	0xd7, 0x13, 0xce, 0xaf, 0x37, 0xab, 0xf0, 0x00, 0x1f, 0x6f, 0xff, 0xae, 0x80, 0xd3, 0x95, 0x9f,
	0x2f, 0x79, 0xf7, 0xed, 0x07, 0xe2, 0x88, 0xbf, 0x36, 0x6a, 0x7b, 0x5d, 0xe2, 0xc0, 0x2d, 0x39,
	0xed, 0x05, 0x14, 0x57, 0x6c, 0x2c, 0xbf, 0x6f, 0x09, 0x0f, 0x0c, 0x81, 0x2b, 0xf2, 0xfe, 0x2c,
	0xb0, 0x8b, 0x2d, 0xea, 0x19, 0xd8, 0x5e, 0x7e, 0xe0, 0x12, 0xef, 0x0e, 0xac, 0x27, 0x03, 0xee,
	0xac, 0x76, 0x6e, 0x1a, 0xd8, 0x6a, 0x0f, 0x0c, 0x58, 0x5d, 0x9d, 0xf7, 0x3d, 0x31, 0x7e, 0x37,
	0xe0, 0xee, 0x6a, 0xcf, 0xb9, 0x04, 0xee, 0xdd, 0xfe, 0x50, 0x40, 0x5d, 0x6e, 0x86, 0xf8, 0x44,
	0xa6, 0x19, 0xcb, 0x66, 0xa9, 0xd4, 0x4b, 0x31, 0x1e, 0xd7, 0xc3, 0xde, 0xc0, 0xa5, 0x5f, 0xb1,
	0xe1, 0xe5, 0xff, 0x24, 0x50, 0x41, 0x3f, 0x81, 0x4f, 0xeb, 0x54, 0xe6, 0x2f, 0x1c, 0x56, 0xd6,
	0x78, 0xb2, 0x2d, 0xfe, 0x90, 0x7a, 0x86, 0xd5, 0x81, 0xdb, 0xb7, 0x1d, 0x70, 0xb4, 0xfc, 0xe9,
	0x94, 0x16, 0xa1, 0x83, 0x2d, 0x6a, 0xf7, 0xb9, 0x4b, 0xab, 0x83, 0x2d, 0xa8, 0xa0, 0x06, 0x38,
	0x2b, 0xa1, 0xcf, 0x83, 0x9c, 0x55, 0xfa, 0x5b, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x57, 0xf3,
	0xf6, 0x7d, 0x15, 0x09, 0x00, 0x00,
}