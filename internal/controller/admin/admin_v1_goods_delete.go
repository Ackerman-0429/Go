package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// GoodsDelete 处理 POST /admin/goods/delete
func (c *ControllerV1) GoodsDelete(ctx context.Context, req *v1.GoodsDeleteReq) (res *v1.GoodsDeleteRes, err error) {
	if err = service.Goods().AdminDelete(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.GoodsDeleteRes{}, nil
}
