package service

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"hello/internal/model/entity"
)

// 订单状态
const (
	OrderStatusPending   uint = 0 // 待付款
	OrderStatusPaid      uint = 1 // 已付款
	OrderStatusShipped   uint = 2 // 已发货
	OrderStatusFinished  uint = 3 // 已完成
	OrderStatusCancelled uint = 4 // 已取消
)

// OrderCreateInput 下单参数，商品来自当前用户购物车中可结算的部分
type OrderCreateInput struct {
	UserID          uint64
	ReceiverName    string
	ReceiverPhone   string
	ReceiverAddress string
	Remark          string
}

// OrderCreateOutput 下单结果
type OrderCreateOutput struct {
	OrderID     uint64
	OrderNo     string
	TotalAmount float64
	Quantity    uint
}

// OrderListInput 订单列表查询条件，Status 传 -1 表示全部
type OrderListInput struct {
	UserID   uint64
	Page     int
	PageSize int
	Status   int
}

// OrderListItem 订单列表项，商品信息由 order_item 聚合而来
type OrderListItem struct {
	ID          uint64
	OrderNo     string
	TotalAmount float64
	Status      uint
	Quantity    uint
	ItemCount   int
	FirstName   string
	FirstCover  string
	CreatedAt   *gtime.Time
}

// OrderDetail 订单详情
type OrderDetail struct {
	Order *entity.Orders
	Items []entity.OrderItem
}

// IOrder 订单业务接口
type IOrder interface {
	// Create 把购物车中可结算的商品生成订单，扣减库存并清空相应购物车项
	Create(ctx context.Context, in OrderCreateInput) (out *OrderCreateOutput, err error)
	// List 分页查询当前用户的订单
	List(ctx context.Context, in OrderListInput) (list []OrderListItem, total int, err error)
	// Detail 查询订单详情，含明细
	Detail(ctx context.Context, userID, orderID uint64) (out *OrderDetail, err error)
	// Cancel 取消待付款订单并回滚库存
	Cancel(ctx context.Context, userID, orderID uint64) (err error)
	// Pay 模拟支付，把待付款改为已付款并累计销量
	Pay(ctx context.Context, userID, orderID uint64) (paidAt *gtime.Time, err error)
}

var localOrder IOrder

// Order 获取订单业务实例
func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

// RegisterOrder 由 logic 层在 init 中调用
func RegisterOrder(i IOrder) {
	localOrder = i
}
