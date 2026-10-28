package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/consts"
	"goshop/internal/service"
)

// OrderDetail 处理 GET /admin/order/detail，管理员可查看任意用户的订单
func (c *ControllerV1) OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	d, err := service.Order().AdminDetail(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 明细件数在这里顺手累加，省得前端再算一遍
	items := make([]v1.OrderGoodsItem, 0, len(d.Items))
	var quantity uint
	for _, it := range d.Items {
		quantity += it.Quantity
		items = append(items, v1.OrderGoodsItem{
			GoodsID:   it.GoodsId,
			GoodsName: it.GoodsName,
			Cover:     it.GoodsCover,
			Price:     it.GoodsPrice,
			Quantity:  it.Quantity,
			Subtotal:  it.Subtotal,
		})
	}

	o := d.Order
	return &v1.OrderDetailRes{
		ID:              o.Id,
		OrderNo:         o.OrderNo,
		UserID:          o.UserId,
		Username:        d.Username,
		TotalAmount:     o.TotalAmount,
		Status:          o.Status,
		StatusText:      consts.OrderStatusText(o.Status),
		Quantity:        quantity,
		ReceiverName:    o.ReceiverName,
		ReceiverPhone:   o.ReceiverPhone,
		ReceiverAddress: o.ReceiverAddress,
		Remark:          o.Remark,
		CreatedAt:       formatTime(o.CreatedAt),
		PaidAt:          formatTime(o.PaidAt),
		Items:           items,
	}, nil
}
