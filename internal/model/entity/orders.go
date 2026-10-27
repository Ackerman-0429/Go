// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure for table orders.
type Orders struct {
	Id              uint64      `json:"id"              orm:"id"               description:"订单ID"`                        // 订单ID
	OrderNo         string      `json:"orderNo"         orm:"order_no"         description:"订单号，对外展示"`                    // 订单号，对外展示
	UserId          uint64      `json:"userId"          orm:"user_id"          description:"下单用户ID"`                      // 下单用户ID
	TotalAmount     float64     `json:"totalAmount"     orm:"total_amount"     description:"订单总金额"`                       // 订单总金额
	Status          uint        `json:"status"          orm:"status"           description:"状态：0待付款 1已付款 2已发货 3已完成 4已取消"` // 状态：0待付款 1已付款 2已发货 3已完成 4已取消
	ReceiverName    string      `json:"receiverName"    orm:"receiver_name"    description:"收货人"`                         // 收货人
	ReceiverPhone   string      `json:"receiverPhone"   orm:"receiver_phone"   description:"收货电话"`                        // 收货电话
	ReceiverAddress string      `json:"receiverAddress" orm:"receiver_address" description:"收货地址"`                        // 收货地址
	Remark          string      `json:"remark"          orm:"remark"           description:"订单备注"`                        // 订单备注
	PaidAt          *gtime.Time `json:"paidAt"          orm:"paid_at"          description:"付款时间"`                        // 付款时间
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`                            //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`                            //
}
