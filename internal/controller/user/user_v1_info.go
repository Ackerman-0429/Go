package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "hello/api/user/v1"
	"hello/internal/service"
)

// Info 处理 GET /user/info，从 Session 中识别当前用户
func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	r := g.RequestFromCtx(ctx)
	sessionVar, err := r.Session.Get("userId")
	if err != nil {
		return nil, err
	}
	userId := sessionVar.Uint64()
	if userId == 0 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "请先登录")
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
	}, nil
}
