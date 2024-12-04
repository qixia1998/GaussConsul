package logic

import (
	"GaussConsul/service/gauss/rpc/types/gauss"
	"context"

	"GaussConsul/service/gauss/api/internal/svc"
	"GaussConsul/service/gauss/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() (resp *types.GaussHealthCheckResp, err error) {
	res, err := l.svcCtx.GaussRpc.HealthCheck(l.ctx, &gauss.HealthCheckRequest{})
	if err != nil {
		return nil, err
	}
	return &types.GaussHealthCheckResp{
		Status: res.Status,
	}, nil
}
