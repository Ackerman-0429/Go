package category

import (
	"context"

	v1 "goshop/api/category/v1"
	"goshop/internal/service"
)

// List 处理 GET /category/list
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	list, err := service.Category().List(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]v1.ListItem, 0, len(list))
	for _, item := range list {
		items = append(items, v1.ListItem{
			ID:       item.Id,
			ParentID: item.ParentId,
			Name:     item.Name,
			Sort:     item.Sort,
		})
	}

	return &v1.ListRes{List: items}, nil
}
