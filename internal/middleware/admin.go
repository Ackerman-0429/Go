package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"goshop/internal/consts"
)

// Admin 管理员鉴权中间件，必须挂在 Auth 之后。
// Auth 已确认账号存在且未被禁用，并把角色写进上下文，
// 这里只判断角色，不再重复查库；角色存 DB，所以降权也是即时生效的。
func Admin(r *ghttp.Request) {
	userID := r.GetCtxVar(CtxUserID).Uint64()
	if userID == 0 {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "请先登录"))
		return
	}
	if r.GetCtxVar(CtxUserRole).Uint() != consts.UserRoleAdmin {
		r.SetError(gerror.NewCode(consts.CodeNotAdmin))
		return
	}

	r.Middleware.Next()
}
