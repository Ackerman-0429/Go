package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goshop/internal/controller/cart"
	"goshop/internal/controller/category"
	"goshop/internal/controller/goods"
	"goshop/internal/controller/hello"
	"goshop/internal/controller/order"
	"goshop/internal/controller/user"
	"goshop/internal/middleware"
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
