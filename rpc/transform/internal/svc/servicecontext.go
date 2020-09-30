package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"shorturl/rpc/transform/internal/config"
	mongoModel "shorturl/rpc/transform/model/mongo"
	mysqlModel "shorturl/rpc/transform/model/mysql"
)

type ServiceContext struct {
	c          config.Config
	Model      *mysqlModel.ShorturlModel
	MongoModel *mongoModel.ShortUrlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:          c,
		Model:      mysqlModel.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
		MongoModel: &mongoModel.ShortUrlModel{},
	}
}
