package order

import (
	"context"

	v1 "hello/api/order/v1"
	"hello/internal/service"
)

// Detail 处理 GET /order/detail
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	out, err := service.Order().Detail(ctx, currentUserID(ctx), req.ID)
	if err != nil {
		return nil, err
	}

	items := make([]v1.DetailItem, 0, len(out.Items))
	for _, it := range out.Items {
		items = append(items, v1.DetailItem{
			GoodsID:   it.GoodsId,
			GoodsName: it.GoodsName,
			Cover:     it.GoodsCover,
			Price:     it.GoodsPrice,
			Quantity:  it.Quantity,
			Subtotal:  it.Subtotal,
		})
	}

	o := out.Order
	return &v1.DetailRes{
		ID:              o.Id,
		OrderNo:         o.OrderNo,
		TotalAmount:     o.TotalAmount,
		Status:          o.Status,
		StatusText:      statusText(o.Status),
		ReceiverName:    o.ReceiverName,
		ReceiverPhone:   o.ReceiverPhone,
		ReceiverAddress: o.ReceiverAddress,
		Remark:          o.Remark,
		PaidAt:          formatTime(o.PaidAt),
		CreatedAt:       formatTime(o.CreatedAt),
		Items:           items,
	}, nil
}
