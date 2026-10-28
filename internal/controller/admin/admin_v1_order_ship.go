package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/consts"
	"goshop/internal/service"
)

// OrderShip 处理 POST /admin/order/ship，已付款 → 已发货
func (c *ControllerV1) OrderShip(ctx context.Context, req *v1.OrderShipReq) (res *v1.OrderShipRes, err error) {
	if err = service.Order().AdminShip(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.OrderShipRes{
		ID:         req.ID,
		Status:     consts.OrderStatusShipped,
		StatusText: consts.OrderStatusText(consts.OrderStatusShipped),
	}, nil
}
