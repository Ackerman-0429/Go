package admin

import (
	"context"

	"goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// ReviewList 处理 GET /admin/review/list，跨商品查询全站评价
func (c *ControllerV1) ReviewList(ctx context.Context, req *v1.ReviewListReq) (res *v1.ReviewListRes, err error) {
	list, total, stats, err := service.Review().AdminList(ctx, model.ReviewAdminListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Status:   req.Status,
		Rating:   req.Rating,
		GoodsID:  req.GoodsId,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.ReviewAdminItem, 0, len(list))
	for _, r := range list {
		items = append(items, v1.ReviewAdminItem{
			ID:        r.ID,
			GoodsID:   r.GoodsID,
			GoodsName: r.GoodsName,
			UserID:    r.UserID,
			Username:  r.Username,
			Nickname:  r.Nickname,
			Rating:    r.Rating,
			Content:   r.Content,
			Anonymous: r.Anonymous,
			Status:    r.Status,
			CreatedAt: formatTime(r.CreatedAt),
		})
	}

	return &v1.ReviewListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
		Visible:  stats.Visible,
		Hidden:   stats.Hidden,
	}, nil
}
