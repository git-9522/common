// Code generated by protoc-gen-go.
// source: ddz_desk.proto
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

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

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

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 房主解散房间(未开局)
type DdzReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqDissolveDesk) Reset()                    { *m = DdzReqDissolveDesk{} }
func (m *DdzReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzReqDissolveDesk) ProtoMessage()               {}
func (*DdzReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *DdzReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type DdzAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckDissolveDesk) Reset()                    { *m = DdzAckDissolveDesk{} }
func (m *DdzAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*DdzAckDissolveDesk) ProtoMessage()               {}
func (*DdzAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *DdzAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *DdzAckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 准备游戏
type DdzReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	IsShowHandPokers *bool        `protobuf:"varint,3,opt,name=isShowHandPokers" json:"isShowHandPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzReqReady) Reset()                    { *m = DdzReqReady{} }
func (m *DdzReqReady) String() string            { return proto.CompactTextString(m) }
func (*DdzReqReady) ProtoMessage()               {}
func (*DdzReqReady) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *DdzReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzReqReady) GetIsShowHandPokers() bool {
	if m != nil && m.IsShowHandPokers != nil {
		return *m.IsShowHandPokers
	}
	return false
}

// 准备游戏的结果
type DdzAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DdzAckReady) Reset()                    { *m = DdzAckReady{} }
func (m *DdzAckReady) String() string            { return proto.CompactTextString(m) }
func (*DdzAckReady) ProtoMessage()               {}
func (*DdzAckReady) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *DdzAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzAckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *DdzAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type DdzBaseWinCoinInfo struct {
	NickName         *string `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BaseValue        *int32  `protobuf:"varint,3,opt,name=baseValue" json:"baseValue,omitempty"`
	WinCoin          *int64  `protobuf:"varint,4,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64  `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	IsDiZhu          *bool   `protobuf:"varint,6,opt,name=isDiZhu" json:"isDiZhu,omitempty"`
	Rate             *int32  `protobuf:"varint,7,opt,name=rate" json:"rate,omitempty"`
	Description      *string `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
	Bomb             *int32  `protobuf:"varint,9,opt,name=bomb" json:"bomb,omitempty"`
	IsSpring         *bool   `protobuf:"varint,10,opt,name=isSpring" json:"isSpring,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzBaseWinCoinInfo) Reset()                    { *m = DdzBaseWinCoinInfo{} }
func (m *DdzBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseWinCoinInfo) ProtoMessage()               {}
func (*DdzBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *DdzBaseWinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *DdzBaseWinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetIsDiZhu() bool {
	if m != nil && m.IsDiZhu != nil {
		return *m.IsDiZhu
	}
	return false
}

func (m *DdzBaseWinCoinInfo) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *DdzBaseWinCoinInfo) GetBomb() int32 {
	if m != nil && m.Bomb != nil {
		return *m.Bomb
	}
	return 0
}

func (m *DdzBaseWinCoinInfo) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

// 本局结果(广播)
type DdzBaCurrentResult struct {
	Header           *ProtoHeader          `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*DdzBaseWinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *DdzBaCurrentResult) Reset()                    { *m = DdzBaCurrentResult{} }
func (m *DdzBaCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*DdzBaCurrentResult) ProtoMessage()               {}
func (*DdzBaCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *DdzBaCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBaCurrentResult) GetWinCoinInfo() []*DdzBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

type DdzBaseEndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	MaxWinCoin       *int32  `protobuf:"varint,6,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	CountBomb        *int32  `protobuf:"varint,7,opt,name=countBomb" json:"countBomb,omitempty"`
	CountWin         *int32  `protobuf:"varint,8,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32  `protobuf:"varint,9,opt,name=countLose" json:"countLose,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DdzBaseEndLotteryInfo) Reset()                    { *m = DdzBaseEndLotteryInfo{} }
func (m *DdzBaseEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*DdzBaseEndLotteryInfo) ProtoMessage()               {}
func (*DdzBaseEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *DdzBaseEndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *DdzBaseEndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *DdzBaseEndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *DdzBaseEndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *DdzBaseEndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *DdzBaseEndLotteryInfo) GetMaxWinCoin() int32 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

func (m *DdzBaseEndLotteryInfo) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *DdzBaseEndLotteryInfo) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *DdzBaseEndLotteryInfo) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

