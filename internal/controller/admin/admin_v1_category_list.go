package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// CategoryList 处理 GET /admin/category/list，含停用分类
func (c *ControllerV1) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	list, err := service.Category().AdminList(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]v1.CategoryItem, 0, len(list))
	for _, c := range list {
		items = append(items, v1.CategoryItem{
			ID:         uint64(c.ID),
			ParentID:   uint64(c.ParentID),
			ParentName: c.ParentName,
			Name:       c.Name,
			Sort:       c.Sort,
			Status:     c.Status,
			GoodsCount: c.GoodsCount,
			Children:   c.Children,
			CreatedAt:  formatTime(c.CreatedAt),
		})
	}

	return &v1.CategoryListRes{List: items}, nil
}
