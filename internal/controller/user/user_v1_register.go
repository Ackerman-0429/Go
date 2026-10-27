package user

import (
	"context"

	v1 "hello/api/user/v1"
	"hello/internal/service"
)

// Register 处理 POST /user/register
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	id, err := service.User().Register(ctx, req.Username, req.Password, req.Nickname)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{ID: id}, nil
}
