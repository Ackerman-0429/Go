package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// GoodsList 处理 GET /admin/goods/list，含下架商品
func (c *ControllerV1) GoodsList(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error) {
	list, total, err := service.Goods().AdminList(ctx, model.GoodsAdminListInput{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Keyword:    req.Keyword,
		CategoryID: req.CategoryID,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.GoodsItem, 0, len(list))
	for _, g := range list {
		items = append(items, v1.GoodsItem{
			ID:         g.Id,
			CategoryID: uint64(g.CategoryId),
			Name:       g.Name,
			Subtitle:   g.Subtitle,
			Price:      g.Price,
			Stock:      g.Stock,
			Cover:      g.Cover,
			Images:     g.Images,
			Detail:     g.Detail,
			Sales:      g.Sales,
			Status:     g.Status,
			CreatedAt:  formatTime(g.CreatedAt),
		})
	}

	return &v1.GoodsListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
