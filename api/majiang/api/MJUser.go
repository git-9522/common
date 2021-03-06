package api

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"casino_common/common/userService"
	ddMJProto "casino_common/proto/ddproto/mjproto"
	"casino_common/utils/agentUtils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"reflect"
	"sync/atomic"
)

type MJUser interface {
	//所有基本信息的优化
	GetUserId() uint32
	GetCoin() int64
	GetSex() int32
	GetHeadUrl() string
	GetOpenId() string
	GetNickName() string
	GetUserSkeleton() interface{}
	GetGameData() interface{}
	GetIsBreak() bool //是否离开
	GetIsLeave() bool //是否掉线
	GetIsReady() bool //是否已准备
	GetDesk() MJDesk  //得到桌子

	//基本功能
	SendOverTurn(p proto.Message) error         //发送overTurn
	SendJiaoInfos() error                       //发送下叫jiaoInfos的提示
	SendTingInfos() (interface{}, error)        //发送下叫tingInfos的提示
	WriteMsg2(p proto.Message) error            //发送信息
	DoReady() error                             //准备
	DoOut(interface{}) error                    //玩家出牌
	DoPeng(...interface{}) (interface{}, error) //碰牌
	DoChi(...interface{}) (interface{}, error)  //吃牌
	DoBu(...interface{}) (interface{}, error)   //补牌
	DoHu(interface{}) (interface{}, error)      //胡牌
	DoGang(...interface{}) (interface{}, error) //杠牌
	DoFly(...interface{}) (interface{}, error)  //飞牌
	DoTi(...interface{}) (interface{}, error)   //提牌

	//业务
	AddBillBean(interface{}) error //增加账单
}

//核心User
type MJUserCore struct {
	UserId     uint32
	URedis     userService.U_REDIS //这里可以使用redis的方法
	TopUser    MJUser              //
	Desk       MJDesk              //关联的desk
	Coin       int64               //金币
	gate.Agent                     //agent
	GameStatus                     //玩家的状态
}

//User的状态信息
type GameStatus struct {
	IsBreak       bool  //掉线
	IsLeave       bool  //离开
	IsReady       bool  //是否已准备
	IsBanker      bool  //是否是庄家
	IsOwner       bool  //是否是房主
	S             int32 //玩家的状态
	CanOut        bool  //是否可以打牌
	ApplyDissolve int32 //申请解散的状态
	IsRobot       bool  //是否是机器人
	IsAgentMode   bool  //是否是托管模式
}

const (
	MJUER_APPLYDISSOLVE_S_REFUSE  int32 = -1 //拒绝解散
	MJUER_APPLYDISSOLVE_S_DEFAULT int32 = 0  //没有处理
	MJUER_APPLYDISSOLVE_S_AGREE   int32 = 1  //同意解散
)

//New一个CoreUser
func NewMJUserCore(userId uint32, a gate.Agent) *MJUserCore {
	ret := &MJUserCore{
		UserId: userId,
		URedis: userService.U_REDIS(userId),
		Agent:  a,
	}
	ret.IsRobot = a == nil
	return ret
}

func (u *MJUserCore) DoBu(...interface{}) (interface{}, error) {
	return nil, nil
}

func (u *MJUserCore) DoChi(...interface{}) (interface{}, error) {
	return nil, nil
}

func (u *MJUserCore) DoFly(...interface{}) (interface{}, error) {
	return nil, nil
}

func (u *MJUserCore) DoTi(...interface{}) (interface{}, error) {
	return nil, nil
}

func (u *MJUserCore) GetUserId() uint32 {
	return u.UserId
}

func (u *MJUserCore) GetSex() int32 {
	return u.URedis.GetSex()
}

func (u *MJUserCore) GetHeadUrl() string {
	return u.URedis.GetHeadUrl()
}

func (u *MJUserCore) GetOpenId() string {
	return u.URedis.GetOpenId()
}

func (u *MJUserCore) GetNickName() string {
	return u.URedis.GetNickName()
}

func (u *MJUserCore) GetIp() string {
	return agentUtils.GetIP(u.Agent)
}

func (u *MJUserCore) GetCoin() int64 {
	return u.Coin
}

func (u *MJUserCore) AddCoin(c int64) int64 {
	return atomic.AddInt64(&u.Coin, c)
}

func (u *MJUserCore) GetIsBreak() bool {
	return u.GameStatus.IsBreak
}

func (u *MJUserCore) GetIsLeave() bool {
	return u.GameStatus.IsLeave
}

func (u *MJUserCore) GetIsReady() bool {
	return u.GameStatus.IsReady
}

func (u *MJUserCore) SendOverTurn(p proto.Message) error {
	return nil
}

