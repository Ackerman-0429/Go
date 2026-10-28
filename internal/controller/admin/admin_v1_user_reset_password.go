package admin

import (
	"context"

	v1 "goshop/api/admin/v1"
	"goshop/internal/service"
)

// UserResetPassword 处理 POST /admin/user/reset-password
func (c *ControllerV1) UserResetPassword(ctx context.Context, req *v1.UserResetPasswordReq) (res *v1.UserResetPasswordRes, err error) {
	if err = service.User().AdminResetPassword(ctx, req.ID, req.Password); err != nil {
		return nil, err
	}
	return &v1.UserResetPasswordRes{}, nil
}
