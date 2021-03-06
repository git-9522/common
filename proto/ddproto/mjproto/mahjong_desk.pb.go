// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mahjong_desk.proto

package mjproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from base.proto

// Ignoring public import of Heartbeat from base.proto

// Ignoring public import of WeixinInfo from base.proto

// Ignoring public import of CardInfo from base.proto

// Ignoring public import of PlayOptions from base.proto

// Ignoring public import of ChangShaPlayOptions from base.proto

// Ignoring public import of BaiShanPlayOptions from base.proto

// Ignoring public import of ZhuanZhuanPlayOptions from base.proto

// Ignoring public import of LiuZhouPlayOptions from base.proto

// Ignoring public import of HaiNanPlayOptions from base.proto

// Ignoring public import of ChangBaiPlayOptions from base.proto

// Ignoring public import of RoomTypeInfo from base.proto

// Ignoring public import of ComposeCard from base.proto

// Ignoring public import of PlayerCard from base.proto

// Ignoring public import of GangBean from base.proto

// Ignoring public import of PlayerInfo from base.proto

// Ignoring public import of DeskGameInfo from base.proto

// Ignoring public import of EProtoId from base.proto

// Ignoring public import of ErrorCode from base.proto

// Ignoring public import of MJOption from base.proto

// Ignoring public import of MJRoomType from base.proto

// Ignoring public import of MahjongColor from base.proto

// Ignoring public import of GangType from base.proto

// Ignoring public import of ComposeCardType from base.proto

// Ignoring public import of HuType from base.proto

// Ignoring public import of PaiType from base.proto

// Ignoring public import of COMMON_ENUM_ROOMCARD_BILL_TYPE from base.proto

// Ignoring public import of MJUserGameStatus from base.proto

// Ignoring public import of DeskGameStatus from base.proto

// Ignoring public import of CHANGBAI_PREOPENING_PLAYOPTIONS from base.proto

// 房主解散房间(未开局)
type Game_DissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_DissolveDesk) Reset()                    { *m = Game_DissolveDesk{} }
func (m *Game_DissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_DissolveDesk) ProtoMessage()               {}
func (*Game_DissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Game_DissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_DissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_DissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

// 解散房间回复
type Game_AckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	PassWord         *string      `protobuf:"bytes,4,opt,name=passWord" json:"passWord,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckDissolveDesk) Reset()                    { *m = Game_AckDissolveDesk{} }
func (m *Game_AckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_AckDissolveDesk) ProtoMessage()               {}
func (*Game_AckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Game_AckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Game_AckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *Game_AckDissolveDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

// 申请解散房间(游戏中)
type Game_ReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_ReqDissolveDesk) Reset()                    { *m = Game_ReqDissolveDesk{} }
func (m *Game_ReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_ReqDissolveDesk) ProtoMessage()               {}
func (*Game_ReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Game_ReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_ReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type Game_AckReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserIdAgree      []uint32     `protobuf:"varint,2,rep,name=userIdAgree" json:"userIdAgree,omitempty"`
	UserIdWait       []uint32     `protobuf:"varint,3,rep,name=userIdWait" json:"userIdWait,omitempty"`
	UserIdDisagree   []uint32     `protobuf:"varint,4,rep,name=userIdDisagree" json:"userIdDisagree,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckReqDissolveDesk) Reset()                    { *m = Game_AckReqDissolveDesk{} }
func (m *Game_AckReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*Game_AckReqDissolveDesk) ProtoMessage()               {}
func (*Game_AckReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Game_AckReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdAgree() []uint32 {
	if m != nil {
		return m.UserIdAgree
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdWait() []uint32 {
	if m != nil {
		return m.UserIdWait
	}
	return nil
}

func (m *Game_AckReqDissolveDesk) GetUserIdDisagree() []uint32 {
	if m != nil {
		return m.UserIdDisagree
	}
	return nil
}

// 准备游戏
type Game_Ready struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_Ready) Reset()                    { *m = Game_Ready{} }
func (m *Game_Ready) String() string            { return proto.CompactTextString(m) }
func (*Game_Ready) ProtoMessage()               {}
func (*Game_Ready) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Game_Ready) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_Ready) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type Game_AckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Msg              *string      `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	UserId           *uint32      `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Game_AckReady) Reset()                    { *m = Game_AckReady{} }
func (m *Game_AckReady) String() string            { return proto.CompactTextString(m) }
func (*Game_AckReady) ProtoMessage()               {}
func (*Game_AckReady) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Game_AckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_AckReady) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

