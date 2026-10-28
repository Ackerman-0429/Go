package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/consts"
	"goshop/internal/model"
	"goshop/internal/service"
)

// OrderList 处理 GET /admin/order/list，跨用户查询全站订单
func (c *ControllerV1) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	list, total, err := service.Order().AdminList(ctx, model.OrderAdminListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.OrderItem, 0, len(list))
	for _, o := range list {
		items = append(items, v1.OrderItem{
			ID:              o.ID,
			OrderNo:         o.OrderNo,
			UserID:          o.UserID,
			Username:        o.Username,
			TotalAmount:     o.TotalAmount,
			Status:          o.Status,
			StatusText:      consts.OrderStatusText(o.Status),
			Quantity:        o.Quantity,
			ItemCount:       o.ItemCount,
			FirstName:       o.FirstName,
			ReceiverName:    o.ReceiverName,
			ReceiverPhone:   o.ReceiverPhone,
			ReceiverAddress: o.ReceiverAddress,
			CreatedAt:       formatTime(o.CreatedAt),
			PaidAt:          formatTime(o.PaidAt),
		})
	}

	return &v1.OrderListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
