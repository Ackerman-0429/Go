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
}

// LogoutReq 退出登录
type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" summary:"退出登录" tags:"用户"`
}

// LogoutRes 退出结果
type LogoutRes struct{}
