package order

import (
	"context"

	v1 "hello/api/order/v1"
	"hello/internal/service"
)

// Cancel 处理 POST /order/cancel
func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	if err = service.Order().Cancel(ctx, currentUserID(ctx), req.ID); err != nil {
		return nil, err
	}
	return &v1.CancelRes{}, nil
}
