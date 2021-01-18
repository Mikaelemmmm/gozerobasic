package logic

import (
	"context"
	"gozerobasic/app/services/message/enum"
	"gozerobasic/app/services/message/messageclient"
	"gozerobasic/app/services/usercenter/usercenterclient"
	"time"

	"gozerobasic/app/gateway/api/v1/usercenter/internal/svc"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	xerrors "github.com/pkg/errors"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterResp, error) {

	//1、校验验证码
	 if _,err := l.svcCtx.MessageRpc.ValidateCaptcha(l.ctx,&messageclient.ValidateCaptchaReq{
		Kind: int64(enum.MessageKindRegister),
		Mobile: req.Mobile,
		Captcha: req.Captcha,
	});err != nil{
		 ///Users/seven/Developer/goenv/pkg/mod/github.com/tal-tech/go-zero@v1.1.2/core/logx/tracelogger.go
		 ///Users/seven/Developer/goenv/pkg/mod/github.com/tal-tech/go-zero@v1.1.2/core/logx/logs.go

	 	return nil, xerrors.WithStack(err)
	 }


	//2、注册
	if _,err:= l.svcCtx.UsercenterRpc.Register(l.ctx,&usercenterclient.RegisterReq{
		Mobile: req.Mobile,
	});err != nil{
		return nil, err
	}

	//3、生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, 1)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		JwtToken:types.JwtToken{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}
