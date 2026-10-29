// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsReview is the golang structure of table goods_review for DAO operations like Where/Data.
type GoodsReview struct {
	g.Meta      `orm:"table:goods_review, do:true"`
	Id          any         // 评价ID
	GoodsId     any         // 商品ID
	OrderId     any         // 订单ID
	OrderItemId any         // 订单明细ID，一条明细只能评一次
	UserId      any         // 评价用户ID
	Rating      any         // 评分 1-5 星
	Content     any         // 评价内容
	Anonymous   any         // 是否匿名：1 是 0 否
	Status      any         // 状态：1 显示 0 隐藏
	CreatedAt   *gtime.Time //
}
