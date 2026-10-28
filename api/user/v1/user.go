// Package v1 定义用户模块 v1 版本的接口出入参结构。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 注册请求
type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" summary:"用户注册" tags:"用户"`
	Username string `json:"username" v:"required|length:3,20#请输入用户名|用户名长度需在3-20之间" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,32#请输入密码|密码长度需在6-32之间" dc:"密码"`
	Nickname string `json:"nickname" v:"max-length:20#昵称不能超过20个字符" dc:"昵称"`
}

// RegisterRes 注册响应
type RegisterRes struct {
	ID uint64 `json:"id" dc:"用户ID"`
}

// LoginReq 登录请求
type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" summary:"用户登录" tags:"用户"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

// LoginRes 登录响应
type LoginRes struct {
	ID       uint64 `json:"id"       dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Role     uint   `json:"role"     dc:"角色：0普通用户 1管理员"`
}

// InfoReq 当前登录用户信息
type InfoReq struct {
	g.Meta `path:"/user/info" method:"get" summary:"获取当前登录用户" tags:"用户"`
}

// InfoRes 用户信息
type InfoRes struct {
	ID       uint64 `json:"id"       dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Email    string `json:"email"    dc:"邮箱"`
	Phone    string `json:"phone"    dc:"手机号"`
	Role     uint   `json:"role"     dc:"角色：0普通用户 1管理员"`
}

// ProfileReq 修改自己的资料。只能改昵称、邮箱、手机号：
// 用户名是登录身份、角色与状态是管理员的权限，这些都不开放给用户自助修改。
type ProfileReq struct {
	g.Meta   `path:"/user/profile" method:"post" summary:"修改个人资料" tags:"用户"`
	Nickname string `json:"nickname" v:"max-length:20#昵称不能超过20个字符"  dc:"昵称，留空则沿用用户名"`
	Email    string `json:"email"    v:"email#邮箱格式不正确"           dc:"邮箱，留空表示不填"`
	Phone    string `json:"phone"    v:"phone-loose#手机号格式不正确"    dc:"手机号，留空表示不填"`
}

// ProfileRes 修改后的资料，前端可直接用来刷新页面顶部的昵称
type ProfileRes struct {
	ID       uint64 `json:"id"       dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Email    string `json:"email"    dc:"邮箱"`
	Phone    string `json:"phone"    dc:"手机号"`
	Role     uint   `json:"role"     dc:"角色：0普通用户 1管理员"`
}

// PasswordReq 修改自己的密码。必须验原密码：只会话还不够，
// 否则别人拿到没锁屏的浏览器就能直接把账号密码换掉。
type PasswordReq struct {
	g.Meta      `path:"/user/password" method:"post" summary:"修改密码" tags:"用户"`
	OldPassword string `json:"oldPassword" v:"required#请输入原密码"                          dc:"原密码"`
	NewPassword string `json:"newPassword" v:"required|length:6,32#请输入新密码|密码长度需在6-32之间" dc:"新密码"`
}

// PasswordRes 修改结果。改完不踢当前会话，用户不用重新登录。
type PasswordRes struct{}

// LogoutReq 退出登录
type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" summary:"退出登录" tags:"用户"`
}

// LogoutRes 退出结果
type LogoutRes struct{}
