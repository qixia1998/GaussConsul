package svc

import (
	"GaussConsul/service/gauss/rpc/internal/config"
	"GaussConsul/service/user/model"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	GaussDBConn *sql.DB
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("opengauss", c.GaussDB.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
