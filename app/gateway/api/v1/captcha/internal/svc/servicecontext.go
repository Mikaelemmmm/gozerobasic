package svc

import (
	"gozerobasic/app/gateway/api/v1/captcha/internal/config"
	"gozerobasic/app/services/message/messageclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	MessageRpc messageclient.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MessageRpc:messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
