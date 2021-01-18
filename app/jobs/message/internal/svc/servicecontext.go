package svc

import (
	"gozerobasic/app/jobs/message/internal/config"
	"github.com/tal-tech/go-queue/dq"
)

type ServiceContext struct {
	Config config.Config
	Consumer      dq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Consumer: dq.NewConsumer(c.DqConf),
	}
}
