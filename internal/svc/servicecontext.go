package svc

import (
	"github.com/xemxx/go-zero-template/internal/config"
	"github.com/xemxx/go-zero-template/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
