// Code generated by protoc-gen-go.
// source: sdy_play.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of cm_offline from common_client.proto

// Ignoring public import of cm_hearbeat from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of sdy_base_userPaiIds from sdy_base.proto

// Ignoring public import of sdy_base_roomTypeInfo from sdy_base.proto

// Ignoring public import of sdy_base_timerInfo from sdy_base.proto

// Ignoring public import of sdy_base_playerInfo from sdy_base.proto

// Ignoring public import of sdy_base_commonRateInfo from sdy_base.proto

// Ignoring public import of sdy_base_deskInfo from sdy_base.proto

// Ignoring public import of sdy_enum_protoId from sdy_base.proto

// Ignoring public import of sdy_enum_errorCode from sdy_base.proto

// Ignoring public import of sdy_enum_actType from sdy_base.proto

// Ignoring public import of sdy_enum_deskStatus from sdy_base.proto

// Ignoring public import of sdy_enum_userStatus from sdy_base.proto

// Ignoring public import of sdy_enum_Option from sdy_base.proto

// Ignoring public import of sdy_enum_roomLevel from sdy_base.proto

// Ignoring public import of sdy_enum_flowers from sdy_base.proto

type SdyEnumJScore int32

const (
	SdyEnumJScore_SDY_SIXTY       SdyEnumJScore = 1
	SdyEnumJScore_SDY_SIXTYFIVE   SdyEnumJScore = 2
	SdyEnumJScore_SDY_SEVENTY     SdyEnumJScore = 3
	SdyEnumJScore_SDY_SEVENTYFIVE SdyEnumJScore = 4
	SdyEnumJScore_SDY_NONE        SdyEnumJScore = 5
)

var SdyEnumJScore_name = map[int32]string{
	1: "SDY_SIXTY",
	2: "SDY_SIXTYFIVE",
	3: "SDY_SEVENTY",
	4: "SDY_SEVENTYFIVE",
	5: "SDY_NONE",
}
var SdyEnumJScore_value = map[string]int32{
	"SDY_SIXTY":       1,
	"SDY_SIXTYFIVE":   2,
	"SDY_SEVENTY":     3,
	"SDY_SEVENTYFIVE": 4,
	"SDY_NONE":        5,
}

func (x SdyEnumJScore) Enum() *SdyEnumJScore {
	p := new(SdyEnumJScore)
	*p = x
	return p
}
func (x SdyEnumJScore) String() string {
	return proto.EnumName(SdyEnumJScore_name, int32(x))
}
func (x *SdyEnumJScore) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SdyEnumJScore_value, data, "SdyEnumJScore")
	if err != nil {
		return err
	}
	*x = SdyEnumJScore(value)
	return nil
}
func (SdyEnumJScore) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

// 开局（接收服务端消息）
type SdyBcOpening struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcOpening) Reset()                    { *m = SdyBcOpening{} }
func (m *SdyBcOpening) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOpening) ProtoMessage()               {}
func (*SdyBcOpening) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SdyBcOpening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOpening) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 发牌
type SdyDealCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []int32      `protobuf:"varint,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyDealCards) Reset()                    { *m = SdyDealCards{} }
func (m *SdyDealCards) String() string            { return proto.CompactTextString(m) }
func (*SdyDealCards) ProtoMessage()               {}
func (*SdyDealCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *SdyDealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyDealCards) GetPlayerPokers() []int32 {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

type SdyBcPlayerPokers struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PlayerPokers     []*SdyBaseUserPaiIds `protobuf:"bytes,2,rep,name=playerPokers" json:"playerPokers,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcPlayerPokers) Reset()                    { *m = SdyBcPlayerPokers{} }
func (m *SdyBcPlayerPokers) String() string            { return proto.CompactTextString(m) }
func (*SdyBcPlayerPokers) ProtoMessage()               {}
func (*SdyBcPlayerPokers) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *SdyBcPlayerPokers) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcPlayerPokers) GetPlayerPokers() []*SdyBaseUserPaiIds {
	if m != nil {
		return m.PlayerPokers
	}
	return nil
}

