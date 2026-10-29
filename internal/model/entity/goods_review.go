// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsReview is the golang structure for table goods_review.
type GoodsReview struct {
	Id          uint64      `json:"id"          orm:"id"            description:"评价ID"`             // 评价ID
	GoodsId     uint64      `json:"goodsId"     orm:"goods_id"      description:"商品ID"`             // 商品ID
	OrderId     uint64      `json:"orderId"     orm:"order_id"      description:"订单ID"`             // 订单ID
	OrderItemId uint64      `json:"orderItemId" orm:"order_item_id" description:"订单明细ID，一条明细只能评一次"` // 订单明细ID，一条明细只能评一次
	UserId      uint64      `json:"userId"      orm:"user_id"       description:"评价用户ID"`           // 评价用户ID
	Rating      uint        `json:"rating"      orm:"rating"        description:"评分 1-5 星"`         // 评分 1-5 星
	Content     string      `json:"content"     orm:"content"       description:"评价内容"`             // 评价内容
	Anonymous   uint        `json:"anonymous"   orm:"anonymous"     description:"是否匿名：1 是 0 否"`     // 是否匿名：1 是 0 否
	Status      uint        `json:"status"      orm:"status"        description:"状态：1 显示 0 隐藏"`     // 状态：1 显示 0 隐藏
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`                 //
}
