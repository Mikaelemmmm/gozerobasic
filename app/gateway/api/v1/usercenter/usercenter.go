package main

import (
	"flag"
	"github.com/tal-tech/go-zero/core/logx"

	"gozerobasic/app/gateway/api/v1/usercenter/internal/config"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/handler"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)


var configFile = flag.String("f", "etc/usercenter-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
