// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hello/internal/model"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IOrder interface {
		// Create 把购物车中可结算的商品生成订单。
		// 扣库存、写订单、写明细、清购物车在同一个事务里，任何一步失败整体回滚。
		Create(ctx context.Context, in model.OrderCreateInput) (*model.OrderCreateOutput, error)
		// List 分页查询当前用户的订单，商品信息由明细聚合
		List(ctx context.Context, in model.OrderListInput) ([]model.OrderListItem, int, error)
		// Detail 查询订单详情，订单与明细都限定在当前用户范围内
		Detail(ctx context.Context, userID uint64, orderID uint64) (*model.OrderDetail, error)
		// Cancel 取消待付款订单并回滚库存
		Cancel(ctx context.Context, userID uint64, orderID uint64) error
		// Pay 模拟支付：待付款改为已付款，并把明细数量累加进商品销量。
		// 真实项目这里要接支付网关与异步回调，此处仅演示状态流转。
		Pay(ctx context.Context, userID uint64, orderID uint64) (*gtime.Time, error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
