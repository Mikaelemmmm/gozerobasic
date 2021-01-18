package logic

import (
	"context"
	"gozerobasic/app/services/message/enum"
	"gozerobasic/app/services/message/internal/model"
	"gozerobasic/lib/tools"
	"gozerobasic/lib/xerr"
	"gozerobasic/lib/xmsg"
	"fmt"
	"github.com/tal-tech/go-zero/core/stringx"

	"gozerobasic/app/services/message/internal/svc"
	"gozerobasic/app/services/message/message"

	"github.com/tal-tech/go-zero/core/logx"
)

type SendCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCaptchaLogic {
	return &SendCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCaptchaLogic) SendCaptcha(in *message.SendCaptchaReq) (*message.SendCaptchaResp, error) {

	var err error
	tplCode, tplSign, err := enum.MessageKindEnum(in.Kind).GetMessageTplConfig()
	if err != nil {
		return nil, err
	}

	captchaKindKey := fmt.Sprintf(model.CaptchaCacheKey, in.Kind, in.Mobile)
	//redis有就用redis中的，否则重新生成
	code, err := l.svcCtx.RedisCli.Get(captchaKindKey)
	if err != nil || !stringx.NotEmpty(code) {
		code = tools.Krand(6, tools.KC_RAND_KIND_NUM)
	}

	content := make(map[string]string)
	content["code"] = code

	//存入redis
	err = l.svcCtx.RedisCli.Setex(captchaKindKey, content["code"], 300)
	if err != nil {
		logx.Errorf("SendSingleMsg err : %v , captchaKindKey : %s ,content : %v", err, captchaKindKey, content)
		return nil, xerr.NewErrMsg("发送验证码失败")
	}

	msgParam := &xmsg.SendSingleMsgParams{
		PhoneNumbers: in.Mobile,
		Content:      content,
		SignName:     tplSign,
		TemplateCode: tplCode,
	}
	err = l.svcCtx.MsgLib.SendSingleMsg(msgParam)
	if err != nil {
		return nil, err
	}

	return &message.SendCaptchaResp{}, nil
}
