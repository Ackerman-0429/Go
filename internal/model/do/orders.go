// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta          `orm:"table:orders, do:true"`
	Id              any         // 订单ID
	OrderNo         any         // 订单号，对外展示
	UserId          any         // 下单用户ID
	TotalAmount     any         // 订单总金额
	Status          any         // 状态：0待付款 1已付款 2已发货 3已完成 4已取消
	ReceiverName    any         // 收货人
	ReceiverPhone   any         // 收货电话
	ReceiverAddress any         // 收货地址
	Remark          any         // 订单备注
	PaidAt          *gtime.Time // 付款时间
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
