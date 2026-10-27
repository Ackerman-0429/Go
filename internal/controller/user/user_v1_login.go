package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "hello/api/user/v1"
	"hello/internal/service"
)

// Login 处理 POST /user/login，校验通过后写入 Session
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := service.User().Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	r := g.RequestFromCtx(ctx)
	if err = r.Session.Set("userId", user.Id); err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		ID:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
	}, nil
}
