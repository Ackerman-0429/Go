package user

import (
	"context"

	v1 "goshop/api/user/v1"
	"goshop/internal/service"
)

// Info 处理 GET /user/info，从 Session 中识别当前用户
func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	userId, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := service.User().GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.InfoRes{
		ID:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
		Role:     user.Role,
	}, nil
}
