package interceptor

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
)

/**
* @Description rpc service logger interceptor
* @Author Mikael
* @Date 2021/1/9 13:35
* @Version 1.0
**/

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp,err = handler(ctx, req)
	if err != nil{
		logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v",err)
	}
	return resp,err
}