// 广播叫分的消息，包含当前谁叫分，下一个叫分的玩家
type SdyBcJiaoFen struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrentUserId    *uint32         `protobuf:"varint,2,opt,name=currentUserId" json:"currentUserId,omitempty"`
	NextUserId       *uint32         `protobuf:"varint,3,opt,name=nextUserId" json:"nextUserId,omitempty"`
	CanJScore        []SdyEnumJScore `protobuf:"varint,4,rep,name=canJScore,enum=yjprotogo.SdyEnumJScore" json:"canJScore,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SdyBcJiaoFen) Reset()                    { *m = SdyBcJiaoFen{} }
func (m *SdyBcJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyBcJiaoFen) ProtoMessage()               {}
func (*SdyBcJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SdyBcJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcJiaoFen) GetCurrentUserId() uint32 {
	if m != nil && m.CurrentUserId != nil {
		return *m.CurrentUserId
	}
	return 0
}

func (m *SdyBcJiaoFen) GetNextUserId() uint32 {
	if m != nil && m.NextUserId != nil {
		return *m.NextUserId
	}
	return 0
}

func (m *SdyBcJiaoFen) GetCanJScore() []SdyEnumJScore {
	if m != nil {
		return m.CanJScore
	}
	return nil
}

// 叫分,玩家选择分数后由客户端请求叫分的消息
type SdyReqJiaoFen struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Score            *SdyEnumJScore `protobuf:"varint,3,opt,name=score,enum=yjprotogo.SdyEnumJScore" json:"score,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyReqJiaoFen) Reset()                    { *m = SdyReqJiaoFen{} }
func (m *SdyReqJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyReqJiaoFen) ProtoMessage()               {}
func (*SdyReqJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *SdyReqJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqJiaoFen) GetScore() SdyEnumJScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJScore_SDY_SIXTY
}

// 叫分回复，失败时回复
type SdyAckJiaoFen struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckJiaoFen) Reset()                    { *m = SdyAckJiaoFen{} }
func (m *SdyAckJiaoFen) String() string            { return proto.CompactTextString(m) }
func (*SdyAckJiaoFen) ProtoMessage()               {}
func (*SdyAckJiaoFen) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *SdyAckJiaoFen) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckJiaoFen) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 叫分成功后广播叫分的结果
type SdyBcJiaoFenResult struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Score            *SdyEnumJScore `protobuf:"varint,3,opt,name=score,enum=yjprotogo.SdyEnumJScore" json:"score,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyBcJiaoFenResult) Reset()                    { *m = SdyBcJiaoFenResult{} }
func (m *SdyBcJiaoFenResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcJiaoFenResult) ProtoMessage()               {}
func (*SdyBcJiaoFenResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *SdyBcJiaoFenResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcJiaoFenResult) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcJiaoFenResult) GetScore() SdyEnumJScore {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return SdyEnumJScore_SDY_SIXTY
}

// 换底结束，开始游戏 (广播)
type SdyBcStartPlay struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Banker           *uint32      `protobuf:"varint,2,opt,name=banker" json:"banker,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcStartPlay) Reset()                    { *m = SdyBcStartPlay{} }
func (m *SdyBcStartPlay) String() string            { return proto.CompactTextString(m) }
func (*SdyBcStartPlay) ProtoMessage()               {}
func (*SdyBcStartPlay) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *SdyBcStartPlay) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcStartPlay) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

// 出牌
type SdyReqOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	OutCards         *int32       `protobuf:"varint,2,opt,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqOutCards) Reset()                    { *m = SdyReqOutCards{} }
func (m *SdyReqOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyReqOutCards) ProtoMessage()               {}
func (*SdyReqOutCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *SdyReqOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqOutCards) GetOutCards() int32 {
	if m != nil && m.OutCards != nil {
		return *m.OutCards
	}
	return 0
}

type SdyAckOutCards struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RemainPokers     *int32       `protobuf:"varint,3,opt,name=remainPokers" json:"remainPokers,omitempty"`
	OutCards         *int32       `protobuf:"varint,4,opt,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckOutCards) Reset()                    { *m = SdyAckOutCards{} }
