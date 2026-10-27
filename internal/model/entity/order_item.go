// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderItem is the golang structure for table order_item.
type OrderItem struct {
	Id         uint64      `json:"id"         orm:"id"          description:"明细ID"`         // 明细ID
	OrderId    uint64      `json:"orderId"    orm:"order_id"    description:"订单ID"`         // 订单ID
	GoodsId    uint64      `json:"goodsId"    orm:"goods_id"    description:"商品ID"`         // 商品ID
	GoodsName  string      `json:"goodsName"  orm:"goods_name"  description:"下单时商品名快照"`     // 下单时商品名快照
	GoodsCover string      `json:"goodsCover" orm:"goods_cover" description:"下单时封面图快照"`     // 下单时封面图快照
	GoodsPrice float64     `json:"goodsPrice" orm:"goods_price" description:"下单时单价快照"`      // 下单时单价快照
	Quantity   uint        `json:"quantity"   orm:"quantity"    description:"购买数量"`         // 购买数量
	Subtotal   float64     `json:"subtotal"   orm:"subtotal"    description:"小计 = 单价 × 数量"` // 小计 = 单价 × 数量
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`             //
}
