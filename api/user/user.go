// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"hello/api/user/v1"
)

type IUserV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
