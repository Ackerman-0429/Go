package main

import (
	// 空白导入 MySQL 驱动。GoFrame v2 起数据库驱动不再内置，缺失会报
	// cannot find database driver for specified database type "mysql"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	// 空白导入 logic 聚合包（由 gf gen service 维护），
	// 触发各业务模块 init() 完成 service 实现注册
	_ "hello/internal/logic"

	"hello/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
