// Code generated by protoc-gen-go.
// source: bainiu_play.proto
// DO NOT EDIT!

/*
Package ddproto is a generated protocol buffer package.

It is generated from these files:
	bainiu_play.proto
	common_client.proto
	common_client_award.proto
	common_client_pay.proto
	common_enum.proto
	common_mj.proto
	common_mj_play.proto
	common_pdk.proto
	common_server.proto
	common_server_award.proto
	common_server_config.proto
	common_server_model.proto
	common_server_msg.proto
	common_server_poker.proto
	common_server_redis.proto
	common_server_task.proto
	ddz_base.proto
	ddz_desk.proto
	ddz_hall.proto
	ddz_play.proto
	ddz_server.proto
	erddz_base.proto
	erddz_desk.proto
	erddz_hall.proto
	erddz_play.proto
	hall.proto
	hall_data.proto
	hall_playback.proto
	hall_playback_pdk.proto
	hot_update.proto
	mj_baishan_base.proto
	niuniu_base.proto
	niuniu_play.proto
	niuniu_server.proto
	pdk_base.proto
	pdk_desk.proto
	pdk_hall.proto
	pdk_play.proto
	pdk_server.proto
	pez_base.proto
	pez_desk.proto
	pez_hall.proto
	pez_play.proto
	rpc_hall.proto
	zjh_base.proto
	zjh_desk.proto
	zjh_hall.proto
	zjh_play.proto
	zjh_serever.proto

It has these top-level messages:
	BainiuAreaInfo
	BainiuClientDesk
	BainiuClientUser
	BainiuEnterDeskReq
	BainiuEnterDeskAck
	BainiuYazhuOt
	BainiuYazhuReq
	BainiuYazhuAck
	BainiuYazhuBc
	BainiuFapaiBc
	BainiuWinUserItem
	BainiuBipaiBc
	ProtoHeader
	Heartbeat
	ServerInfo
	QuickConn
	AckQuickConn
	WeixinInfo
	CommonReqReg
	CommonReqRegViaInput
	CommonAckReg
	CommonReqGameLogin
	CommonReqGameLoginViaInput
	CommonAckGameLogin
	CommonReqQrLogin
	CommonAckQrLogin
	CommonReqQrWxInfo
	CommonAckQrWxInfo
	CommonAckReconnect
	CommonReqReconnect
	CommonReqGameState
	CommonAckGameState
	CommonReqLogout
	CommonAckLogout
	CommonReqFeedback
	ClientBasePoker
	CommonReqMessage
	CommonBcMessage
	CommonReqNotice
	CommonAckNotice
	CommonReqEnterAgentMode
	CommonAckEnterAgentMode
	CommonReqQuitAgentMode
	CommonAckQuitAgentMode
	CommonReqLeaveDesk
	CommonAckLeaveDesk
	CommonBcKickout
	CommonReqAllowance
	CommonAckAllowance
	CommonReqApplyDissolve
	CommonBcApplyDissolve
	CommonReqApplyDissolveBack
	CommonAckApplyDissolveBack
	CommonAckTimeout
	CommonBcUserBreak
	CommonReqClickStatistic
	CommonReqOffline
	AwardReqOnline
	WardAckOnline
	AwardReqOnlineInfo
	AwardAckOnlineInfo
	AwardReqGetNewUser
	AwardAckGetNewUser
	PayBaseProduct
	PayBasePaymodel
	PayBaseDetails
	WxpayReqUnifiedorder
	WxpayAckUnifiedorder
	WxpayReqSyncChecker
	WxpayAckSyncChecker
	ApplepayReqRechargecb
	ApplepayAcksRechargecb
	CardInfo
	RoomTypeInfo
	ComposeCard
	PlayerCard
	GameReqBuxiazi
	GameAckBuxiazi
	PdkBaseRoomTypeInfo
	PdkBaseTimerInfo
	GameSession
	CommonSrvGameUser
	CommonSrvGameDesk
	Taward
	AwardMOnline
	TConfigSys
	TConfigDrawLottery
	User
	TNotice
	TGameCounts
	TUserTask
	Push
	CommonSrvMsgGameCount
	CommonSrvPokerPai
	RedisGameServerStatus
	TTask
	DdzBaseRoomTypeInfo
	DdzBasePlayerInfo
	DdzBasePlayerRateInfo
	DdzBaseCommonRateInfo
	DdzBaseTimerInfo
	DdzBaseDeskInfo
	DdzReqDissolveDesk
	DdzAckDissolveDesk
	DdzReqReady
	DdzAckReady
	DdzBaseWinCoinInfo
	DdzBaCurrentResult
	DdzBaseEndLotteryInfo
	DdzBcEndLottery
	DdzReqCreateDesk
	DdzAckCreateDesk
	DdzReqGameRecord
	DdzBaseUserRecord
	DdzBaseGameRecord
	DdzAckGameRecord
	DdzReqEnterDesk
	DdzAckEnterDesk
	DdzBcOpening
	DdzBcDealCards
	DdzReqShowHandPokers
	DdzAckShowHandPokers
	DdzReqJiaoDiZhu
	DdzAckJiaoDiZhu
	DdzReqRobDiZhu
	DdzAckRobDiZhu
	DdzReqDouble
	DdzAckDouble
	DdzBcStartPlay
	DdzReqMenuZhua
	DdzAckMenuZhua
	DdzReqSeeCards
	DdzAckSeeCards
	DdzReqPull
	DdzAckPull
	DdzReqOutCards
	DdzAckOutCards
	DdzReqActGuo
	DdzAckGuoAck
	DdzBcOverTurn
	DdzBcGameInfo
	DdzSrvOutPokerPais
	DdzSrvDeskTongJi
	DdzSrvDesk
	DdzSrvGameData
	DdzSrvBillBean
	DdzSrvBill
	DdzSrvUserStatisticsRound
	DdzSrvUserStatistics
	DdzSrvUser
	DdzSrvRoom
	DdzSrvBak
	ErddzBaseRoomTypeInfo
	ErddzBasePlayerInfo
	ErddzBasePlayerRateInfo
	ErddzBaseCommonRateInfo
	ErddzBaseTimerInfo
	ErddzBaseDeskInfo
	ErddzReqDissolveDesk
	ErddzAckDissolveDesk
	ErddzReqReady
	ErddzAckReady
	ErddzBaseWinCoinInfo
	ErddzBcCurrentResult
	ErddzBaseEndLotteryInfo
	ErddzBcEndLottery
	ErddzReqCreateDesk
	ErddzAckCreateDesk
	ErddzReqEnterDesk
	ErddzAckEnterDesk
	ErddzBcOpening
	ErddzBcDealCards
	ErddzReqJiaoDiZhu
	ErddzAckJiaoDiZhu
	ErddzReqRobDiZhu
	ErddzAckRobDiZhu
	ErddzReqRangcards
	ErddzAckRangcards
	ErddzReqDouble
	ErddzAckDouble
	ErddzBcStartPlay
	ErddzReqOutCards
	ErddzAckOutCards
	ErddzReqActGuo
	ErddzAckActGuo
	ErddzBcOverTurn
	ErddzBcGameInfo
	HallReqEvent
	HallAckEvent
	HallLotteryItem
	HallSignLotteryInfo
	HallDrawLotteryInfo
	HallReqMailList
	HallAckMailList
	HallReqTask
	HallAckTask
	HallReqCheckTask
	HallAckCheckTask
	HallReqTaskSum
	HallAckTaskSum
	HallReqCheckBonus
	HallAckCheckBonus
	HallReqBagItems
	HallAckBagItems
	HallReqUserData
	HallAckUserData
	HallReqUpdateRealData
	HallAckUpdateRealData
	HallReqGoodsList
	HallAckGoodsList
	HallReqGoodsBuy
	HallAckGoodsBuy
	HallGoodsItemMsg
	HallReqRank
	HallAckRank
	HallReqDrawLottery
	HallAckDrawLottery
	HallReqDsLotteryInfo
	HallAckDsLotteryInfo
	HallReqFriendsList
	HallAckFriendsList
	HallReqRecommendUserList
	HallAckRecommendUserList
	HallReqFriendsSearch
	HallAckFriendsSearch
	HallReqAddFriend
	HallAckAddFriend
	HallReqDelFriend
	HallAckDelFriend
	HallFriendState
	HallUserInfo
	HallAckStrongboxInfo
	HallReqStrongboxInfo
	HallReqStrongboxAccess
	HallAckStrongboxAccess
	Game_GameRecord
	BeanUserRecord
	BeanGameRecord
	Game_AckGameRecord
	HallItemEvent
	HallMailItem
	HallBagItem
	HallItemTask
	HallUserData
	HallRankItem
	CoinZone
	DiamondZone
	ExchangeZone
	BuyZone
	GoodsItem
	HallStrongboxInfo
	PlaybackReqPage
	PlaybackAckPage
	PlayerInfo
	PlaybackSnapshot
	PlaybackDeskInfo
	PlaybackItem
	PlaybackReqPageByGameid
	PlaybackPdkAckPage
	PdkPlayerInfo
	PdkPlaybackSnapshot
	PdkPlaybackDeskInfo
	VersionInfo
	AssetInfo
	HotupdateReqVersionInfo
	HotupdateAckVersionInfo
	HotupdateReqAssetsInfo
	HotupdateAckAssetsInfo
	HotupdateReqGameAssetsInfo
	HotupdateAckGameAssetsInfo
	NiuniuClientPoker
	NiuniuUserBill
	NiuniuDeskOption
	NiuniuClientDesk
	NiuniuClientUser
	NiuCreateDeskReq
	NiuEnterDeskReq
	NiuEnterDeskAck
	NiuEnterDeskBc
	NiuSwitchReadyReq
	NiuSwitchReadyAck
	NiuSwitchReadyBc
	NiuStartGameOt
	NiuQiangzhuangOt
	NiuQiangzhuangReq
	NiuQiangzhuangAck
	NiuQiangzhuangResItem
	NiuQiangzhuangResBc
	NiuJiabeiOt
	NiuJiabeiReq
	NiuJiabeiAck
	NiuJiabeiBc
	NiuBipaiResultItem
	NiuBipaiResultBc
	NiuGameEnd
	NiuDeskDissolveDoneBc
	NiuOwnerDissolveReq
	NiuOwnerDissolveAck
	NiuOfflineBc
	NiuniuSrvPoker
	NiuniuSrvDesk
	NiuniuSrvUser
	NiuniuSrvRoom
	PdkBasePlayerInfo
	PdkBasePlayerRateInfo
	PdkBaseCommonRateInfo
	PdkBaseDeskInfo
	PdkReqDissolveDesk
	PdkAckDissolveDesk
	PdkReqReady
	PdkAckReady
	PdkBaseWinCoinInfo
	PdkBaCurrentResult
	PdkBaseEndLotteryInfo
	PdkBcEndLottery
	PdkReqCreateDesk
	PdkAckCreateDesk
	PdkReqGameRecord
	PdkBaseUserRecord
	PdkBaseGameRecord
	PdkAckGameRecord
	PdkReqEnterDesk
	PdkAckEnterDesk
	PdkBcOpening
	PdkBcDealCards
	PdkReqShowHandPokers
	PdkAckShowHandPokers
	PdkReqJiaoDiZhu
	PdkAckJiaoDiZhu
	PdkReqRobDiZhu
	PdkAckRobDiZhu
	PdkReqDouble
	PdkAckDouble
	PdkBcStartPlay
	PdkReqMenuZhua
	PdkAckMenuZhua
	PdkReqSeeCards
	PdkAckSeeCards
	PdkReqPull
	PdkAckPull
	PdkReqOutCards
	PdkAckOutCards
	PdkReqActGuo
	PdkAckGuoAck
	PdkBcOverTurn
	PdkBcGameInfo
	PdkSrvOutPokerPais
	PdkSrvDeskTongJi
	PdkSrvDesk
	PdkSrvGameData
	PdkSrvBillBean
	PdkSrvBill
	PdkSrvUserStatisticsRound
	PdkSrvUserStatistics
	PdkSrvUser
	PdkSrvRoom
	PdkSrvBak
	PezBase_PaiInfo
	PezBase_PlayConf
	PezBase_RoomTypeInfo
	PezBaseTimerInfo
	PezBase_PaiValue
	PezBase_PlayerCard
	PezLastScore
	PezBase_PlayerInfo
	PezBase_DeskGameInfo
	PezTuozi
	Pez_DissolveDesk
	Pez_AckDissolveDesk
	Pez_Ready
	Pez_AckReady
	Pez_EndLotteryInfo
	Pez_SendCurrentResult
	EndLottery
	Pez_SendEndLottery
	Pez_UserBean
	Pez_Bill
	PezReq_CreateRoom
	PezReq_AckCreateRoom
	PezReq_EnterRoom
	PezReq_AckEnterRoom
	Pez_Notice
	Pez_AckNotice
	Pez_GameRecord
	PezBeanUserRecord
	PezBeanGameRecord
	Pez_AckGameRecord
	Pez_Feedback
	PezUserCoinBean
	Pez_Opening
	Pez_DealCards
	Pez_Bet
	Pez_AckBet
	Pez_BCOpenPai
	Pez_SendGameInfo
	RpcHallUpdateConfig
	ZjhReqGetRoomList
	ZjhBaseRoomInfo
	ZjhAckRoomList
	ZjhReqEnterDesk
	ZjhBaseUserInfo
	ZjhDeskStateAck
	ZjhBcGameInfo
	ZjhBaseDeskInfo
	ZjhBcNewPlayerEnter
	ZjhReqLeave
	ZjhBcLeave
	ZjhBasePlayerInfo
	ZjhBcOpening
	ZjhReqReady
	ZjhBcReady
	ZjhBcOverTurn
	ZjhReqCheckCards
	ZjhBcCheckCards
	ZjhReqFold
	ZjhBcFold
	ZjhReqCall
	ZjhBcCall
	ZjhReqBloodFight
	ZjhBcBloodFight
	ZjhBcBloodEnd
	ZjhReqRaiseFight
	ZjhBcRaiseAck
	ZjhReqCompare
	ZjhBcCompare
	ZjhBcGameEnd
	ZjhBcDeskState
	ZjhReqDaShang
	ZjhBcDaShang
	ZjhSrvPoker
	ZjhSrvBill
	ZjhSrv_GameData
	ZjhSrvUser
	ZjhSrvDesk
	ZjhSrvRoom
*/
package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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

