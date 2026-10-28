package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"goshop/internal/middleware"
)

// formatTime 时间格式化，空值返回空字符串，避免前端显示 null
func formatTime(t *gtime.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("Y-m-d H:i:s")
}

// currentUserID 从 Auth 中间件写入的上下文变量中取当前登录管理员 ID
func currentUserID(ctx context.Context) uint64 {
	return g.RequestFromCtx(ctx).GetCtxVar(middleware.CtxUserID).Uint64()
}
