// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"

	g1 "github.com/xemxx/go-zero-template/internal/handler/g1"
	"github.com/xemxx/go-zero-template/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/ping",
					Handler: g1.PingHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: g1.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: g1.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: g1.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/form/example",
					Handler: g1.FormExampleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/path/example/:id",
					Handler: g1.PathExampleHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
		rest.WithMaxBytes(1048576),
	)
}
