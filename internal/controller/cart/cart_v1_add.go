package cart

import (
	"context"

	v1 "hello/api/cart/v1"
	"hello/internal/service"
)

// Add 处理 POST /cart/add
func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	id, quantity, err := service.Cart().Add(ctx, service.CartAddInput{
		UserID:   currentUserID(ctx),
		GoodsID:  req.GoodsID,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddRes{ID: id, Quantity: quantity}, nil
}
