package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "hello/api/user/v1"
)

// Logout 处理 POST /user/logout，清除会话中的用户标识
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := g.RequestFromCtx(ctx)
	if err = r.Session.Remove("userId"); err != nil {
		return nil, err
	}
	return &v1.LogoutRes{}, nil
}
