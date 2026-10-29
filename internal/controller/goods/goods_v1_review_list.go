package goods

import (
	"context"

	v1 "goshop/api/goods/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// ReviewList 处理 GET /goods/reviews，返回评价分页 + 评分汇总
func (c *ControllerV1) ReviewList(ctx context.Context, req *v1.ReviewListReq) (res *v1.ReviewListRes, err error) {
	list, total, err := service.Review().List(ctx, model.ReviewListInput{
		GoodsID:  req.GoodsID,
		Page:     req.Page,
		PageSize: req.PageSize,
		Rating:   req.Rating,
	})
	if err != nil {
		return nil, err
	}

	// 汇总不受星级筛选影响，保证「全部 12 条 / 5星 8 条」这种数字始终是全局的
	summary, err := service.Review().Summary(ctx, req.GoodsID)
	if err != nil {
		return nil, err
	}

	items := make([]v1.ReviewItem, 0, len(list))
	for _, r := range list {
		items = append(items, v1.ReviewItem{
			ID:        r.ID,
			Nickname:  r.Nickname,
			Rating:    r.Rating,
			Content:   r.Content,
			Anonymous: r.Anonymous,
			CreatedAt: formatTime(r.CreatedAt),
		})
	}

	return &v1.ReviewListRes{
		List:      items,
		Total:     total,
		Page:      req.Page,
		PageSize:  req.PageSize,
		AvgRating: summary.AvgRating,
		Count:     summary.Count,
		Counts:    summary.Counts[:],
	}, nil
}
