package cart

import (
	"context"

	v1 "hello/api/cart/v1"
	"hello/internal/service"
)

// List 处理 GET /cart/list
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	out, err := service.Cart().List(ctx, currentUserID(ctx))
	if err != nil {
		return nil, err
	}

	items := make([]v1.ListItem, 0, len(out.List))
	for _, item := range out.List {
		items = append(items, v1.ListItem{
			ID:        item.ID,
			GoodsID:   item.GoodsID,
			Name:      item.Name,
			Subtitle:  item.Subtitle,
			Cover:     item.Cover,
			Price:     item.Price,
			Quantity:  item.Quantity,
			Stock:     item.Stock,
			Subtotal:  item.Subtotal,
			Available: item.Available,
			Reason:    item.Reason,
		})
	}

	return &v1.ListRes{
		List:          items,
		TotalQuantity: out.TotalQuantity,
		TotalAmount:   out.TotalAmount,
		InvalidCount:  out.InvalidCount,
	}, nil
}
