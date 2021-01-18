package logic

import (
	"context"
	"gozerobasic/app/services/usercenter/enum"
	"gozerobasic/app/services/usercenter/internal/model"
	"gozerobasic/app/services/usercenter/internal/svc"
	"gozerobasic/app/services/usercenter/usercenter"
	"gozerobasic/lib/tools"
	"gozerobasic/lib/uniqueid"
	"gozerobasic/lib/xerr"
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"strconv"
)
var errMobileExists = xerr.NewErrMsg("手机号已存在")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {

	user,err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if user != nil{
		return nil,errMobileExists
	}

	//注册
	var insertUser model.User
	if err = l.svcCtx.UserModel.Trans(func(session sqlx.Session) error {
		var err error
		userId := uniqueid.GenId()
		insertUser = model.User{
			Id: userId,
			Mobile:in.Mobile,
			Nickname:strconv.FormatInt(userId,10),
			InviteCode:tools.Krand(tools.KC_RAND_KIND_ALL,8),
		}
		err = l.svcCtx.UserModel.TranInsert(session,insertUser)
		if err != nil{
			return err
		}

		userAuth:= model.UserAuth{
			Id: uniqueid.GenId(),
			UserId: userId,
			AuthType: int64(enum.AuthTypeMobile),
			AuthKey: in.Mobile,
		}
		err =  l.svcCtx.UserAuthModel.TranInsert(session,userAuth)
		if err != nil{
			return err
		}

		return 	nil
	});err != nil{
		return nil, err
	}

	var userResp *usercenter.User
	_ = copier.Copy(&userResp,&insertUser)

	return &usercenter.RegisterResp{
		User:userResp,
	}, nil
}