// 牌局结束(广播)
type DdzBcEndLottery struct {
	Header           *ProtoHeader             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*DdzBaseEndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *DdzBcEndLottery) Reset()                    { *m = DdzBcEndLottery{} }
func (m *DdzBcEndLottery) String() string            { return proto.CompactTextString(m) }
func (*DdzBcEndLottery) ProtoMessage()               {}
func (*DdzBcEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *DdzBcEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DdzBcEndLottery) GetCoinInfo() []*DdzBaseEndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*DdzReqDissolveDesk)(nil), "ddproto.ddz_req_dissolveDesk")
	proto.RegisterType((*DdzAckDissolveDesk)(nil), "ddproto.ddz_ack_dissolveDesk")
	proto.RegisterType((*DdzReqReady)(nil), "ddproto.ddz_req_ready")
	proto.RegisterType((*DdzAckReady)(nil), "ddproto.ddz_ack_ready")
	proto.RegisterType((*DdzBaseWinCoinInfo)(nil), "ddproto.ddz_base_winCoinInfo")
	proto.RegisterType((*DdzBaCurrentResult)(nil), "ddproto.ddz_ba_currentResult")
	proto.RegisterType((*DdzBaseEndLotteryInfo)(nil), "ddproto.ddz_base_endLotteryInfo")
	proto.RegisterType((*DdzBcEndLottery)(nil), "ddproto.ddz_bc_endLottery")
}