// Ignoring public import of niuniu_client_poker from niuniu_base.proto

// Ignoring public import of niuniu_user_bill from niuniu_base.proto

// Ignoring public import of niuniu_desk_option from niuniu_base.proto

// Ignoring public import of niuniu_enum_protoid from niuniu_base.proto

// Ignoring public import of niuniu_enum_PokerType from niuniu_base.proto

// Ignoring public import of niuniu_enum_desk_state from niuniu_base.proto

// Ignoring public import of niuniu_enum_banker_rule from niuniu_base.proto

// ==============================百人牛牛==================================
type BainiuDeskStatus int32

const (
	BainiuDeskStatus_BAINIU_WAIT_YAZHU BainiuDeskStatus = 1
	BainiuDeskStatus_BAINIU_WAIT_FAPAI BainiuDeskStatus = 2
	BainiuDeskStatus_BAINIU_WAIT_BIPAI BainiuDeskStatus = 3
)

var BainiuDeskStatus_name = map[int32]string{
	1: "BAINIU_WAIT_YAZHU",
	2: "BAINIU_WAIT_FAPAI",
	3: "BAINIU_WAIT_BIPAI",
}
var BainiuDeskStatus_value = map[string]int32{
	"BAINIU_WAIT_YAZHU": 1,
	"BAINIU_WAIT_FAPAI": 2,
	"BAINIU_WAIT_BIPAI": 3,
}

