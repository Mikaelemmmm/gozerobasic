package logic

import (
	"context"
	"gozerobasic/app/gateway/api/v1/captcha/internal/svc"
	"gozerobasic/app/gateway/api/v1/captcha/internal/types"
	"gozerobasic/app/services/message/enum"
	"gozerobasic/app/services/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.CaptchaReq) (*types.CaptchaResp, error) {

	_,err := l.svcCtx.MessageRpc.SendCaptcha(l.ctx,&message.SendCaptchaReq{
		Kind: int64(enum.MessageKindLogin),
		Mobile: req.Mobile,
	})
	if err != nil{
		return nil, err
	}

	return &types.CaptchaResp{
		Ok: true,
	}, nil
}
