package svc

import (
	"GaussConsul/service/gauss/api/internal/config"
	"GaussConsul/service/gauss/rpc/gaussclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	GaussRpc gaussclient.Gauss
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GaussRpc: gaussclient.NewGauss(zrpc.MustNewClient(c.GaussRpc)),
	}
}
