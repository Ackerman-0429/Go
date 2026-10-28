// Package cart 定义购物车模块的接口契约，由 controller 层实现。
package cart

import (
	"context"

	v1 "hello/api/cart/v1"
)

// ICartV1 购物车模块接口，全部方法要求登录态
type ICartV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Clear(ctx context.Context, req *v1.ClearReq) (res *v1.ClearRes, err error)
}
