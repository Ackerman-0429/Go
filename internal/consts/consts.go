// Package consts 定义项目级常量与枚举，避免魔法值散落在各层。
package consts

// 用户角色。与 user.role 列取值一一对应。
const (
	UserRoleNormal uint = 0 // 普通用户
	UserRoleAdmin  uint = 1 // 管理员
)

// 用户状态。与 user.status 列取值一一对应。
const (
	UserStatusDisabled uint = 0 // 禁用
	UserStatusNormal   uint = 1 // 正常
)

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

// OrderStatusText 订单状态中文文案，前后台共用一份，避免两处 switch 各写一遍。
func OrderStatusText(status uint) string {
	switch status {
	case OrderStatusPending:
		return "待付款"
	case OrderStatusPaid:
		return "已付款"
	case OrderStatusShipped:
		return "已发货"
	case OrderStatusFinished:
		return "已完成"
	case OrderStatusCancelled:
		return "已取消"
	}
	return "未知状态"
}