func (x BainiuDeskStatus) Enum() *BainiuDeskStatus {
	p := new(BainiuDeskStatus)
	*p = x
	return p
}
func (x BainiuDeskStatus) String() string {
	return proto.EnumName(BainiuDeskStatus_name, int32(x))
}
func (x *BainiuDeskStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuDeskStatus_value, data, "BainiuDeskStatus")
	if err != nil {
		return err
	}
	*x = BainiuDeskStatus(value)
	return nil
}
func (BainiuDeskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 押注区域
type BainiuAreaName int32

const (
	BainiuAreaName_BAINIU_AREA_BANKER BainiuAreaName = 0
	BainiuAreaName_BAINIU_AREA_TIAN   BainiuAreaName = 1
	BainiuAreaName_BAINIU_AREA_DI     BainiuAreaName = 2
	BainiuAreaName_BAINIU_AREA_XUAN   BainiuAreaName = 3
	BainiuAreaName_BAINIU_AREA_HUANG  BainiuAreaName = 4
)

var BainiuAreaName_name = map[int32]string{
	0: "BAINIU_AREA_BANKER",
	1: "BAINIU_AREA_TIAN",
	2: "BAINIU_AREA_DI",
	3: "BAINIU_AREA_XUAN",
	4: "BAINIU_AREA_HUANG",
}
var BainiuAreaName_value = map[string]int32{
	"BAINIU_AREA_BANKER": 0,
	"BAINIU_AREA_TIAN":   1,
	"BAINIU_AREA_DI":     2,
	"BAINIU_AREA_XUAN":   3,
	"BAINIU_AREA_HUANG":  4,
}

func (x BainiuAreaName) Enum() *BainiuAreaName {
	p := new(BainiuAreaName)
	*p = x
	return p
}
func (x BainiuAreaName) String() string {
	return proto.EnumName(BainiuAreaName_name, int32(x))
}
func (x *BainiuAreaName) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuAreaName_value, data, "BainiuAreaName")
	if err != nil {
		return err
	}
	*x = BainiuAreaName(value)
	return nil
}
func (BainiuAreaName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// ===============================比牌结果================================
// 比牌类型
type BainiuWinType int32

const (
	BainiuWinType_BAINIU_BIPAI_PINGJU  BainiuWinType = 1
	BainiuWinType_BAINIU_BIPAI_SHENGLI BainiuWinType = 2
	BainiuWinType_BAINIU_BIPAI_SHIBAI  BainiuWinType = 3
	BainiuWinType_BAINIU_BIPAI_ZJTC    BainiuWinType = 4
	BainiuWinType_BAINIU_BIPAI_ZJTP    BainiuWinType = 5
)

var BainiuWinType_name = map[int32]string{
	1: "BAINIU_BIPAI_PINGJU",
	2: "BAINIU_BIPAI_SHENGLI",
	3: "BAINIU_BIPAI_SHIBAI",
	4: "BAINIU_BIPAI_ZJTC",
	5: "BAINIU_BIPAI_ZJTP",
}
var BainiuWinType_value = map[string]int32{
	"BAINIU_BIPAI_PINGJU":  1,
	"BAINIU_BIPAI_SHENGLI": 2,
	"BAINIU_BIPAI_SHIBAI":  3,
	"BAINIU_BIPAI_ZJTC":    4,
	"BAINIU_BIPAI_ZJTP":    5,
}

func (x BainiuWinType) Enum() *BainiuWinType {
	p := new(BainiuWinType)
	*p = x
	return p
}
func (x BainiuWinType) String() string {
	return proto.EnumName(BainiuWinType_name, int32(x))
}
func (x *BainiuWinType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BainiuWinType_value, data, "BainiuWinType")
	if err != nil {
		return err
	}
	*x = BainiuWinType(value)
	return nil
}
func (BainiuWinType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 押注区域
type BainiuAreaInfo struct {
	Name             *BainiuAreaName    `protobuf:"varint,1,opt,name=name,enum=ddproto.BainiuAreaName" json:"name,omitempty"`
	Poker            *NiuniuClientPoker `protobuf:"bytes,2,opt,name=poker" json:"poker,omitempty"`
	AllChips         *int64             `protobuf:"varint,3,opt,name=all_chips" json:"all_chips,omitempty"`
	MyChips          *int64             `protobuf:"varint,4,opt,name=my_chips" json:"my_chips,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *BainiuAreaInfo) Reset()                    { *m = BainiuAreaInfo{} }
func (m *BainiuAreaInfo) String() string            { return proto.CompactTextString(m) }
func (*BainiuAreaInfo) ProtoMessage()               {}
func (*BainiuAreaInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BainiuAreaInfo) GetName() BainiuAreaName {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuAreaInfo) GetPoker() *NiuniuClientPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *BainiuAreaInfo) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuAreaInfo) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

// 房间状态
type BainiuClientDesk struct {
	Status           *BainiuDeskStatus `protobuf:"varint,1,opt,name=status,enum=ddproto.BainiuDeskStatus" json:"status,omitempty"`
	Areas            *BainiuAreaInfo   `protobuf:"bytes,2,opt,name=areas" json:"areas,omitempty"`
	SurplusTime      *int64            `protobuf:"varint,3,opt,name=surplus_time" json:"surplus_time,omitempty"`
	Users            *BainiuClientUser `protobuf:"bytes,4,opt,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BainiuClientDesk) Reset()                    { *m = BainiuClientDesk{} }
func (m *BainiuClientDesk) String() string            { return proto.CompactTextString(m) }
func (*BainiuClientDesk) ProtoMessage()               {}
func (*BainiuClientDesk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BainiuClientDesk) GetStatus() BainiuDeskStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return BainiuDeskStatus_BAINIU_WAIT_YAZHU
}

func (m *BainiuClientDesk) GetAreas() *BainiuAreaInfo {
	if m != nil {
		return m.Areas
	}
	return nil
}

func (m *BainiuClientDesk) GetSurplusTime() int64 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

func (m *BainiuClientDesk) GetUsers() *BainiuClientUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// 用户状态
type BainiuClientUser struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	Image            *string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	NickName         *string `protobuf:"bytes,3,opt,name=nick_name" json:"nick_name,omitempty"`
	Coin             *int64  `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuClientUser) Reset()                    { *m = BainiuClientUser{} }
func (m *BainiuClientUser) String() string            { return proto.CompactTextString(m) }
func (*BainiuClientUser) ProtoMessage()               {}
func (*BainiuClientUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BainiuClientUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuClientUser) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

func (m *BainiuClientUser) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BainiuClientUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// ===============================进房===================================
// 进房请求
type BainiuEnterDeskReq struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuEnterDeskReq) Reset()                    { *m = BainiuEnterDeskReq{} }
func (m *BainiuEnterDeskReq) String() string            { return proto.CompactTextString(m) }
func (*BainiuEnterDeskReq) ProtoMessage()               {}
func (*BainiuEnterDeskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BainiuEnterDeskReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 进房回复
type BainiuEnterDeskAck struct {
	Header           *ProtoHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskState        *BainiuClientDesk `protobuf:"bytes,2,opt,name=desk_state" json:"desk_state,omitempty"`
	UserInfo         *BainiuClientUser `protobuf:"bytes,3,opt,name=user_info" json:"user_info,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *BainiuEnterDeskAck) Reset()                    { *m = BainiuEnterDeskAck{} }
func (m *BainiuEnterDeskAck) String() string            { return proto.CompactTextString(m) }
func (*BainiuEnterDeskAck) ProtoMessage()               {}
func (*BainiuEnterDeskAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BainiuEnterDeskAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BainiuEnterDeskAck) GetDeskState() *BainiuClientDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *BainiuEnterDeskAck) GetUserInfo() *BainiuClientUser {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

// ===============================押注===================================
// 押注overturn
type BainiuYazhuOt struct {
	DeskTime         *int64 `protobuf:"varint,1,opt,name=desk_time" json:"desk_time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BainiuYazhuOt) Reset()                    { *m = BainiuYazhuOt{} }
func (m *BainiuYazhuOt) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuOt) ProtoMessage()               {}
func (*BainiuYazhuOt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BainiuYazhuOt) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

// 押注请求
type BainiuYazhuReq struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,2,opt,name=area_name,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *uint32         `protobuf:"varint,3,opt,name=chips" json:"chips,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuReq) Reset()                    { *m = BainiuYazhuReq{} }
func (m *BainiuYazhuReq) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuReq) ProtoMessage()               {}
func (*BainiuYazhuReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BainiuYazhuReq) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuReq) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuReq) GetChips() uint32 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

