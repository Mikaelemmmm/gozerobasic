package logic

import (
	"context"
	"gozerobasic/app/gateway/api/v1/captcha/internal/svc"
	"gozerobasic/app/gateway/api/v1/captcha/internal/types"
	"gozerobasic/app/services/message/enum"
	"gozerobasic/app/services/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type ForgetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ForgetPwdLogic {
	return ForgetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPwdLogic) ForgetPwd(req types.CaptchaReq) (*types.CaptchaResp, error) {

	_,err := l.svcCtx.MessageRpc.SendCaptcha(l.ctx,&message.SendCaptchaReq{
		Kind: int64(enum.MessageKindForgetPwd),
		Mobile: req.Mobile,
	})
	if err != nil{
		return nil, err
	}


	return &types.CaptchaResp{
		Ok: true,
	}, nil
}
