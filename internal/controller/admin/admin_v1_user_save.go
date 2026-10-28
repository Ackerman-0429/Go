package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// UserSave 处理 POST /admin/user/save，改昵称/角色/状态。
// 传入当前登录管理员 ID，用于拦住「把自己降级或禁用」的操作。
func (c *ControllerV1) UserSave(ctx context.Context, req *v1.UserSaveReq) (res *v1.UserSaveRes, err error) {
	if _, err = service.User().AdminSave(ctx, model.UserSaveInput{
		ID:       req.ID,
		Nickname: req.Nickname,
		Role:     req.Role,
		Status:   req.Status,
	}, currentUserID(ctx)); err != nil {
		return nil, err
	}
	return &v1.UserSaveRes{
		ID:     req.ID,
		Role:   req.Role,
		Status: req.Status,
	}, nil
}
