package svc

import (
	"gozerobasic/app/services/message/internal/config"
	"gozerobasic/lib/xmsg"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type ServiceContext struct {
	c        config.Config
	MsgLib   xmsg.IMsg
	RedisCli *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:        c,
		RedisCli: redis.NewRedis(c.Redis.Host,c.Redis.Type),
		MsgLib:   xmsg.NewMsgInstance(c.MsgLib),
	}
}
