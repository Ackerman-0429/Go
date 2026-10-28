package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// CategorySave 处理 POST /admin/category/save，ID 为空即新建
func (c *ControllerV1) CategorySave(ctx context.Context, req *v1.CategorySaveReq) (res *v1.CategorySaveRes, err error) {
	id, err := service.Category().AdminSave(ctx, model.CategorySaveInput{
		ID:       req.ID,
		ParentID: req.ParentID,
		Name:     req.Name,
		Sort:     req.Sort,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CategorySaveRes{ID: id}, nil
}
