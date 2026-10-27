package user

import (
	"context"

	v1 "hello/api/user/v1"
)

// IUserV1 用户模块接口，由 controller 实现
type IUserV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
}