func (m *Game_AckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

type BirdInfo struct {
	BirdPai          *CardInfo `protobuf:"bytes,1,opt,name=birdPai" json:"birdPai,omitempty"`
	ZhuaUser         *uint32   `protobuf:"varint,2,opt,name=zhuaUser" json:"zhuaUser,omitempty"`
	BirdUser         *uint32   `protobuf:"varint,3,opt,name=birdUser" json:"birdUser,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BirdInfo) Reset()                    { *m = BirdInfo{} }
func (m *BirdInfo) String() string            { return proto.CompactTextString(m) }
func (*BirdInfo) ProtoMessage()               {}
func (*BirdInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *BirdInfo) GetBirdPai() *CardInfo {
	if m != nil {
		return m.BirdPai
	}
	return nil
}

func (m *BirdInfo) GetZhuaUser() uint32 {
	if m != nil && m.ZhuaUser != nil {
		return *m.ZhuaUser
	}
	return 0
}

func (m *BirdInfo) GetBirdUser() uint32 {
	if m != nil && m.BirdUser != nil {
		return *m.BirdUser
	}
	return 0
}

// 赢牌信息：谁赢了多少
type WinCoinInfo struct {
	NickName         *string     `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32     `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	WinCoin          *int64      `protobuf:"varint,3,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64      `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	CardTitle        *string     `protobuf:"bytes,5,opt,name=cardTitle" json:"cardTitle,omitempty"`
	Cards            *PlayerCard `protobuf:"bytes,6,opt,name=cards" json:"cards,omitempty"`
	IsDealer         *bool       `protobuf:"varint,7,opt,name=isDealer" json:"isDealer,omitempty"`
	HuCount          *int32      `protobuf:"varint,8,opt,name=huCount" json:"huCount,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *WinCoinInfo) Reset()                    { *m = WinCoinInfo{} }
func (m *WinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*WinCoinInfo) ProtoMessage()               {}
func (*WinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *WinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *WinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *WinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WinCoinInfo) GetCardTitle() string {
	if m != nil && m.CardTitle != nil {
		return *m.CardTitle
	}
	return ""
}

func (m *WinCoinInfo) GetCards() *PlayerCard {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *WinCoinInfo) GetIsDealer() bool {
	if m != nil && m.IsDealer != nil {
		return *m.IsDealer
	}
	return false
}

func (m *WinCoinInfo) GetHuCount() int32 {
	if m != nil && m.HuCount != nil {
		return *m.HuCount
	}
	return 0
}

type EndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	CountHu          *int32  `protobuf:"varint,6,opt,name=countHu" json:"countHu,omitempty"`
	CountZiMo        *int32  `protobuf:"varint,7,opt,name=countZiMo" json:"countZiMo,omitempty"`
	CountDianPao     *int32  `protobuf:"varint,8,opt,name=countDianPao" json:"countDianPao,omitempty"`
	CountAnGang      *int32  `protobuf:"varint,9,opt,name=countAnGang" json:"countAnGang,omitempty"`
	CountMingGang    *int32  `protobuf:"varint,10,opt,name=countMingGang" json:"countMingGang,omitempty"`
	CountDianGang    *int32  `protobuf:"varint,11,opt,name=countDianGang" json:"countDianGang,omitempty"`
	CountChaJiao     *int32  `protobuf:"varint,12,opt,name=countChaJiao" json:"countChaJiao,omitempty"`
	BestGunner       *bool   `protobuf:"varint,13,opt,name=bestGunner" json:"bestGunner,omitempty"`
	CountLianZhuang  *int32  `protobuf:"varint,14,opt,name=countLianZhuang" json:"countLianZhuang,omitempty"`
	CountZhaMa       *int32  `protobuf:"varint,15,opt,name=countZhaMa" json:"countZhaMa,omitempty"`
	CountChi         *int32  `protobuf:"varint,16,opt,name=countChi" json:"countChi,omitempty"`
	CountPeng        *int32  `protobuf:"varint,17,opt,name=countPeng" json:"countPeng,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EndLotteryInfo) Reset()                    { *m = EndLotteryInfo{} }
func (m *EndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*EndLotteryInfo) ProtoMessage()               {}
func (*EndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *EndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *EndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *EndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *EndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *EndLotteryInfo) GetCountHu() int32 {
	if m != nil && m.CountHu != nil {
		return *m.CountHu
	}
	return 0
}

func (m *EndLotteryInfo) GetCountZiMo() int32 {
	if m != nil && m.CountZiMo != nil {
		return *m.CountZiMo
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianPao() int32 {
	if m != nil && m.CountDianPao != nil {
		return *m.CountDianPao
	}
	return 0
}

func (m *EndLotteryInfo) GetCountAnGang() int32 {
	if m != nil && m.CountAnGang != nil {
		return *m.CountAnGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountMingGang() int32 {
	if m != nil && m.CountMingGang != nil {
		return *m.CountMingGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountDianGang() int32 {
	if m != nil && m.CountDianGang != nil {
		return *m.CountDianGang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountChaJiao() int32 {
	if m != nil && m.CountChaJiao != nil {
		return *m.CountChaJiao
	}
	return 0
}

func (m *EndLotteryInfo) GetBestGunner() bool {
	if m != nil && m.BestGunner != nil {
		return *m.BestGunner
	}
	return false
}

func (m *EndLotteryInfo) GetCountLianZhuang() int32 {
	if m != nil && m.CountLianZhuang != nil {
		return *m.CountLianZhuang
	}
	return 0
}

func (m *EndLotteryInfo) GetCountZhaMa() int32 {
	if m != nil && m.CountZhaMa != nil {
		return *m.CountZhaMa
	}
	return 0
}

func (m *EndLotteryInfo) GetCountChi() int32 {
	if m != nil && m.CountChi != nil {
		return *m.CountChi
	}
	return 0
}

func (m *EndLotteryInfo) GetCountPeng() int32 {
	if m != nil && m.CountPeng != nil {
		return *m.CountPeng
	}
	return 0
}

// 本局结果(广播)
type Game_SendCurrentResult struct {
	Header           *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*WinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	BridInfo         []*BirdInfo    `protobuf:"bytes,3,rep,name=bridInfo" json:"bridInfo,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Game_SendCurrentResult) Reset()                    { *m = Game_SendCurrentResult{} }
func (m *Game_SendCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*Game_SendCurrentResult) ProtoMessage()               {}
func (*Game_SendCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Game_SendCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_SendCurrentResult) GetWinCoinInfo() []*WinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

func (m *Game_SendCurrentResult) GetBridInfo() []*BirdInfo {
	if m != nil {
		return m.BridInfo
	}
	return nil
}

// 牌局结束(广播)
type Game_SendEndLottery struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*EndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Game_SendEndLottery) Reset()                    { *m = Game_SendEndLottery{} }
func (m *Game_SendEndLottery) String() string            { return proto.CompactTextString(m) }
func (*Game_SendEndLottery) ProtoMessage()               {}
func (*Game_SendEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *Game_SendEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Game_SendEndLottery) GetCoinInfo() []*EndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Game_DissolveDesk)(nil), "mjproto.game_DissolveDesk")
	proto.RegisterType((*Game_AckDissolveDesk)(nil), "mjproto.game_AckDissolveDesk")
	proto.RegisterType((*Game_ReqDissolveDesk)(nil), "mjproto.game_ReqDissolveDesk")
	proto.RegisterType((*Game_AckReqDissolveDesk)(nil), "mjproto.game_AckReqDissolveDesk")
	proto.RegisterType((*Game_Ready)(nil), "mjproto.game_Ready")
	proto.RegisterType((*Game_AckReady)(nil), "mjproto.game_AckReady")
	proto.RegisterType((*BirdInfo)(nil), "mjproto.BirdInfo")
	proto.RegisterType((*WinCoinInfo)(nil), "mjproto.WinCoinInfo")
	proto.RegisterType((*EndLotteryInfo)(nil), "mjproto.EndLotteryInfo")
	proto.RegisterType((*Game_SendCurrentResult)(nil), "mjproto.game_SendCurrentResult")
	proto.RegisterType((*Game_SendEndLottery)(nil), "mjproto.game_SendEndLottery")
}

func init() { proto.RegisterFile("mahjong_desk.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x5f, 0x6b, 0x13, 0x4f,
	0x14, 0xfd, 0x6d, 0x93, 0x34, 0x9b, 0x9b, 0xa6, 0x7f, 0xa6, 0xa5, 0x1d, 0xca, 0x0f, 0x09, 0x8b,
	0x48, 0x44, 0xed, 0x83, 0x82, 0xef, 0x35, 0x91, 0xb6, 0xd2, 0x6a, 0x18, 0x95, 0x40, 0x11, 0xca,
	0x24, 0x3b, 0x6e, 0xa6, 0x49, 0x66, 0xdb, 0x99, 0x5d, 0x6b, 0xfd, 0x14, 0x7e, 0x06, 0xdf, 0x7d,
	0xf0, 0x8b, 0xf9, 0x19, 0x64, 0xee, 0xfe, 0xc9, 0xa4, 0xe2, 0x43, 0xb1, 0xf8, 0x12, 0xe6, 0x9c,
	0x39, 0x7b, 0xcf, 0xbd, 0x27, 0x7b, 0x17, 0xc8, 0x8c, 0x8f, 0xcf, 0x63, 0x15, 0x9d, 0x85, 0xc2,
	0x4c, 0xf6, 0x2e, 0x74, 0x9c, 0xc4, 0xa4, 0x3e, 0x3b, 0xc7, 0xc3, 0x2e, 0x0c, 0xb9, 0x11, 0x19,
	0x19, 0x5c, 0xc2, 0x46, 0xc4, 0x67, 0xe2, 0xac, 0x27, 0x8d, 0x89, 0xa7, 0x9f, 0x44, 0x4f, 0x98,
	0x09, 0x79, 0x0c, 0xcb, 0x63, 0xc1, 0x43, 0xa1, 0xa9, 0xd7, 0xf6, 0x3a, 0xcd, 0xa7, 0x5b, 0x7b,
	0xf9, 0xa3, 0x7b, 0x7d, 0xfb, 0x7b, 0x88, 0x77, 0x2c, 0xd7, 0x90, 0x6d, 0x58, 0x4e, 0x8d, 0xd0,
	0x47, 0x21, 0x5d, 0x6a, 0x7b, 0x9d, 0x16, 0xcb, 0x91, 0xe5, 0xad, 0xfb, 0x51, 0x48, 0x2b, 0x6d,
	0xaf, 0x53, 0x63, 0x39, 0x0a, 0xbe, 0x7a, 0xb0, 0x85, 0x9e, 0xfb, 0xa3, 0xc9, 0xbf, 0xb3, 0x25,
	0xbb, 0xe0, 0x5f, 0x70, 0x63, 0x06, 0xb1, 0x0e, 0x69, 0xb5, 0xed, 0x75, 0x1a, 0xac, 0xc4, 0xc1,
	0x87, 0xbc, 0x23, 0x26, 0x2e, 0xef, 0xbe, 0xa3, 0xe0, 0x87, 0x07, 0x3b, 0xc5, 0xc0, 0x7f, 0xe7,
	0xd0, 0x86, 0x66, 0x56, 0x73, 0x3f, 0xd2, 0x42, 0xd0, 0xa5, 0x76, 0xa5, 0xd3, 0x62, 0x2e, 0x45,
	0xee, 0x01, 0x64, 0x70, 0xc0, 0x65, 0x42, 0x2b, 0x28, 0x70, 0x18, 0xf2, 0x00, 0x56, 0x33, 0xd4,
	0x93, 0x86, 0x63, 0x91, 0x2a, 0x6a, 0x6e, 0xb0, 0x01, 0x03, 0xc8, 0x13, 0xe1, 0xe1, 0xf5, 0x1d,
	0xe5, 0x10, 0x41, 0x6b, 0x1e, 0xc3, 0xed, 0xcb, 0xae, 0x43, 0x65, 0x66, 0x22, 0xac, 0xd9, 0x60,
	0xf6, 0xe8, 0x18, 0x55, 0x16, 0x8c, 0x62, 0xf0, 0x5f, 0x48, 0x1d, 0x1e, 0xa9, 0x8f, 0x31, 0x79,
	0x04, 0xf5, 0xa1, 0xd4, 0x61, 0x9f, 0xcb, 0xdc, 0x64, 0xa3, 0x34, 0xe9, 0xf2, 0x4c, 0xc3, 0x0a,
	0x85, 0x7d, 0x47, 0xbe, 0x8c, 0x53, 0xfe, 0xde, 0x08, 0x9d, 0xf7, 0x5e, 0x62, 0x7b, 0x67, 0x65,
	0x78, 0x97, 0xd9, 0x95, 0x38, 0xf8, 0xe9, 0x41, 0x73, 0x20, 0x55, 0x37, 0x96, 0x0a, 0x4d, 0x77,
	0xc1, 0x57, 0x72, 0x34, 0x79, 0xcd, 0x67, 0x02, 0x5d, 0x1b, 0xac, 0xc4, 0x7f, 0x7c, 0x6f, 0x29,
	0xd4, 0xaf, 0xb2, 0x12, 0x58, 0xbe, 0xc2, 0x0a, 0x48, 0x08, 0x54, 0x47, 0x96, 0xae, 0x22, 0x8d,
	0x67, 0xf2, 0x3f, 0x34, 0x46, 0x5c, 0x87, 0xef, 0x64, 0x32, 0x15, 0xb4, 0x86, 0x16, 0x73, 0x82,
	0x3c, 0x84, 0x9a, 0x05, 0x86, 0x2e, 0xe3, 0xc8, 0x9b, 0xf3, 0x5c, 0xa7, 0xfc, 0x5a, 0x68, 0x3b,
	0x38, 0xcb, 0x14, 0xb6, 0x55, 0x69, 0x7a, 0x82, 0x4f, 0x85, 0xa6, 0xf5, 0xb6, 0xd7, 0xf1, 0x59,
	0x89, 0x6d, 0x4b, 0xe3, 0xb4, 0x1b, 0xa7, 0x2a, 0xa1, 0x3e, 0xee, 0x52, 0x01, 0x83, 0x6f, 0x55,
	0x58, 0x7d, 0xa9, 0xc2, 0xe3, 0x38, 0x49, 0x84, 0xbe, 0xc6, 0x99, 0xe7, 0x73, 0x79, 0x0b, 0x73,
	0xb9, 0x59, 0x2c, 0xfd, 0x9e, 0xc5, 0x50, 0x46, 0x83, 0x7c, 0x64, 0x9f, 0xe5, 0xc8, 0x1a, 0x4b,
	0xf3, 0xe6, 0x4a, 0x09, 0x8d, 0x43, 0xfb, 0xac, 0x80, 0x6e, 0x4a, 0xb5, 0xc5, 0x94, 0x28, 0xd4,
	0x47, 0xb6, 0xb7, 0xc3, 0x14, 0xa7, 0xae, 0xb1, 0x02, 0x62, 0x56, 0xf6, 0x78, 0x2a, 0x4f, 0x62,
	0x9c, 0xb1, 0xc6, 0xe6, 0x04, 0x09, 0x60, 0x05, 0x41, 0x4f, 0x72, 0xd5, 0xe7, 0x71, 0x3e, 0xe9,
	0x02, 0x67, 0xf7, 0x0e, 0xf1, 0xbe, 0x3a, 0xe0, 0x2a, 0xa2, 0x0d, 0x94, 0xb8, 0x14, 0xb9, 0x0f,
	0x2d, 0x84, 0x27, 0x52, 0x45, 0xa8, 0x01, 0xd4, 0x2c, 0x92, 0xa5, 0xca, 0xd6, 0x45, 0x55, 0xd3,
	0x51, 0x15, 0x64, 0xd9, 0x51, 0x77, 0xcc, 0x5f, 0x49, 0x1e, 0xd3, 0x15, 0xa7, 0xa3, 0x9c, 0xb3,
	0x7b, 0x3e, 0x14, 0x26, 0x39, 0x48, 0x95, 0x0d, 0xa9, 0x85, 0x21, 0x39, 0x0c, 0xe9, 0xc0, 0x1a,
	0xea, 0x8f, 0x25, 0x57, 0xa7, 0xe3, 0xd4, 0x7a, 0xad, 0x62, 0x99, 0x9b, 0xb4, 0xad, 0x94, 0x85,
	0x31, 0xe6, 0x27, 0x9c, 0xae, 0xa1, 0xc8, 0x61, 0xec, 0xff, 0x97, 0x3b, 0x4b, 0xba, 0x8e, 0xb7,
	0x25, 0x2e, 0x93, 0xed, 0x0b, 0x15, 0xd1, 0x0d, 0x27, 0x59, 0x4b, 0x04, 0xdf, 0x3d, 0xd8, 0xc6,
	0x85, 0x7f, 0x2b, 0x54, 0xd8, 0x4d, 0xb5, 0x16, 0x2a, 0x61, 0xc2, 0xa4, 0xd3, 0xe4, 0x96, 0x9b,
	0xff, 0x1c, 0x9a, 0x57, 0xf3, 0xed, 0xc2, 0xcf, 0x9e, 0xfb, 0x88, 0xb3, 0x79, 0xcc, 0x15, 0x92,
	0x27, 0xe0, 0x0f, 0xb5, 0xc4, 0x1d, 0xc7, 0x4f, 0xa1, 0xbb, 0xfc, 0xc5, 0x07, 0x82, 0x95, 0x92,
	0xe0, 0x33, 0x6c, 0x96, 0xed, 0xce, 0x5f, 0xee, 0x5b, 0xf6, 0xfa, 0xcc, 0xc6, 0xb5, 0xd0, 0xe8,
	0x4e, 0xa9, 0x5f, 0xdc, 0x18, 0x56, 0x0a, 0xfb, 0xff, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x53,
	0xc2, 0x02, 0xb5, 0xb1, 0x07, 0x00, 0x00,
}
