// Code generated by goctl. DO NOT EDIT!
// Source: userservice.proto

package main

import (
	"flag"
	"fmt"

	"shorturl/rpc/userservice/internal/config"
	"shorturl/rpc/userservice/internal/server"
	"shorturl/rpc/userservice/internal/svc"
	user_service "shorturl/rpc/userservice/pb"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/userservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	userServiceSrv := server.NewUserServiceServer(ctx)

	s, err := zrpc.NewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_service.RegisterUserServiceServer(grpcServer, userServiceSrv)
	})
	logx.Must(err)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
