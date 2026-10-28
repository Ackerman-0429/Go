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
	OrderShip(ctx context.Context, req *v1.OrderShipReq) (res *v1.OrderShipRes, err error)
	OrderFinish(ctx context.Context, req *v1.OrderFinishReq) (res *v1.OrderFinishRes, err error)
}
