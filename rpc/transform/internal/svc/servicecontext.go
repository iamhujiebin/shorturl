package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"shorturl/rpc/transform/internal/config"
	mongoModel "shorturl/rpc/transform/model/mongo"
	mysqlModel "shorturl/rpc/transform/model/mysql"
	"shorturl/rpc/userservice/userservice"
)

type ServiceContext struct {
	c           config.Config
	UserService userservice.UserService
	Model       *mysqlModel.ShorturlModel
	MongoModel  *mongoModel.ShortUrlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:           c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
		Model:       mysqlModel.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
		MongoModel:  &mongoModel.ShortUrlModel{},
	}
}
