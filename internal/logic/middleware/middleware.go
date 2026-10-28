// Package middleware 提供 HTTP 中间件，目前包含登录态校验。
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	// SessionUserID Session 中存放当前用户 ID 的键，与用户模块登录时写入的键一致
	SessionUserID = "userId"

	// CtxUserID 校验通过后写入请求上下文的用户 ID，供 controller 读取
	CtxUserID = "ctxUserId"
)

// Auth 登录态校验中间件。
// 未登录时不调用 Next，而是把错误交给上层统一响应中间件格式化，
// 这样返回结构与业务接口完全一致（code/message/data）。
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

	r.SetCtxVar(CtxUserID, userID)
	r.Middleware.Next()
}
