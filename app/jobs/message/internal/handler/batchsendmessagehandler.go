package handler

import (
	"context"
	"gozerobasic/app/jobs/message/internal/logic"
	"gozerobasic/app/jobs/message/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

func batchSendMessageHandler(ctx *svc.ServiceContext){

	rootCxt:= context.Background()
	l := logic.NewBatchSendMessageLogic(context.Background(), ctx)
	err := l.BatchSendMessage()
	if err != nil{
		logx.WithContext(rootCxt).Error("【JOB-ERR】 : %+v ",err)
	}
}
