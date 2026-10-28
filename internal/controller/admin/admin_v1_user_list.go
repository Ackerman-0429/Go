package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// UserList 处理 GET /admin/user/list，支持关键字/角色/状态筛选
func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	list, total, err := service.User().AdminList(ctx, model.UserAdminListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
		Role:     req.Role,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	items := make([]v1.UserItem, 0, len(list))
	for _, u := range list {
		items = append(items, v1.UserItem{
			ID:         u.ID,
			Username:   u.Username,
			Nickname:   u.Nickname,
			Email:      u.Email,
			Phone:      u.Phone,
			Role:       u.Role,
			Status:     u.Status,
			OrderCount: u.OrderCount,
			TotalSpent: u.TotalSpent,
			CreatedAt:  formatTime(u.CreatedAt),
		})
	}

	return &v1.UserListRes{
		List:     items,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