func (m *SdyAckOutCards) String() string            { return proto.CompactTextString(m) }
func (*SdyAckOutCards) ProtoMessage()               {}
func (*SdyAckOutCards) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *SdyAckOutCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckOutCards) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckOutCards) GetRemainPokers() int32 {
	if m != nil && m.RemainPokers != nil {
		return *m.RemainPokers
	}
	return 0
}

func (m *SdyAckOutCards) GetOutCards() int32 {
	if m != nil && m.OutCards != nil {
		return *m.OutCards
	}
	return 0
}

// 轮到谁操作
type SdyBcOverTurn struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcOverTurn) Reset()                    { *m = SdyBcOverTurn{} }
func (m *SdyBcOverTurn) String() string            { return proto.CompactTextString(m) }
func (*SdyBcOverTurn) ProtoMessage()               {}
func (*SdyBcOverTurn) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *SdyBcOverTurn) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcOverTurn) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 广播上一个玩家出了什么牌的消息
type SdyBcWhatPai struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PaiId            *int32       `protobuf:"varint,2,opt,name=paiId" json:"paiId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcWhatPai) Reset()                    { *m = SdyBcWhatPai{} }
func (m *SdyBcWhatPai) String() string            { return proto.CompactTextString(m) }
func (*SdyBcWhatPai) ProtoMessage()               {}
func (*SdyBcWhatPai) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

func (m *SdyBcWhatPai) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcWhatPai) GetPaiId() int32 {
	if m != nil && m.PaiId != nil {
		return *m.PaiId
	}
	return 0
}

// 一圈打完后发送是否闲家得分的广播，以及得的多少分
type SdyBcScorePai struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ScorePai         []int32      `protobuf:"varint,2,rep,name=scorePai" json:"scorePai,omitempty"`
	IsPoPai          *bool        `protobuf:"varint,3,opt,name=isPoPai" json:"isPoPai,omitempty"`
	IsShangChe       *bool        `protobuf:"varint,4,opt,name=isShangChe" json:"isShangChe,omitempty"`
	IsKouDi          *bool        `protobuf:"varint,5,opt,name=isKouDi" json:"isKouDi,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcScorePai) Reset()                    { *m = SdyBcScorePai{} }
func (m *SdyBcScorePai) String() string            { return proto.CompactTextString(m) }
func (*SdyBcScorePai) ProtoMessage()               {}
func (*SdyBcScorePai) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *SdyBcScorePai) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcScorePai) GetScorePai() []int32 {
	if m != nil {
		return m.ScorePai
	}
	return nil
}

func (m *SdyBcScorePai) GetIsPoPai() bool {
	if m != nil && m.IsPoPai != nil {
		return *m.IsPoPai
	}
	return false
}

func (m *SdyBcScorePai) GetIsShangChe() bool {
	if m != nil && m.IsShangChe != nil {
		return *m.IsShangChe
	}
	return false
}

func (m *SdyBcScorePai) GetIsKouDi() bool {
	if m != nil && m.IsKouDi != nil {
		return *m.IsKouDi
	}
	return false
}

