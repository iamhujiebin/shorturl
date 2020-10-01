package svc

import (
	"shorturl/rpc/userservice/internal/config"
	"shorturl/rpc/userservice/model/mongo"
)

type ServiceContext struct {
	c         config.Config
	UserModel mongo.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:         c,
		UserModel: mongo.UserModel{},
	}
}
