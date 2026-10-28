package order

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"hello/internal/logic/middleware"
	"hello/internal/service"
)

// ControllerV1 订单控制器。全部接口都要求登录态，鉴权由 Auth 中间件统一完成。
type ControllerV1 struct{}

// NewV1 创建订单控制器
func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// currentUserID 从 Auth 中间件写入的上下文变量中取当前登录用户
func currentUserID(ctx context.Context) uint64 {
	return g.RequestFromCtx(ctx).GetCtxVar(middleware.CtxUserID).Uint64()
}

// statusText 订单状态文案
func statusText(status uint) string {
	switch status {
	case service.OrderStatusPending:
		return "待付款"
	case service.OrderStatusPaid:
		return "已付款"
	case service.OrderStatusShipped:
		return "已发货"
	case service.OrderStatusFinished:
		return "已完成"
	case service.OrderStatusCancelled:
		return "已取消"
	}
	return "未知状态"
}

// formatTime 时间格式化，空值返回空字符串
func formatTime(t *gtime.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("Y-m-d H:i:s")
}
