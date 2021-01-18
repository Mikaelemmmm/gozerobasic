package handler

import (
	"gozerobasic/lib/xgo"
	"gozerobasic/app/jobs/message/internal/svc"
)


/**
* @Description 启动job
* @Author Mikael
* @Date 2021/1/18 12:05
* @Version 1.0
**/

func JobRun(serverCtx *svc.ServiceContext)  {

	xgo.Go(func() {
		//批量发送短信
		batchSendMessageHandler(serverCtx)
		//...更多短信业务相关的job
	})
}