package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"hello/internal/model/entity"
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
