// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"hello/api/goods/v1"
)

type IGoodsV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
}
