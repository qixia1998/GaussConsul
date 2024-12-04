package main

import (
	"flag"
	"fmt"

	"GaussConsul/service/gauss/rpc/internal/config"
	"GaussConsul/service/gauss/rpc/internal/server"
	"GaussConsul/service/gauss/rpc/internal/svc"
	"GaussConsul/service/gauss/rpc/types/gauss"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gauss.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		gauss.RegisterGaussServer(grpcServer, server.NewGaussServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// register service to consul
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
