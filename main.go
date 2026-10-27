// main.go —— GoFrame v2 入门示例
//
// 运行：
//   go run .
// 访问：
//   http://127.0.0.1:8000/hello?name=Ackerman
//
// GoFrame 的"规范路由"（Standard Router）核心思路：
// 用 g.Meta 在结构体上声明路径和方法，框架自动完成
// 参数绑定 -> 参数校验 -> 调用方法 -> 结果序列化。
package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// HelloReq 请求参数。g.Meta 声明：GET /hello
type HelloReq struct {
	g.Meta `path:"/hello" method:"get" summary:"打招呼接口"`
	Name   string `v:"required#name 不能为空" dc:"你的名字"`
}

// HelloRes 响应数据
type HelloRes struct {
	Message string `json:"message" dc:"问候语"`
}

// HelloController 控制器，方法签名固定为 (ctx, req) -> (res, err)
type HelloController struct{}

// Hello 处理 GET /hello
func (c *HelloController) Hello(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	return &HelloRes{
		Message: fmt.Sprintf("你好，%s！运行环境 %s", req.Name, runtime.Version()),
	}, nil
}

func main() {
	s := g.Server()
	s.SetPort(8000)

	// 开启自动生成的 OpenAPI 文档（v2.10 默认关闭）
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")

	s.Group("/", func(group *ghttp.RouterGroup) {
		// 统一响应中间件：自动包成 {code, message, data}
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(new(HelloController))
	})

	s.Run()
}
