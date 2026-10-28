package cart

import (
	"context"

	v1 "goshop/api/cart/v1"
	"goshop/internal/service"
)

// Update 处理 POST /cart/update
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	if err = service.Cart().UpdateQuantity(ctx, currentUserID(ctx), req.ID, req.Quantity); err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Quantity: req.Quantity}, nil
}
