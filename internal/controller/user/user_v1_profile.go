package user

import (
	"context"

	v1 "goshop/api/user/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Profile 处理 POST /user/profile，用户修改自己的昵称、邮箱、手机号
func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	userId, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := service.User().UpdateProfile(ctx, model.UserProfileInput{
		ID:       userId,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &v1.ProfileRes{
		ID:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}, nil
}
