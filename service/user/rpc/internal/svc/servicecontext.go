package svc

import (
	"GaussConsul/service/user/model"
	"GaussConsul/service/user/rpc/internal/config"
	"database/sql"
	_ "gitee.com/opengauss/openGauss-connector-go-pq"
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
