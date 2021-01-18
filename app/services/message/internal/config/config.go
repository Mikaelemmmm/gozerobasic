package config

import (
	"gozerobasic/lib/xmsg"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Redis   redis.RedisConf
	MsgLib  xmsg.MsgConf
}
