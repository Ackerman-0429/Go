package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// Stats 处理 GET /admin/stats，后台首页概览
func (c *ControllerV1) Stats(ctx context.Context, req *v1.StatsReq) (res *v1.StatsRes, err error) {
	stats, err := service.Order().Stats(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.StatsRes{
		TodayOrderCount:  stats.TodayOrderCount,
		TodaySales:       stats.TodaySales,
		PendingShipCount: stats.PendingShipCount,
		TotalOrderCount:  stats.TotalOrderCount,
		OnSaleGoodsCount: stats.OnSaleGoodsCount,
		UserCount:        stats.UserCount,
	}, nil
}
