package order

import (
	"context"

	v1 "goshop/api/order/v1"
	"goshop/internal/consts"
	"goshop/internal/service"
)

// Pay 处理 POST /order/pay，模拟支付
func (c *ControllerV1) Pay(ctx context.Context, req *v1.PayReq) (res *v1.PayRes, err error) {
	paidAt, err := service.Order().Pay(ctx, currentUserID(ctx), req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.PayRes{
		Status:     consts.OrderStatusPaid,
		StatusText: consts.OrderStatusText(consts.OrderStatusPaid),
		PaidAt:     formatTime(paidAt),
	}, nil
}
