// Package user 实现用户模块的控制器，只负责参数接收与响应组装。
package user

import (
	v1api "hello/api/user"
)

// ControllerV1 用户模块控制器
type ControllerV1 struct{}

// NewV1 由 cmd 层调用，完成接口注入
func NewV1() v1api.IUserV1 {
	return &ControllerV1{}
}
