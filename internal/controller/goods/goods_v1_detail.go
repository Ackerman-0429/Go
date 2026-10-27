package goods

import (
	"context"
	"encoding/json"

	v1 "hello/api/goods/v1"
	"hello/internal/service"
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

	return &v1.DetailRes{
		ID:         item.Id,
		CategoryID: item.CategoryId,
		Name:       item.Name,
		Subtitle:   item.Subtitle,
		Price:      item.Price,
		Stock:      item.Stock,
		Cover:      item.Cover,
		Images:     images,
		Detail:     item.Detail,
		Sales:      item.Sales,
	}, nil
}
