package goods

import (
	"context"
	"encoding/json"

	v1 "goshop/api/goods/v1"
	"goshop/internal/service"
)

// Detail 处理 GET /goods/detail
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	item, err := service.Goods().GetById(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// images 在库里存的是 JSON 数组字符串，这里解析成真正的数组返回
	images := make([]string, 0)
	if item.Images != "" {
		_ = json.Unmarshal([]byte(item.Images), &images)
	}

	// 评分汇总跟详情一起带回，前端不用再请求一次评价接口才画得出星级
	summary, err := service.Review().Summary(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &v1.DetailRes{
		ID:          item.Id,
		CategoryID:  item.CategoryId,
		Name:        item.Name,
		Subtitle:    item.Subtitle,
		Price:       item.Price,
		Stock:       item.Stock,
		Cover:       item.Cover,
		Images:      images,
		Detail:      item.Detail,
		Sales:       item.Sales,
		Rating:      summary.AvgRating,
		ReviewCount: summary.Count,
	}, nil
}
