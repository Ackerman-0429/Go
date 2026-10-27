// Package service 定义业务接口与注册机制，具体实现位于 internal/logic。
package service

import (
	"context"

	"hello/internal/model/entity"
)

// IUser 用户业务接口。
// 定义在 service 层、实现在 logic 层，controller 只依赖这个接口。
type IUser interface {
	// Register 注册，返回新用户ID
	Register(ctx context.Context, username, password, nickname string) (id uint64, err error)
	// Login 登录，校验通过返回用户实体
	Login(ctx context.Context, username, password string) (user *entity.User, err error)
	// GetById 按ID查询用户
	GetById(ctx context.Context, id uint64) (user *entity.User, err error)
}

var localUser IUser

// User 获取用户业务实例
func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// RegisterUser 由 logic 层在 init 中调用，完成实现注册
func RegisterUser(i IUser) {
	localUser = i
}
