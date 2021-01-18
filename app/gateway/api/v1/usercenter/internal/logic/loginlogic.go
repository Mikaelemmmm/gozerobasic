package logic

import (
	"context"
	msgEnum "gozerobasic/app/services/message/enum"
	"gozerobasic/app/services/usercenter/enum"
	"gozerobasic/app/services/message/messageclient"
	"gozerobasic/app/services/usercenter/usercenterclient"
	"gozerobasic/lib/xerr"
	"github.com/dgrijalva/jwt-go"
	"time"

	"gozerobasic/app/gateway/api/v1/usercenter/internal/svc"
	"gozerobasic/app/gateway/api/v1/usercenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)
var errAuthFailed = xerr.NewErrMsg("授权失败")
var errUserNotFound= xerr.NewErrMsg("用户不存在")

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

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {

	var userId int64
	var err error
	//1、校验
	switch req.AuthType {
	case int64(enum.AuthTypeMobile):
		userId,err =  l.loginByMobile(req)
	default:
		return nil,errAuthFailed
	}
	if err != nil{
		return nil, err
	}

	//2、生成token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, userId)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		JwtToken:types.JwtToken{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

/**
	生成jwt token
 */
func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func  (l *LoginLogic) loginByMobile(req types.LoginReq) (userId int64,err error) {

	//1、校验验证码
	if _,err := l.svcCtx.MessageRpc.ValidateCaptcha(l.ctx,&messageclient.ValidateCaptchaReq{
		Kind: int64(msgEnum.MessageKindLogin),
		Mobile: req.AuthKey,
		Captcha: req.Captcha,
	});err != nil{
		return 0, err
	}

	//2、查询用户
	resp,err := l.svcCtx.UsercenterRpc.GetUserByMobile(l.ctx,&usercenterclient.GetUserByMobileReq{
		Mobile: req.AuthKey,
	})
	if err != nil{
		return 0,err
	}
	if resp.User == nil{
		return 0,errUserNotFound
	}

	return resp.User.Id,nil
}
