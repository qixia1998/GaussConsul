package logic

import (
	"context"
	"database/sql"
	"fmt"

	"GaussConsul/service/gauss/rpc/internal/svc"
	"GaussConsul/service/gauss/rpc/types/gauss"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HealthCheckLogic) HealthCheck(in *gauss.HealthCheckRequest) (*gauss.HealthCheckResponse, error) {

	conn, err := sql.Open("opengauss", l.svcCtx.Config.GaussDB.DataSource)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		return &gauss.HealthCheckResponse{
			Status: fmt.Sprintf("GaussDB health check failed %v", err),
		}, nil
	}

	return &gauss.HealthCheckResponse{
		Status: "GaussDB is healthy",
	}, nil
}