// 游戏信息(广播)
type SdyBcGameInfo struct {
	Header           *ProtoHeader         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	SenderUserId     *uint32              `protobuf:"varint,2,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool                `protobuf:"varint,3,opt,name=isReconnect" json:"isReconnect,omitempty"`
	PlayerInfo       []*SdyBasePlayerInfo `protobuf:"bytes,4,rep,name=playerInfo" json:"playerInfo,omitempty"`
	DeskInfo         *SdyBaseDeskInfo     `protobuf:"bytes,5,opt,name=deskInfo" json:"deskInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SdyBcGameInfo) Reset()                    { *m = SdyBcGameInfo{} }
func (m *SdyBcGameInfo) String() string            { return proto.CompactTextString(m) }
func (*SdyBcGameInfo) ProtoMessage()               {}
func (*SdyBcGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{13} }

func (m *SdyBcGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *SdyBcGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func (m *SdyBcGameInfo) GetPlayerInfo() []*SdyBasePlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *SdyBcGameInfo) GetDeskInfo() *SdyBaseDeskInfo {
	if m != nil {
		return m.DeskInfo
	}
	return nil
}

// 定主， 广播消息，通知其他玩家谁在定主，包含定主的人
type SdyBcDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	FootPokers       []int32      `protobuf:"varint,3,rep,name=footPokers" json:"footPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcDingZhu) Reset()                    { *m = SdyBcDingZhu{} }
func (m *SdyBcDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyBcDingZhu) ProtoMessage()               {}
func (*SdyBcDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{14} }

func (m *SdyBcDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyBcDingZhu) GetFootPokers() []int32 {
	if m != nil {
		return m.FootPokers
	}
	return nil
}

// 选好花色后请求定主
type SdyReqDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Flower           *int32       `protobuf:"varint,3,opt,name=flower" json:"flower,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqDingZhu) Reset()                    { *m = SdyReqDingZhu{} }
func (m *SdyReqDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyReqDingZhu) ProtoMessage()               {}
func (*SdyReqDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{15} }

func (m *SdyReqDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqDingZhu) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

type SdyBcDingZhuResult struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Flower           *int32       `protobuf:"varint,2,opt,name=flower" json:"flower,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcDingZhuResult) Reset()                    { *m = SdyBcDingZhuResult{} }
func (m *SdyBcDingZhuResult) String() string            { return proto.CompactTextString(m) }
func (*SdyBcDingZhuResult) ProtoMessage()               {}
func (*SdyBcDingZhuResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{16} }

func (m *SdyBcDingZhuResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcDingZhuResult) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

// 定主的回复
type SdyAckDingZhu struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckDingZhu) Reset()                    { *m = SdyAckDingZhu{} }
func (m *SdyAckDingZhu) String() string            { return proto.CompactTextString(m) }
func (*SdyAckDingZhu) ProtoMessage()               {}
func (*SdyAckDingZhu) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{17} }

func (m *SdyAckDingZhu) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckDingZhu) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 广播换底，定主完毕后广播换底，
type SdyBcHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyBcHuanDi) Reset()                    { *m = SdyBcHuanDi{} }
func (m *SdyBcHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyBcHuanDi) ProtoMessage()               {}
func (*SdyBcHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{18} }

func (m *SdyBcHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyBcHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type SdyReqHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	OutCards         []int32      `protobuf:"varint,3,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqHuanDi) Reset()                    { *m = SdyReqHuanDi{} }
func (m *SdyReqHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyReqHuanDi) ProtoMessage()               {}
func (*SdyReqHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{19} }

func (m *SdyReqHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyReqHuanDi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

type SdyAckHuanDi struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsHuanDiOver     *bool        `protobuf:"varint,3,opt,name=isHuanDiOver" json:"isHuanDiOver,omitempty"`
	OutCards         []int32      `protobuf:"varint,4,rep,name=outCards" json:"outCards,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyAckHuanDi) Reset()                    { *m = SdyAckHuanDi{} }
func (m *SdyAckHuanDi) String() string            { return proto.CompactTextString(m) }
func (*SdyAckHuanDi) ProtoMessage()               {}
func (*SdyAckHuanDi) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{20} }

func (m *SdyAckHuanDi) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckHuanDi) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *SdyAckHuanDi) GetIsHuanDiOver() bool {
	if m != nil && m.IsHuanDiOver != nil {
		return *m.IsHuanDiOver
	}
	return false
}

func (m *SdyAckHuanDi) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

// 三打一换桌message
type SdyReqChangeDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsChangeDesk     *bool        `protobuf:"varint,2,opt,name=isChangeDesk" json:"isChangeDesk,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SdyReqChangeDesk) Reset()                    { *m = SdyReqChangeDesk{} }
func (m *SdyReqChangeDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyReqChangeDesk) ProtoMessage()               {}
func (*SdyReqChangeDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{21} }

func (m *SdyReqChangeDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyReqChangeDesk) GetIsChangeDesk() bool {
	if m != nil && m.IsChangeDesk != nil {
		return *m.IsChangeDesk
	}
	return false
}

type SdyAckChangeDesk struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	IsChangeOk       *bool          `protobuf:"varint,2,opt,name=isChangeOk" json:"isChangeOk,omitempty"`
	GameInfo         *SdyBcGameInfo `protobuf:"bytes,3,opt,name=gameInfo" json:"gameInfo,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SdyAckChangeDesk) Reset()                    { *m = SdyAckChangeDesk{} }
func (m *SdyAckChangeDesk) String() string            { return proto.CompactTextString(m) }
func (*SdyAckChangeDesk) ProtoMessage()               {}
func (*SdyAckChangeDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{22} }

func (m *SdyAckChangeDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SdyAckChangeDesk) GetIsChangeOk() bool {
	if m != nil && m.IsChangeOk != nil {
		return *m.IsChangeOk
	}
	return false
}

func (m *SdyAckChangeDesk) GetGameInfo() *SdyBcGameInfo {
	if m != nil {
		return m.GameInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SdyBcOpening)(nil), "yjprotogo.sdy_bc_opening")
	proto.RegisterType((*SdyDealCards)(nil), "yjprotogo.sdy_dealCards")
	proto.RegisterType((*SdyBcPlayerPokers)(nil), "yjprotogo.sdy_bc_playerPokers")
	proto.RegisterType((*SdyBcJiaoFen)(nil), "yjprotogo.sdy_bc_jiaoFen")
	proto.RegisterType((*SdyReqJiaoFen)(nil), "yjprotogo.sdy_req_jiaoFen")
	proto.RegisterType((*SdyAckJiaoFen)(nil), "yjprotogo.sdy_ack_jiaoFen")
	proto.RegisterType((*SdyBcJiaoFenResult)(nil), "yjprotogo.sdy_bc_jiaoFenResult")
	proto.RegisterType((*SdyBcStartPlay)(nil), "yjprotogo.sdy_bc_startPlay")
	proto.RegisterType((*SdyReqOutCards)(nil), "yjprotogo.sdy_req_outCards")
	proto.RegisterType((*SdyAckOutCards)(nil), "yjprotogo.sdy_ack_outCards")
	proto.RegisterType((*SdyBcOverTurn)(nil), "yjprotogo.sdy_bc_overTurn")
	proto.RegisterType((*SdyBcWhatPai)(nil), "yjprotogo.sdy_bc_whatPai")
	proto.RegisterType((*SdyBcScorePai)(nil), "yjprotogo.sdy_bc_scorePai")
	proto.RegisterType((*SdyBcGameInfo)(nil), "yjprotogo.sdy_bc_gameInfo")
	proto.RegisterType((*SdyBcDingZhu)(nil), "yjprotogo.sdy_bc_dingZhu")
	proto.RegisterType((*SdyReqDingZhu)(nil), "yjprotogo.sdy_req_dingZhu")
	proto.RegisterType((*SdyBcDingZhuResult)(nil), "yjprotogo.sdy_bc_dingZhuResult")
	proto.RegisterType((*SdyAckDingZhu)(nil), "yjprotogo.sdy_ack_dingZhu")
	proto.RegisterType((*SdyBcHuanDi)(nil), "yjprotogo.sdy_bc_huanDi")
	proto.RegisterType((*SdyReqHuanDi)(nil), "yjprotogo.sdy_req_huanDi")
	proto.RegisterType((*SdyAckHuanDi)(nil), "yjprotogo.sdy_ack_huanDi")
	proto.RegisterType((*SdyReqChangeDesk)(nil), "yjprotogo.sdy_req_changeDesk")
	proto.RegisterType((*SdyAckChangeDesk)(nil), "yjprotogo.sdy_ack_changeDesk")
	proto.RegisterEnum("yjprotogo.SdyEnumJScore", SdyEnumJScore_name, SdyEnumJScore_value)
}

var fileDescriptor11 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x41, 0x6f, 0xe2, 0x46,
	0x14, 0xae, 0x21, 0xa4, 0xce, 0x23, 0x6c, 0xe8, 0x64, 0x15, 0x21, 0x54, 0x45, 0x68, 0xd4, 0x03,
	0xea, 0x01, 0x55, 0x39, 0x54, 0x7b, 0xea, 0x61, 0x03, 0xab, 0xd0, 0x4a, 0x89, 0x65, 0xd2, 0x74,
	0xd9, 0xc3, 0xa2, 0xc1, 0x9e, 0x80, 0x03, 0xcc, 0xb0, 0x33, 0xf6, 0xee, 0x72, 0xec, 0xb1, 0xea,
	0xa5, 0xc7, 0xaa, 0xea, 0x9f, 0xe8, 0x3f, 0xe8, 0x4f, 0xab, 0x66, 0xc6, 0x36, 0x36, 0x4d, 0x2a,
	0xe1, 0x5a, 0xda, 0x53, 0xf2, 0xbe, 0x79, 0xef, 0x7b, 0xdf, 0xbc, 0x99, 0xf9, 0x30, 0x3c, 0x93,
	0xfe, 0x66, 0xb2, 0x5e, 0x92, 0x4d, 0x6f, 0x2d, 0x78, 0xc8, 0xd1, 0xd1, 0xe6, 0x41, 0xff, 0x33,
	0xe3, 0xed, 0x53, 0x8f, 0xaf, 0x56, 0x9c, 0x4d, 0xbc, 0x65, 0x40, 0x59, 0x68, 0xd6, 0xdb, 0x3a,
	0x7f, 0x4a, 0x24, 0x35, 0x31, 0x7e, 0x6d, 0x18, 0xa6, 0xde, 0x84, 0xaf, 0x29, 0x0b, 0xd8, 0x0c,
	0xf5, 0xe0, 0x70, 0x4e, 0x89, 0x4f, 0x45, 0xcb, 0xea, 0x58, 0xdd, 0xfa, 0xc5, 0x59, 0x2f, 0xa5,
	0xec, 0x39, 0xea, 0xef, 0x95, 0x5e, 0x75, 0xe3, 0x2c, 0x74, 0x06, 0x87, 0x91, 0xa4, 0x62, 0xe8,
	0xb7, 0x2a, 0x1d, 0xab, 0xdb, 0x70, 0xe3, 0x08, 0x7b, 0xd0, 0x50, 0xcc, 0x3e, 0x25, 0xcb, 0x4b,
	0x22, 0x7c, 0xb9, 0x37, 0x31, 0x86, 0x63, 0xb5, 0x31, 0x2a, 0x1c, 0xbe, 0xa0, 0x42, 0xb6, 0x2a,
	0x9d, 0x6a, 0xb7, 0xe6, 0xe6, 0x30, 0xfc, 0x8b, 0x05, 0xa7, 0xb1, 0xfe, 0x2c, 0xbe, 0x77, 0xaf,
	0x97, 0x8f, 0xf4, 0xaa, 0x5f, 0x9c, 0x67, 0xaa, 0x92, 0xb9, 0x4d, 0xd4, 0xf6, 0x1c, 0x12, 0x0c,
	0x7d, 0xb9, 0xa3, 0xe5, 0x6f, 0x2b, 0x9d, 0xe5, 0x43, 0x40, 0xf8, 0x2b, 0xca, 0xf6, 0x96, 0xf1,
	0x15, 0x34, 0xbc, 0x48, 0x08, 0xca, 0xc2, 0x1f, 0xb3, 0x23, 0xcd, 0x83, 0xe8, 0x1c, 0x80, 0xd1,
	0x8f, 0x49, 0x4a, 0x55, 0xa7, 0x64, 0x10, 0xf4, 0x02, 0x8e, 0x3c, 0xc2, 0xbe, 0x1f, 0x79, 0x5c,
	0xd0, 0xd6, 0x41, 0xa7, 0xda, 0x7d, 0x76, 0xd1, 0xde, 0xd9, 0x09, 0x65, 0xd1, 0x6a, 0xf2, 0xa0,
	0x33, 0xdc, 0x6d, 0x32, 0xfe, 0xd5, 0x82, 0x13, 0xb5, 0x2c, 0xe8, 0xbb, 0xc2, 0x7b, 0x78, 0xe2,
	0x3e, 0xa0, 0x6f, 0xa0, 0x26, 0xb5, 0x22, 0x25, 0xf8, 0xbf, 0x15, 0x99, 0x44, 0x3c, 0x36, 0x62,
	0x88, 0xb7, 0x28, 0x5b, 0x0c, 0xfe, 0xcd, 0x82, 0xe7, 0xf9, 0xb3, 0x72, 0xa9, 0x8c, 0x96, 0xe1,
	0x27, 0xdc, 0xed, 0x1b, 0x68, 0xc6, 0x8a, 0x64, 0x48, 0x44, 0xe8, 0x2c, 0xc9, 0xa6, 0x88, 0x9a,
	0x29, 0x61, 0x0b, 0x2a, 0x12, 0x35, 0x26, 0xc2, 0x6f, 0x0d, 0xb7, 0x3a, 0x56, 0x1e, 0x85, 0xc5,
	0x9e, 0x63, 0x1b, 0xec, 0xa4, 0x56, 0xb3, 0xd7, 0xdc, 0x34, 0xc6, 0x7f, 0x58, 0xa6, 0x81, 0x3a,
	0xaa, 0xc2, 0x0d, 0x9e, 0x1a, 0x25, 0x86, 0x63, 0x41, 0x57, 0x24, 0x60, 0xf1, 0xdb, 0xac, 0xea,
	0xe6, 0x39, 0x2c, 0x27, 0xee, 0x60, 0x47, 0x5c, 0x7c, 0x8d, 0x94, 0xc5, 0xbd, 0xa7, 0xe2, 0x36,
	0x12, 0xe5, 0x5d, 0xa3, 0xbb, 0xf4, 0xc5, 0x7f, 0x98, 0x93, 0xd0, 0x21, 0xc1, 0xde, 0xcc, 0xcf,
	0xa1, 0xb6, 0x56, 0x66, 0x12, 0x8f, 0xd4, 0x04, 0xf8, 0x2f, 0x2b, 0xd5, 0xac, 0x2f, 0x47, 0x11,
	0xe6, 0x36, 0xd8, 0x49, 0x6d, 0x6c, 0x9d, 0x69, 0x8c, 0x5a, 0xf0, 0x79, 0x20, 0x1d, 0xae, 0x96,
	0xd4, 0x34, 0x6d, 0x37, 0x09, 0x95, 0xb7, 0x04, 0x72, 0x34, 0x27, 0x6c, 0x76, 0x39, 0xa7, 0x7a,
	0x94, 0xb6, 0x9b, 0x41, 0x4c, 0xe5, 0x0f, 0x3c, 0xea, 0x07, 0xad, 0x5a, 0x52, 0xa9, 0x43, 0xfc,
	0x73, 0x25, 0xd5, 0x3c, 0x23, 0x2b, 0x3a, 0x64, 0xf7, 0xbc, 0x88, 0xe5, 0x4b, 0xca, 0x7c, 0x2a,
	0x72, 0xf6, 0x97, 0xc3, 0x50, 0x07, 0xea, 0x81, 0x74, 0xa9, 0xc7, 0x19, 0xa3, 0x5e, 0x18, 0xeb,
	0xcf, 0x42, 0xe8, 0x3b, 0x00, 0x63, 0xcc, 0x4a, 0x83, 0x36, 0xc0, 0x27, 0xac, 0x7c, 0x9b, 0xe5,
	0x66, 0x2a, 0xd0, 0x0b, 0xb0, 0x7d, 0x2a, 0x17, 0xba, 0xba, 0xa6, 0x75, 0x7f, 0xf9, 0x58, 0x75,
	0x92, 0xe3, 0xa6, 0xd9, 0xf8, 0x63, 0x7a, 0x1f, 0xfc, 0x80, 0xcd, 0xde, 0xcc, 0xa3, 0xd2, 0x1e,
	0xc1, 0x39, 0xc0, 0x3d, 0xe7, 0x61, 0xfa, 0x04, 0xd4, 0x79, 0x66, 0x10, 0xfc, 0x6e, 0x6b, 0xdc,
	0x65, 0xb7, 0x3e, 0x83, 0xc3, 0xfb, 0x25, 0xff, 0x40, 0x45, 0xfc, 0xf2, 0xe2, 0x08, 0xbf, 0x4d,
	0x2d, 0x34, 0xee, 0x58, 0xdc, 0x42, 0x63, 0xfe, 0x4a, 0x8e, 0x3f, 0x63, 0xff, 0x25, 0x6f, 0x09,
	0xff, 0x64, 0xbe, 0x4d, 0xa6, 0xde, 0x64, 0x1e, 0x11, 0xd6, 0x0f, 0x4a, 0x23, 0x0e, 0xcd, 0x05,
	0x50, 0xc7, 0x50, 0x2e, 0x73, 0xce, 0xe1, 0xcc, 0xf1, 0x6f, 0x1d, 0xee, 0xf7, 0xf8, 0xcb, 0x43,
	0x8d, 0xaa, 0xe4, 0xb6, 0x18, 0x8e, 0x03, 0x79, 0xa5, 0x39, 0x6f, 0xde, 0xc7, 0x57, 0xc0, 0x76,
	0x73, 0xd8, 0x8e, 0xf9, 0xe6, 0xa5, 0xcd, 0x01, 0x25, 0x03, 0xf1, 0x94, 0x87, 0xd0, 0x3e, 0x95,
	0x8b, 0x22, 0xbe, 0x10, 0xc8, 0xcb, 0xb4, 0x5e, 0x6b, 0xd4, 0x2a, 0xb6, 0x18, 0xfe, 0xd3, 0x32,
	0xad, 0xd4, 0x10, 0xfe, 0x47, 0x2b, 0x6d, 0x80, 0x86, 0xf6, 0x26, 0x69, 0x94, 0x41, 0xd0, 0xb7,
	0x60, 0x27, 0xf6, 0xa6, 0x87, 0x51, 0xff, 0xd7, 0x6f, 0x7b, 0xc6, 0x00, 0xdd, 0x34, 0xf7, 0xeb,
	0x99, 0xb9, 0xcd, 0x99, 0x1f, 0x7e, 0xd4, 0x80, 0xa3, 0x51, 0x7f, 0x3c, 0x19, 0x0d, 0x5f, 0xdf,
	0x8e, 0x9b, 0x16, 0xfa, 0x02, 0x1a, 0x69, 0xf8, 0x6a, 0x78, 0x37, 0x68, 0x56, 0xd0, 0x09, 0xd4,
	0x35, 0x34, 0xb8, 0x1b, 0x5c, 0xdf, 0x8e, 0x9b, 0x55, 0x74, 0x0a, 0x27, 0x19, 0x40, 0x67, 0x1d,
	0xa0, 0x63, 0xb0, 0x15, 0x78, 0x7d, 0x73, 0x3d, 0x68, 0xd6, 0x5e, 0x56, 0xae, 0xaa, 0xce, 0x67,
	0x8e, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0xd1, 0x1f, 0x3f, 0x1a, 0x0c, 0x00, 0x00,
}