func (u *MJUserCore) WriteMsg2(p proto.Message) error {
	if p == nil {
		log.W("%v玩家[%v]WriteMsg2() 协议为空 不发送消息", u.GetDesk().DlogDes(), u.GetUserId())
		return nil
	}
	if u.Agent == nil {
		log.W("%v玩家[%v]WriteMsg2(%v) Agent连接为空 不发送消息", u.GetDesk().DlogDes(), u.GetUserId(), reflect.TypeOf(p).String())
		return nil
	}

	//todo判断条件
	//if u.GameStatus.IsBreak {
	//	log.T("%v玩家[%v]掉线, 协议不发送 type[%v] msg[%v]", u.Desk.DlogDes(), u.GetUserId(), reflect.TypeOf(p).String(), p)
	//	return nil
	//}
	//
	log.T("%v开始给玩家[%v]发送type[%v]，msg[%v]", u.GetDesk().DlogDes(), u.GetUserId(), reflect.TypeOf(p).String(), p)
	u.Agent.WriteMsg(p)
	return nil
}

func (u *MJUserCore) GetDesk() MJDesk {
	return u.Desk
}

func (u *MJUserCore) GetGameData() interface{} {
	return nil
}

func (u *MJUserCore) SendJiaoInfos() error {
	defer Error.ErrorRecovery(fmt.Sprintf("%v给玩家[%v]发送jiaoInfos提示时异常, 已捕获待处理", u.GetDesk().DlogDes(), u.GetUserId()))
	ack := &ddMJProto.GameAckJiaoinfos{}
	ack.Header = &ddMJProto.ProtoHeader{
		UserId: proto.Uint32(u.GetUserId()),
	}
	//判断碰牌之后的叫info
	jiaoInfos, err := u.GetDesk().GetParser().GetJiaoInfos(u.GetDesk().GetUserById(u.GetUserId()).GetGameData(), u.GetDesk().GetAllMingPai(u.GetUserId()))
	if err != nil {
		log.E("%v 获取玩家[%v]jiaoinfo 时出错:err %v ", u.GetDesk().DlogDes(), u.GetUserId(), err)
		return err
	}

	log.T("%v 获取到玩家[%v]jiaoinfo[%+v]", u.GetDesk().DlogDes(), u.GetUserId(), jiaoInfos)
	if jiaoInfos == nil {
		log.T("%v 玩家[%v]jiaoinfo 为空 不发送jiaoInfo", u.GetDesk().DlogDes(), u.GetUserId())
		return nil
	}

	jfs := jiaoInfos.([]*JiaoInfo)
	if jfs == nil || len(jfs) <= 0 {
		log.T("%v 玩家[%v]jiaoinfo 为空 不发送jiaoInfo", u.GetDesk().DlogDes(), u.GetUserId())
		return nil
	}

	if jiaoInfos != nil {
		//得到叫牌的信息
		for _, jf := range jfs {
			j := &ddMJProto.JiaoInfo{}
			j.OutCard = jf.OutPai.GetCardInfo()
			for _, jfb := range jf.Jiaos {
				j.PaiInfos = append(j.PaiInfos, &ddMJProto.JiaoPaiInfo{
					HuCard: jfb.HuPai.GetCardInfo(),
					Fan:    proto.Int32(jfb.Fan),
					Count:  proto.Int32(jfb.Count),
				})
			}
			ack.JiaoInfos = append(ack.JiaoInfos, j)
		}
	}

	u.WriteMsg2(ack)
	return nil
}

func (u *MJUserCore) SendTingInfos() (interface{}, error) {
	defer Error.ErrorRecovery(fmt.Sprintf("%v给玩家[%v]发送tingInfos提示时异常, 已捕获待处理", u.GetDesk().DlogDes(), u.GetUserId()))
	ack := &ddMJProto.GameAckTinginfos{}
	ack.Header = &ddMJProto.ProtoHeader{
		UserId: proto.Uint32(u.GetUserId()),
	}
	//获取tinginfos
	tingInfoBeans, err := u.GetDesk().GetParser().GetTingInfos(u.GetDesk().GetUserById(u.GetUserId()).GetGameData(), u.GetDesk().GetAllMingPai(u.GetUserId()))
	if err != nil {
		log.E("%v 获取玩家[%v]tinginfo 时出错:err %v 发送空的tingInfo", u.GetDesk().DlogDes(), u.GetUserId(), err)
		u.WriteMsg2(ack)
		return nil, err
	}

	log.T("%v 获取到玩家[%v]tinginfo[%+v]", u.GetDesk().DlogDes(), u.GetUserId(), tingInfoBeans)
	if tingInfoBeans == nil {
		log.W("%v 玩家[%v]tinginfo 为空 发送空的tingInfo", u.GetDesk().DlogDes(), u.GetUserId())
		u.WriteMsg2(ack)
		return nil, nil
	}

	tfs := tingInfoBeans.([]*JiaoInfoBean)
	if tfs == nil || len(tfs) <= 0 {
		log.T("%v 玩家[%v]tinginfo 为空 发送空的tingInfo", u.GetDesk().DlogDes(), u.GetUserId())
		u.WriteMsg2(ack)
		return nil, nil
	}

	//得到叫牌的信息
	for _, tf := range tfs {
		j := &ddMJProto.JiaoPaiInfo{}
		j.HuCard = tf.HuPai.GetCardInfo()
		j.Fan = proto.Int32(tf.Fan)
		j.Count = proto.Int32(tf.Count)
		ack.PaiInfos = append(ack.PaiInfos, j)
	}

	u.WriteMsg2(ack)
	return tingInfoBeans, nil
}
