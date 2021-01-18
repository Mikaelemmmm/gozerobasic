package logic

import (
	"context"
	"errors"
	"gozerobasic/app/services/message/internal/model"
	"fmt"
	xerrors "github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/stringx"

	"gozerobasic/app/services/message/internal/svc"
	"gozerobasic/app/services/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type ValidateCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateCaptchaLogic {
	return &ValidateCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateCaptchaLogic) ValidateCaptcha(in *message.ValidateCaptchaReq) (*message.ValidateCaptchaResp, error) {

	captchaKindKey := fmt.Sprintf(model.CaptchaCacheKey,in.Kind,in.Mobile)
	srvCaptcha, err := l.svcCtx.RedisCli.Get(captchaKindKey)
	if err != nil || stringx.HasEmpty(srvCaptcha) {
		logx.Errorf("验证码失效 captchaKindKey:%s", captchaKindKey)
		return nil, xerrors.WithStack(errors.New("验证码失效，请重新获取"))
	}
	if srvCaptcha != in.Captcha {
		return nil, xerrors.WithStack(errors.New("验证码不正确，请重新输入"))
	}

	return &message.ValidateCaptchaResp{
		Ok: true,
	}, nil
}
