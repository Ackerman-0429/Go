// Package user 实现用户模块的业务逻辑：注册、登录、查询。
package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"

	"hello/internal/dao"
	"hello/internal/model/do"
	"hello/internal/model/entity"
	"hello/internal/service"
)

// sUser 是 service.IUser 的实现
type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

// Register 注册新用户。密码用 bcrypt 哈希后入库，绝不存明文。
func (s *sUser) Register(ctx context.Context, username, password, nickname string) (uint64, error) {
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Count()
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, gerror.NewCode(gcode.CodeValidationFailed, "用户名已被占用")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	if nickname == "" {
		nickname = username
	}

	result, err := dao.User.Ctx(ctx).Data(do.User{
		Username: username,
		Password: string(hashed),
		Nickname: nickname,
		Status:   1,
	}).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// Login 登录校验。用户名不存在和密码错误统一提示，避免暴露账号是否存在。
func (s *sUser) Login(ctx context.Context, username, password string) (*entity.User, error) {
	var user *entity.User
	if err := dao.User.Ctx(ctx).Where(dao.User.Columns().Username, username).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "用户名或密码错误")
	}

	if user.Status != 1 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "账号已被禁用")
	}

	return user, nil
}

// GetById 按主键查询用户
func (s *sUser) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(id).Scan(&user); err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	return user, nil
}
