package api

import (
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
)

type MJUser interface {
	WriteMsg(p proto.Message) error //发送信息
}

//核心User
type MJUserCore struct {
	UserId uint32
	uRedis userService.U_REDIS //这里可以使用redis的方法
}
