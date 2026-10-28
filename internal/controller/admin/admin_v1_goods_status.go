package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// GoodsStatus 处理 POST /admin/goods/status，上下架
func (c *ControllerV1) GoodsStatus(ctx context.Context, req *v1.GoodsStatusReq) (res *v1.GoodsStatusRes, err error) {
	if err = service.Goods().AdminSetStatus(ctx, req.ID, req.Status); err != nil {
		return nil, err
	}
	return &v1.GoodsStatusRes{ID: req.ID, Status: req.Status}, nil
}
