// Package order 定义订单模块的接口契约，由 controller 层实现。
package order

import (
	"context"

	v1 "hello/api/order/v1"
)

// IOrderV1 订单模块接口，全部方法要求登录态
type IOrderV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	Pay(ctx context.Context, req *v1.PayReq) (res *v1.PayRes, err error)
}
