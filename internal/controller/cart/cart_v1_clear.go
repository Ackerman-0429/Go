package cart

import (
	"context"

	v1 "goshop/api/cart/v1"
	"goshop/internal/service"
)

// Clear 处理 POST /cart/clear
func (c *ControllerV1) Clear(ctx context.Context, req *v1.ClearReq) (res *v1.ClearRes, err error) {
	deleted, err := service.Cart().Clear(ctx, currentUserID(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.ClearRes{Deleted: deleted}, nil
}
