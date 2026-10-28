// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
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
		// AdminList 后台用户列表。支持关键字（用户名/昵称/邮箱/手机号）、角色、状态筛选，
		// 并附带订单数与累计消费，方便判断这个账号值不值得留着。
		AdminList(ctx context.Context, in model.UserAdminListInput) ([]model.UserAdminItem, int, error)
		// AdminSave 编辑用户：昵称、角色、状态。
		// operatorID 是当前登录的管理员，用来拦住「把自己锁在门外」的操作。
		AdminSave(ctx context.Context, in model.UserSaveInput, operatorID uint64) (uint64, error)
		// AdminCreate 后台直接新建用户。复用注册那套校验，额外放开角色与状态，
		// 让管理员可以一次性建出「已启用的管理员」或「先建好但禁用」的账号。
		AdminCreate(ctx context.Context, in model.UserCreateInput) (uint64, error)
		// AdminResetPassword 管理员重置密码，不需要知道原密码。
		AdminResetPassword(ctx context.Context, id uint64, password string) error
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
