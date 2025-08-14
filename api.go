package main

import (
	"flag"

	"ks_api_service/internal/config"
	"ks_api_service/internal/server"
	"ks_api_service/internal/svc"
	"ks_api_service/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	path = flag.String("f", "etc/api.yaml", "the config file path")
	key  = flag.String("k", "api.yaml", "the remote key name")
)

func main() {
	flag.Parse()

	c := config.Config{}
	c.Parse(*path, *key)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		api.RegisterApiServer(grpcServer, server.NewApiServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