// 押注回复
type BainiuYazhuAck struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32         `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,3,opt,name=area_name,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *int64          `protobuf:"varint,4,opt,name=chips" json:"chips,omitempty"`
	AllChips         *int64          `protobuf:"varint,5,opt,name=all_chips" json:"all_chips,omitempty"`
	SurplusCoin      *int64          `protobuf:"varint,6,opt,name=surplus_coin" json:"surplus_coin,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuAck) Reset()                    { *m = BainiuYazhuAck{} }
func (m *BainiuYazhuAck) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuAck) ProtoMessage()               {}
func (*BainiuYazhuAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BainiuYazhuAck) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BainiuYazhuAck) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuAck) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuAck) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *BainiuYazhuAck) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuYazhuAck) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

// 押注广播
type BainiuYazhuBc struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=user_id" json:"user_id,omitempty"`
	AreaName         *BainiuAreaName `protobuf:"varint,2,opt,name=area_name,enum=ddproto.BainiuAreaName" json:"area_name,omitempty"`
	Chips            *int64          `protobuf:"varint,3,opt,name=chips" json:"chips,omitempty"`
	AllChips         *int64          `protobuf:"varint,4,opt,name=all_chips" json:"all_chips,omitempty"`
	SurplusCoin      *int64          `protobuf:"varint,5,opt,name=surplus_coin" json:"surplus_coin,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BainiuYazhuBc) Reset()                    { *m = BainiuYazhuBc{} }
func (m *BainiuYazhuBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuYazhuBc) ProtoMessage()               {}
func (*BainiuYazhuBc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BainiuYazhuBc) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BainiuYazhuBc) GetAreaName() BainiuAreaName {
	if m != nil && m.AreaName != nil {
		return *m.AreaName
	}
	return BainiuAreaName_BAINIU_AREA_BANKER
}

func (m *BainiuYazhuBc) GetChips() int64 {
	if m != nil && m.Chips != nil {
		return *m.Chips
	}
	return 0
}

func (m *BainiuYazhuBc) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuYazhuBc) GetSurplusCoin() int64 {
	if m != nil && m.SurplusCoin != nil {
		return *m.SurplusCoin
	}
	return 0
}

// ===============================发牌=================================
// 发牌广播
type BainiuFapaiBc struct {
	Poker            *NiuniuClientPoker `protobuf:"bytes,1,opt,name=poker" json:"poker,omitempty"`
	DeskTime         *int64             `protobuf:"varint,2,opt,name=desk_time" json:"desk_time,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *BainiuFapaiBc) Reset()                    { *m = BainiuFapaiBc{} }
