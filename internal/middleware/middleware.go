// Package middleware 提供 HTTP 层中间件：登录态校验等，与业务逻辑解耦。
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"

	"goshop/internal/consts"
	"goshop/internal/service"
)

const (
	// SessionUserID Session 中存放当前用户 ID 的键，与用户模块登录时写入的键一致
	SessionUserID = "userId"

	// CtxUserID 校验通过后写入请求上下文的用户 ID，供 controller 读取
	CtxUserID = "ctxUserId"

	// CtxUserRole 登录用户角色，由本中间件写入，Admin 中间件直接复用，避免重复查库
	CtxUserRole = "ctxUserRole"
)

// Auth 登录态校验中间件。
// 除了「有没有登录」，还会回查一次 user 表确认账号仍存在且未被禁用：
// 否则管理员禁用某个账号后，对方拿着旧 Session 依然能下单。
// 角色也在这里读出并放进上下文，所以后台降权同样是即时生效的。
func Auth(r *ghttp.Request) {
	sessionVar, err := r.Session.Get(SessionUserID)
	if err != nil {
		r.SetError(gerror.NewCode(gcode.CodeInternalError, "会话读取失败"))
		return
	}

	userID := sessionVar.Uint64()
	if userID == 0 {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "请先登录"))
		return
	}

	user, err := service.User().GetById(r.Context(), userID)
	if err != nil || user == nil {
		// 账号已被删除：清掉脏 Session，避免每次请求都带着无效 ID
		_ = r.Session.Remove(SessionUserID)
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "请先登录"))
		return
	}
	if user.Status != consts.UserStatusNormal {
		_ = r.Session.Remove(SessionUserID)
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用"))
		return
	}

	r.SetCtxVar(CtxUserID, userID)
	r.SetCtxVar(CtxUserRole, user.Role)
	r.Middleware.Next()
}
