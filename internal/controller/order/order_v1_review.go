package order

import (
	"context"

	v1 "goshop/api/order/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Review 处理 POST /order/review，对自己已完成订单中的商品发表评价
func (c *ControllerV1) Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error) {
	id, err := service.Review().Create(ctx, currentUserID(ctx), model.ReviewCreateInput{
		OrderItemID: req.OrderItemID,
		Rating:      req.Rating,
		Content:     req.Content,
		Anonymous:   req.Anonymous,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ReviewRes{ID: id}, nil
}
