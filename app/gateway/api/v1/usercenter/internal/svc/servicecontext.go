package svc

import (
	"gozerobasic/app/gateway/api/v1/usercenter/internal/config"
	"gozerobasic/app/services/message/messageclient"
	"gozerobasic/app/services/usercenter/usercenterclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	UsercenterRpc usercenterclient.Usercenter
	MessageRpc messageclient.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UsercenterRpc:usercenterclient.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpc)),
		MessageRpc:messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
