package admin

import (
	"context"

	"goshop/api/admin/v1"
	"goshop/internal/service"
)

// ReviewStatus 处理 POST /admin/review/status，显示或隐藏一条评价
func (c *ControllerV1) ReviewStatus(ctx context.Context, req *v1.ReviewStatusReq) (res *v1.ReviewStatusRes, err error) {
	if err = service.Review().AdminSetStatus(ctx, req.ID, req.Status); err != nil {
		return nil, err
	}
	return &v1.ReviewStatusRes{ID: req.ID, Status: req.Status}, nil
}
