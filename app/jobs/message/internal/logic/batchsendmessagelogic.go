package logic

import (
	"context"
	"gozerobasic/app/jobs/message/internal/svc"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
)

type BatchSendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) BatchSendMessageLogic {
	return BatchSendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *BatchSendMessageLogic) BatchSendMessage() error {

	fmt.Printf("job BatchSendMessage start \n")

	l.svcCtx.Consumer.Consume(func(body []byte) {
		fmt.Printf("job BatchSendMessage %s \n" + string(body))
	})

	fmt.Printf("job BatchSendMessage finish \n")
	return nil
}
