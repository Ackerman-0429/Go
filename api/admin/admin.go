// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"goshop/api/admin/v1"
)

type IAdminV1 interface {
	Stats(ctx context.Context, req *v1.StatsReq) (res *v1.StatsRes, err error)
	GoodsList(ctx context.Context, req *v1.GoodsListReq) (res *v1.GoodsListRes, err error)
	GoodsSave(ctx context.Context, req *v1.GoodsSaveReq) (res *v1.GoodsSaveRes, err error)
	GoodsStatus(ctx context.Context, req *v1.GoodsStatusReq) (res *v1.GoodsStatusRes, err error)
	GoodsDelete(ctx context.Context, req *v1.GoodsDeleteReq) (res *v1.GoodsDeleteRes, err error)
	OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
	OrderShip(ctx context.Context, req *v1.OrderShipReq) (res *v1.OrderShipRes, err error)
	OrderFinish(ctx context.Context, req *v1.OrderFinishReq) (res *v1.OrderFinishRes, err error)
	CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error)
	CategorySave(ctx context.Context, req *v1.CategorySaveReq) (res *v1.CategorySaveRes, err error)
	CategoryDelete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserSave(ctx context.Context, req *v1.UserSaveReq) (res *v1.UserSaveRes, err error)
	UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error)
	UserResetPassword(ctx context.Context, req *v1.UserResetPasswordReq) (res *v1.UserResetPasswordRes, err error)
}
