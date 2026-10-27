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

	// GoFrame v2 起数据库驱动不再内置，必须显式空白导入才会注册
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

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

// DbReq 无入参
type DbReq struct {
	g.Meta `path:"/db" method:"get" summary:"数据库连通性测试"`
}

// DbRes 数据库信息
type DbRes struct {
	Version  string `json:"version"  dc:"MySQL 版本"`
	Database string `json:"database" dc:"当前数据库"`
	Tables   int    `json:"tables"   dc:"表数量"`
}

// Db 查询数据库，验证 GoFrame 与 MySQL 的链路
func (c *HelloController) Db(ctx context.Context, req *DbReq) (res *DbRes, err error) {
	db := g.DB()

	one, err := db.GetOne(ctx, "SELECT VERSION() AS version, DATABASE() AS dbname")
	if err != nil {
		return nil, err
	}

	count, err := db.GetCount(ctx,
		"SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE()")
	if err != nil {
		return nil, err
	}

	return &DbRes{
		Version:  one["version"].String(),
		Database: one["dbname"].String(),
		Tables:   count,
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