var fileDescriptor11 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xc9, 0xd2, 0xb5, 0xc9, 0x2b, 0x1b, 0x5b, 0x36, 0x09, 0x0b, 0x09, 0x69, 0x8a, 0x38,
	0x70, 0xea, 0x61, 0x7f, 0xc2, 0xd8, 0xa1, 0x93, 0x2a, 0xa8, 0x86, 0x44, 0x25, 0x2e, 0x91, 0x6b,
	0x9b, 0xd6, 0x6a, 0x63, 0x07, 0xdb, 0xa1, 0x8c, 0xff, 0x8c, 0x0b, 0x7f, 0x1b, 0xf6, 0x4b, 0x53,
	0x52, 0x8d, 0x43, 0x41, 0xbb, 0x54, 0x7d, 0xbf, 0xbe, 0xef, 0x7d, 0x3f, 0x0e, 0x9c, 0x72, 0xfe,
	0xa3, 0xe0, 0xc2, 0xae, 0x46, 0x95, 0xd1, 0x4e, 0x67, 0x03, 0xce, 0xf1, 0xcf, 0xab, 0x0b, 0xa6,
	0xcb, 0x52, 0xab, 0x82, 0xad, 0xa5, 0x50, 0xae, 0xa9, 0xe6, 0x13, 0xb8, 0x0c, 0xfd, 0x46, 0x7c,
	0x2d, 0xb8, 0xb4, 0x56, 0xaf, 0xbf, 0x89, 0x5b, 0x3f, 0x9b, 0xbd, 0x81, 0xfe, 0x52, 0x50, 0x2e,
	0x0c, 0x89, 0xae, 0xa2, 0xb7, 0xc3, 0xeb, 0xcb, 0xd1, 0x56, 0x66, 0x34, 0x0d, 0xbf, 0x63, 0xac,
	0x65, 0xa7, 0xd0, 0xaf, 0xad, 0x30, 0x77, 0x9c, 0x1c, 0xf9, 0xae, 0x93, 0x5c, 0x35, 0x6a, 0x94,
	0xad, 0x9e, 0x40, 0x2d, 0xc4, 0xc1, 0x87, 0x8f, 0x63, 0x1f, 0x1f, 0x67, 0x67, 0x90, 0x54, 0xd4,
	0xda, 0x99, 0x36, 0x9c, 0xf4, 0x7c, 0x26, 0xcd, 0x0b, 0x38, 0x69, 0xaf, 0x37, 0x5e, 0xe3, 0xe1,
	0x3f, 0x17, 0x11, 0x38, 0x93, 0xf6, 0xe3, 0x52, 0x6f, 0xc6, 0x54, 0xf1, 0xa9, 0x5e, 0x09, 0x63,
	0x71, 0x65, 0x92, 0xdf, 0x37, 0x0b, 0x82, 0xa1, 0x7f, 0x59, 0x30, 0x84, 0xb8, 0xb4, 0x0b, 0x54,
	0x4f, 0x3b, 0xdb, 0x62, 0x84, 0xf4, 0x2b, 0x6a, 0x28, 0xcd, 0xa9, 0x15, 0xc5, 0x46, 0xaa, 0x77,
	0x5a, 0xaa, 0x3b, 0xf5, 0x45, 0x07, 0x7f, 0x4a, 0xb2, 0xd5, 0x7b, 0x5a, 0x0a, 0x54, 0x4f, 0x1f,
	0x1d, 0x7a, 0x0e, 0x69, 0x98, 0xfa, 0x44, 0xd7, 0xb5, 0xd8, 0x42, 0x79, 0x01, 0x83, 0xad, 0x06,
	0x32, 0x89, 0xb3, 0xe7, 0xd0, 0x63, 0x21, 0x3a, 0xc6, 0xc8, 0x97, 0xa5, 0xbd, 0x95, 0x9f, 0x97,
	0x35, 0xe9, 0x07, 0x47, 0xa1, 0x6c, 0xa8, 0x13, 0x64, 0x80, 0xd3, 0x17, 0x30, 0xf4, 0x88, 0x99,
	0x91, 0x95, 0x93, 0x5a, 0x91, 0x04, 0xb7, 0xfa, 0x96, 0xb9, 0x2e, 0xe7, 0x24, 0x6d, 0xa9, 0x7b,
	0x38, 0x95, 0x91, 0x6a, 0x41, 0x00, 0xa1, 0x54, 0xed, 0xfd, 0x05, 0xab, 0x8d, 0xf1, 0xdf, 0xd2,
	0xbd, 0xb0, 0xf5, 0xda, 0x1d, 0xc8, 0xe6, 0x1a, 0x86, 0x1d, 0xd3, 0xde, 0x58, 0xec, 0x5b, 0x5f,
	0xef, 0x5a, 0xff, 0x46, 0x26, 0xff, 0x19, 0xc1, 0xcb, 0x5d, 0x41, 0x28, 0x3e, 0xd1, 0xce, 0x09,
	0xf3, 0x80, 0xd4, 0xfe, 0x30, 0x8a, 0x90, 0x51, 0x97, 0xe2, 0xee, 0x01, 0xe6, 0x72, 0x31, 0xf3,
	0x4c, 0xf0, 0x51, 0x1b, 0x26, 0x1f, 0x36, 0xca, 0x1f, 0xda, 0x6b, 0x13, 0x2d, 0xc3, 0x86, 0x5a,
	0x06, 0x50, 0xd2, 0xef, 0xb3, 0x6d, 0xae, 0x8f, 0x1c, 0x3c, 0x7b, 0xa6, 0x6b, 0xe5, 0x6e, 0x02,
	0x9a, 0x41, 0x8b, 0x06, 0x53, 0x41, 0x3a, 0xd9, 0x6b, 0x9a, 0x68, 0x2b, 0x1a, 0x7e, 0x79, 0x09,
	0xe7, 0x78, 0x3a, 0xeb, 0x1c, 0x7e, 0x30, 0xaa, 0x84, 0xed, 0x73, 0xba, 0x7a, 0xcc, 0x69, 0x1f,
	0xc7, 0xcd, 0xd1, 0x38, 0x9e, 0x3e, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf0, 0x02, 0x9b,
	0x05, 0x04, 0x00, 0x00,
}
