package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// CategoryDelete 处理 POST /admin/category/delete
func (c *ControllerV1) CategoryDelete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error) {
	if err = service.Category().AdminDelete(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.CategoryDeleteRes{}, nil
}
