// Package consts 定义项目级常量与枚举，避免魔法值散落在各层。
package consts

// 商品上下架状态。与 goods.status 列取值一一对应。
const (
	GoodsStatusOffSale uint = 0 // 下架
	GoodsStatusOnSale  uint = 1 // 上架
)

// 订单状态。与 orders.status 列取值一一对应。
const (
	OrderStatusPending   uint = 0 // 待付款
	OrderStatusPaid      uint = 1 // 已付款
	OrderStatusShipped   uint = 2 // 已发货
	OrderStatusFinished  uint = 3 // 已完成
	OrderStatusCancelled uint = 4 // 已取消
)
