// Package goods 定义商品模块的接口契约，由 controller 层实现。
package goods

import (
	"context"

	v1 "hello/api/goods/v1"
)

// IGoodsV1 商品模块接口
type IGoodsV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
}
