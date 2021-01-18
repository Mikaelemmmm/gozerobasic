package svc

import (
	"gozerobasic/app/services/usercenter/internal/config"
	"gozerobasic/app/services/usercenter/internal/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c             config.Config
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:             c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UserAuthModel: model.NewUserAuthModel(sqlx.NewMysql(c.DataSource),c.Cache),
	}
}
