package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"hello/internal/controller/cart"
	"hello/internal/controller/category"
	"hello/internal/controller/goods"
	"hello/internal/controller/hello"
	"hello/internal/controller/order"
	"hello/internal/controller/user"
	"hello/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 公开接口：无需登录
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
					goods.NewV1(),
					category.NewV1(),
				)
			})
			// 需登录接口：Auth 中间件校验会话并写入用户 ID
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, middleware.Auth)
				group.Bind(
					cart.NewV1(),
					order.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
