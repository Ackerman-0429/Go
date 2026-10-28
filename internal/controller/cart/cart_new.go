package cart

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"hello/internal/logic/middleware"
)

// ControllerV1 购物车控制器。全部接口都要求登录态，鉴权由 Auth 中间件统一完成。
type ControllerV1 struct{}

// NewV1 创建购物车控制器
func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// currentUserID 从 Auth 中间件写入的上下文变量中取当前登录用户
func currentUserID(ctx context.Context) uint64 {
	return g.RequestFromCtx(ctx).GetCtxVar(middleware.CtxUserID).Uint64()
}
