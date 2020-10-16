package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformer"
	"shorturl/rpc/userservice/userservice"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
		//UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
	}
}
