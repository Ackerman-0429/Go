package user

import (
	"context"

	v1 "goshop/api/user/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Password 处理 POST /user/password，用户凭原密码修改自己的登录密码
func (c *ControllerV1) Password(ctx context.Context, req *v1.PasswordReq) (res *v1.PasswordRes, err error) {
	userId, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}

	if err = service.User().ChangePassword(ctx, model.UserPasswordInput{
		ID:          userId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}); err != nil {
		return nil, err
	}
	return &v1.PasswordRes{}, nil
}
