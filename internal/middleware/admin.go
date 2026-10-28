package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"goshop/internal/consts"
	"goshop/internal/service"
)

// Admin 管理员鉴权中间件，必须挂在 Auth 之后。
// Auth 只保证「已登录」，这里再确认「是管理员」且「账号未被禁用」。
// 角色存在数据库而不是 Session 里，管理员被降权后立即生效，无需等会话过期。
func Admin(r *ghttp.Request) {
	userID := r.GetCtxVar(CtxUserID).Uint64()
	if userID == 0 {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "请先登录"))
		return
	}

	user, err := service.User().GetById(r.Context(), userID)
	if err != nil {
		r.SetError(err)
		return
	}
	if user.Status != consts.UserStatusNormal {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用"))
		return
	}
	if user.Role != consts.UserRoleAdmin {
		r.SetError(gerror.NewCode(consts.CodeNotAdmin))
		return
	}

	r.Middleware.Next()
}
