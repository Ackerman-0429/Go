package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"hello/internal/controller/category"
	"hello/internal/controller/goods"
	"hello/internal/controller/hello"
	"hello/internal/controller/user"

	// 导入 logic 层，触发其 init() 完成 service 实现的注册
	_ "hello/internal/logic/category"
	_ "hello/internal/logic/goods"
	_ "hello/internal/logic/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
					goods.NewV1(),
					category.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
