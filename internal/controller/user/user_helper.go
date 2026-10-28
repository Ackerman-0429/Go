package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/middleware"
)

// currentUserID 从 Session 读当前登录用户 ID。
// 用户控制器整体挂在公开分组（登录/注册本身不需要登录态），
// 所以需要登录的接口只能在这里单独校验，不能依赖 Auth 中间件。
func currentUserID(ctx context.Context) (uint64, error) {
	sessionVar, err := g.RequestFromCtx(ctx).Session.Get(middleware.SessionUserID)
	if err != nil {
		return 0, err
	}
	userId := sessionVar.Uint64()
	if userId == 0 {
		return 0, gerror.NewCode(gcode.CodeNotAuthorized, "请先登录")
	}
	return userId, nil
}
