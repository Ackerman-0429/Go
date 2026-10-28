package order

import (
	"context"

	v1 "goshop/api/order/v1"
	"goshop/internal/consts"
	"goshop/internal/model"
	"goshop/internal/service"
)

// List 处理 GET /order/list
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := service.Order().List(ctx, model.OrderListInput{
		UserID:   currentUserID(ctx),
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.ListItem, 0, len(list))
	for _, o := range list {
		items = append(items, v1.ListItem{
			ID:          o.ID,
			OrderNo:     o.OrderNo,
			TotalAmount: o.TotalAmount,
			Status:      o.Status,
			StatusText:  consts.OrderStatusText(o.Status),
			Quantity:    o.Quantity,
			ItemCount:   o.ItemCount,
			FirstName:   o.FirstName,
			FirstCover:  o.FirstCover,
			CreatedAt:   formatTime(o.CreatedAt),
		})
	}

	return &v1.ListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
