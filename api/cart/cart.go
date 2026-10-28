// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package cart

import (
	"context"

	"goshop/api/cart/v1"
)

type ICartV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Clear(ctx context.Context, req *v1.ClearReq) (res *v1.ClearRes, err error)
}
