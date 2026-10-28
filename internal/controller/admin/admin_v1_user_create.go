package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// UserCreate 处理 POST /admin/user/create，后台直接建号。
// 密码由管理员指定，与注册接口一样用 bcrypt 哈希后入库。
func (c *ControllerV1) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	id, err := service.User().AdminCreate(ctx, model.UserCreateInput{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserCreateRes{ID: id}, nil
}
