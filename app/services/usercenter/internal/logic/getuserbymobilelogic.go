package logic

import (
	"context"
	"gozerobasic/app/services/usercenter/internal/model"
	"github.com/jinzhu/copier"

	"gozerobasic/app/services/usercenter/internal/svc"
	"gozerobasic/app/services/usercenter/usercenter"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *usercenter.GetUserByMobileReq) (*usercenter.GetUserByMobileResp, error) {

	user,err:= l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	userResp:= new(usercenter.User)
	if user != nil{
		_ = copier.Copy(userResp,user)
	}

	return &usercenter.GetUserByMobileResp{
		User:userResp,
	}, nil

}
