package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserService zrpc.RpcClientConf
	DataSource  string
	Table       string
	Cache       cache.CacheConf
}
