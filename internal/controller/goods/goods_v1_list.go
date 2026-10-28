package goods

import (
	"context"

	v1 "hello/api/goods/v1"
	"hello/internal/model"
	"hello/internal/service"
)

// List 处理 GET /goods/list
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, total, err := service.Goods().List(ctx, model.GoodsListInput{
		Page:       req.Page,
		PageSize:   req.PageSize,
		CategoryID: req.CategoryID,
		Keyword:    req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.ListItem, 0, len(list))
	for _, item := range list {
		items = append(items, v1.ListItem{
			ID:         item.Id,
			Name:       item.Name,
			Subtitle:   item.Subtitle,
			Price:      item.Price,
			Cover:      item.Cover,
			Sales:      item.Sales,
			Stock:      item.Stock,
			CategoryID: item.CategoryId,
		})
	}

	return &v1.ListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
