package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Transform   zrpc.RpcClientConf
	UserService zrpc.RpcClientConf
	Log         struct {
		LogMode string
		Path    string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
