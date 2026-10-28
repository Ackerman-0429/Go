// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model/entity"
)

type (
	IUser interface {
		// Register 注册新用户。密码用 bcrypt 哈希后入库，绝不存明文。
		Register(ctx context.Context, username string, password string, nickname string) (uint64, error)
		// Login 登录校验。用户名不存在和密码错误统一提示，避免暴露账号是否存在。
		Login(ctx context.Context, username string, password string) (*entity.User, error)
		// GetById 按主键查询用户
		GetById(ctx context.Context, id uint64) (*entity.User, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