func (m *BainiuFapaiBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuFapaiBc) ProtoMessage()               {}
func (*BainiuFapaiBc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BainiuFapaiBc) GetPoker() *NiuniuClientPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *BainiuFapaiBc) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

// 赢家列表
type BainiuWinUserItem struct {
	NickName         *string `protobuf:"bytes,1,opt,name=nick_name" json:"nick_name,omitempty"`
	Bill             *int64  `protobuf:"varint,2,opt,name=bill" json:"bill,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BainiuWinUserItem) Reset()                    { *m = BainiuWinUserItem{} }
func (m *BainiuWinUserItem) String() string            { return proto.CompactTextString(m) }
func (*BainiuWinUserItem) ProtoMessage()               {}
func (*BainiuWinUserItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BainiuWinUserItem) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *BainiuWinUserItem) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

// 比牌结果广播
type BainiuBipaiBc struct {
	DeskTime         *int64               `protobuf:"varint,1,opt,name=desk_time" json:"desk_time,omitempty"`
	Coin             *int64               `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	Type             *BainiuWinType       `protobuf:"varint,3,opt,name=type,enum=ddproto.BainiuWinType" json:"type,omitempty"`
	MyChips          *int64               `protobuf:"varint,4,opt,name=my_chips" json:"my_chips,omitempty"`
	MyBill           *int64               `protobuf:"varint,5,opt,name=my_bill" json:"my_bill,omitempty"`
	AllChips         *int64               `protobuf:"varint,6,opt,name=all_chips" json:"all_chips,omitempty"`
	BankerBill       *int64               `protobuf:"varint,7,opt,name=banker_bill" json:"banker_bill,omitempty"`
	Winer            []*BainiuWinUserItem `protobuf:"bytes,8,rep,name=winer" json:"winer,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *BainiuBipaiBc) Reset()                    { *m = BainiuBipaiBc{} }
func (m *BainiuBipaiBc) String() string            { return proto.CompactTextString(m) }
func (*BainiuBipaiBc) ProtoMessage()               {}
func (*BainiuBipaiBc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BainiuBipaiBc) GetDeskTime() int64 {
	if m != nil && m.DeskTime != nil {
		return *m.DeskTime
	}
	return 0
}

func (m *BainiuBipaiBc) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *BainiuBipaiBc) GetType() BainiuWinType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return BainiuWinType_BAINIU_BIPAI_PINGJU
}

func (m *BainiuBipaiBc) GetMyChips() int64 {
	if m != nil && m.MyChips != nil {
		return *m.MyChips
	}
	return 0
}

func (m *BainiuBipaiBc) GetMyBill() int64 {
	if m != nil && m.MyBill != nil {
		return *m.MyBill
	}
	return 0
}

func (m *BainiuBipaiBc) GetAllChips() int64 {
	if m != nil && m.AllChips != nil {
		return *m.AllChips
	}
	return 0
}

func (m *BainiuBipaiBc) GetBankerBill() int64 {
	if m != nil && m.BankerBill != nil {
		return *m.BankerBill
	}
	return 0
}

func (m *BainiuBipaiBc) GetWiner() []*BainiuWinUserItem {
	if m != nil {
		return m.Winer
	}
	return nil
}

func init() {
	proto.RegisterType((*BainiuAreaInfo)(nil), "ddproto.bainiu_area_info")
	proto.RegisterType((*BainiuClientDesk)(nil), "ddproto.bainiu_client_desk")
	proto.RegisterType((*BainiuClientUser)(nil), "ddproto.bainiu_client_user")
	proto.RegisterType((*BainiuEnterDeskReq)(nil), "ddproto.bainiu_enter_desk_req")
	proto.RegisterType((*BainiuEnterDeskAck)(nil), "ddproto.bainiu_enter_desk_ack")
	proto.RegisterType((*BainiuYazhuOt)(nil), "ddproto.bainiu_yazhu_ot")
	proto.RegisterType((*BainiuYazhuReq)(nil), "ddproto.bainiu_yazhu_req")
	proto.RegisterType((*BainiuYazhuAck)(nil), "ddproto.bainiu_yazhu_ack")
	proto.RegisterType((*BainiuYazhuBc)(nil), "ddproto.bainiu_yazhu_bc")
	proto.RegisterType((*BainiuFapaiBc)(nil), "ddproto.bainiu_fapai_bc")
	proto.RegisterType((*BainiuWinUserItem)(nil), "ddproto.bainiu_win_user_item")
	proto.RegisterType((*BainiuBipaiBc)(nil), "ddproto.bainiu_bipai_bc")
	proto.RegisterEnum("ddproto.BainiuDeskStatus", BainiuDeskStatus_name, BainiuDeskStatus_value)
	proto.RegisterEnum("ddproto.BainiuAreaName", BainiuAreaName_name, BainiuAreaName_value)
	proto.RegisterEnum("ddproto.BainiuWinType", BainiuWinType_name, BainiuWinType_value)
}

var fileDescriptor0 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x13, 0xbb, 0x6d, 0x6e, 0xff, 0x9c, 0x69, 0xfa, 0x7d, 0xa1, 0x80, 0x54, 0x45, 0x15,
	0x54, 0xa1, 0x0a, 0x52, 0x37, 0xac, 0x1d, 0x28, 0x8d, 0x0b, 0x8a, 0x42, 0xdb, 0x08, 0xda, 0xcd,
	0x30, 0x71, 0x5c, 0x3a, 0x8a, 0x63, 0x9b, 0xd8, 0x11, 0x0a, 0x3b, 0x76, 0x48, 0x88, 0x57, 0xe0,
	0x0d, 0x10, 0x0f, 0xc3, 0x0b, 0x31, 0x73, 0x3d, 0x75, 0x5c, 0x27, 0x40, 0x17, 0x6c, 0xda, 0xe8,
	0xcc, 0x99, 0x7b, 0xcf, 0x39, 0xf7, 0x8e, 0xa1, 0xdc, 0x63, 0xdc, 0xe7, 0x63, 0x1a, 0x7a, 0x6c,
	0xd2, 0x08, 0x47, 0x41, 0x1c, 0x90, 0xa5, 0x7e, 0x1f, 0x7f, 0x6c, 0x6f, 0x3a, 0xc1, 0x70, 0x18,
	0xf8, 0xd4, 0xf1, 0xb8, 0xeb, 0xc7, 0xc9, 0xe9, 0x76, 0x59, 0xb0, 0xe5, 0x85, 0x1e, 0x8b, 0xdc,
	0x04, 0xaa, 0x7d, 0xd5, 0xc0, 0x54, 0x65, 0xd8, 0xc8, 0x65, 0x94, 0xfb, 0x97, 0x01, 0x79, 0x08,
	0xba, 0xcf, 0x86, 0x6e, 0x55, 0xdb, 0xd1, 0xf6, 0xd6, 0x0f, 0xee, 0x34, 0x54, 0xd1, 0x46, 0x96,
	0x28, 0x09, 0xe4, 0x11, 0x18, 0x61, 0x30, 0x70, 0x47, 0xd5, 0x82, 0x60, 0xae, 0x1c, 0xdc, 0x4b,
	0x99, 0xaa, 0x51, 0xd2, 0x9d, 0x22, 0x87, 0x94, 0xa1, 0xc4, 0x3c, 0x8f, 0x3a, 0x57, 0x3c, 0x8c,
	0xaa, 0x45, 0x71, 0xa1, 0x48, 0x4c, 0x58, 0x1e, 0x4e, 0x14, 0xa2, 0x4b, 0xa4, 0xf6, 0x5d, 0x03,
	0xa2, 0xda, 0xa8, 0xcb, 0x7d, 0x37, 0x1a, 0x88, 0x46, 0x8b, 0x51, 0xcc, 0xe2, 0x71, 0xa4, 0x34,
	0xdd, 0xcd, 0x6b, 0x92, 0x2c, 0x9a, 0x50, 0xc8, 0x1e, 0x18, 0x52, 0x62, 0xa4, 0x54, 0xcd, 0xd7,
	0x8f, 0x46, 0x2b, 0xb0, 0x1a, 0x8d, 0x47, 0xa1, 0x37, 0x8e, 0x68, 0xcc, 0x85, 0xe1, 0x44, 0x55,
	0x1d, 0x8c, 0x71, 0xe4, 0x8e, 0x12, 0x49, 0x2b, 0xb3, 0xbd, 0x94, 0x30, 0xc9, 0xa9, 0x75, 0xf3,
	0x72, 0x25, 0x4a, 0x36, 0x60, 0x49, 0xfe, 0xa7, 0xbc, 0x8f, 0x7a, 0xd7, 0xc8, 0x1a, 0x18, 0x7c,
	0xc8, 0xde, 0xb9, 0x28, 0xa9, 0x24, 0xa3, 0xf0, 0xb9, 0x33, 0xc0, 0x10, 0xb1, 0x69, 0x89, 0xac,
	0x82, 0xee, 0x04, 0xdc, 0x57, 0x31, 0xec, 0xc1, 0x96, 0x2a, 0x2b, 0x6a, 0x8a, 0x42, 0x68, 0x6f,
	0xe4, 0xbe, 0x9f, 0xa9, 0x5c, 0xfb, 0xa6, 0xcd, 0xa3, 0x32, 0x67, 0x40, 0x76, 0x61, 0xf1, 0xca,
	0x65, 0x7d, 0x31, 0x1d, 0x0d, 0x7d, 0x54, 0x52, 0x1f, 0x1d, 0xf9, 0xb7, 0x85, 0x67, 0xe4, 0x31,
	0x40, 0x9a, 0x9d, 0xab, 0x12, 0xfb, 0x9d, 0x63, 0x1c, 0x45, 0x03, 0x4a, 0x89, 0x02, 0x11, 0x20,
	0x6a, 0xff, 0x4b, 0x42, 0xbb, 0xb0, 0xa1, 0xd0, 0x09, 0xfb, 0x78, 0x35, 0xa6, 0x41, 0x2c, 0xed,
	0x63, 0x4f, 0xcc, 0x5c, 0x43, 0xc3, 0x6f, 0xd3, 0x35, 0x4c, 0x58, 0xf3, 0xbc, 0x92, 0x7d, 0xb1,
	0x41, 0xd7, 0xbb, 0x87, 0x52, 0xff, 0xb8, 0x9c, 0x22, 0xf3, 0xe9, 0xae, 0xad, 0xd5, 0x7e, 0x68,
	0xb9, 0x16, 0xb7, 0xcf, 0x28, 0x23, 0xa4, 0x30, 0x2b, 0xa4, 0x78, 0x6b, 0x21, 0x38, 0xdb, 0x9b,
	0xef, 0xc0, 0x40, 0x28, 0xb3, 0x87, 0xb8, 0x04, 0x8b, 0x98, 0xc9, 0x67, 0x2d, 0x17, 0x5d, 0xcf,
	0xf9, 0xa7, 0x99, 0xe4, 0xa4, 0xe8, 0x73, 0xa5, 0xa0, 0xc0, 0xda, 0xab, 0x54, 0xc9, 0x25, 0x0b,
	0x19, 0x97, 0x4a, 0xd2, 0xb7, 0xaf, 0xdd, 0xee, 0xed, 0x4f, 0x27, 0x5e, 0xc0, 0x92, 0x4f, 0xa0,
	0xa2, 0x4a, 0x7e, 0xe0, 0x3e, 0x4d, 0x4c, 0xc5, 0xee, 0xf0, 0xe6, 0xdb, 0xd0, 0xae, 0xdf, 0x46,
	0x8f, 0x7b, 0x9e, 0xba, 0xf8, 0x73, 0x1a, 0x4b, 0x8f, 0x2b, 0x31, 0xb3, 0x1b, 0x95, 0x3e, 0x28,
	0xbc, 0x44, 0x1e, 0x80, 0x1e, 0x4f, 0xc2, 0xeb, 0x61, 0x55, 0xf3, 0x09, 0x49, 0x09, 0xf2, 0x7c,
	0xf6, 0x8b, 0x24, 0x13, 0x17, 0x08, 0xf6, 0x37, 0x66, 0x43, 0xc3, 0x49, 0x91, 0x4d, 0x58, 0xe9,
	0x31, 0x5f, 0x18, 0x4d, 0x78, 0x4b, 0x08, 0xee, 0x83, 0x21, 0xca, 0x8a, 0x80, 0x96, 0x77, 0x8a,
	0x22, 0xa0, 0xfb, 0xf3, 0x7a, 0xa6, 0xb6, 0xeb, 0xe7, 0xe9, 0x87, 0x24, 0xfb, 0x29, 0xdb, 0x82,
	0x72, 0xd3, 0xb2, 0xdb, 0x76, 0x97, 0xbe, 0xb6, 0xec, 0x33, 0x7a, 0x6e, 0x5d, 0xb4, 0xba, 0xa6,
	0x96, 0x87, 0x9f, 0x5b, 0x1d, 0xcb, 0x36, 0x0b, 0x79, 0xb8, 0x69, 0x4b, 0xb8, 0x58, 0xff, 0x94,
	0xfb, 0xc6, 0xe3, 0x26, 0xfc, 0x07, 0x44, 0x71, 0xad, 0x93, 0x43, 0x8b, 0x36, 0xad, 0xf6, 0x8b,
	0xc3, 0x13, 0x73, 0x41, 0xcc, 0xdf, 0xcc, 0xe2, 0x67, 0xb6, 0xd5, 0x16, 0x0d, 0x09, 0xac, 0x67,
	0xd1, 0x67, 0xb2, 0x5b, 0x8e, 0xf9, 0xa6, 0x2b, 0x98, 0xc5, 0x8c, 0x06, 0x44, 0x5b, 0x02, 0x3d,
	0x32, 0xf5, 0xfa, 0x97, 0xe9, 0xd0, 0xd2, 0xac, 0xff, 0x87, 0x4d, 0x45, 0x45, 0xa5, 0xb4, 0x63,
	0xb7, 0x8f, 0x8e, 0xa5, 0xbd, 0x2a, 0x54, 0x6e, 0x1c, 0x9c, 0xb6, 0x0e, 0xdb, 0x47, 0x2f, 0x65,
	0xcf, 0xfc, 0x95, 0xd3, 0x96, 0xdd, 0x94, 0x1e, 0x33, 0x6d, 0x93, 0x83, 0x8b, 0xe3, 0xb3, 0xa7,
	0xa6, 0x3e, 0x0f, 0xee, 0x98, 0x46, 0x67, 0xa1, 0xa3, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x1a,
	0xc1, 0xfd, 0x82, 0x3e, 0x07, 0x00, 0x00,
}