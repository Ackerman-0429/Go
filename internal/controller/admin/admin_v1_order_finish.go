package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/consts"
	"goshop/internal/service"
)

// OrderFinish 处理 POST /admin/order/finish，已发货 → 已完成
func (c *ControllerV1) OrderFinish(ctx context.Context, req *v1.OrderFinishReq) (res *v1.OrderFinishRes, err error) {
	if err = service.Order().AdminFinish(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.OrderFinishRes{
		ID:         req.ID,
		Status:     consts.OrderStatusFinished,
		StatusText: consts.OrderStatusText(consts.OrderStatusFinished),
	}, nil
}
