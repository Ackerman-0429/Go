package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// GoodsSave 处理 POST /admin/goods/save，ID 为空即新建
func (c *ControllerV1) GoodsSave(ctx context.Context, req *v1.GoodsSaveReq) (res *v1.GoodsSaveRes, err error) {
	id, err := service.Goods().AdminSave(ctx, model.GoodsSaveInput{
		ID:         req.ID,
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Subtitle:   req.Subtitle,
		Price:      req.Price,
		Stock:      req.Stock,
		Cover:      req.Cover,
		Images:     req.Images,
		Detail:     req.Detail,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsSaveRes{ID: id}, nil
}
