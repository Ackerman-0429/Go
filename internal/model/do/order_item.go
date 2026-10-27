// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderItem is the golang structure of table order_item for DAO operations like Where/Data.
type OrderItem struct {
	g.Meta     `orm:"table:order_item, do:true"`
	Id         any         // 明细ID
	OrderId    any         // 订单ID
	GoodsId    any         // 商品ID
	GoodsName  any         // 下单时商品名快照
	GoodsCover any         // 下单时封面图快照
	GoodsPrice any         // 下单时单价快照
	Quantity   any         // 购买数量
	Subtotal   any         // 小计 = 单价 × 数量
	CreatedAt  *gtime.Time //
}
