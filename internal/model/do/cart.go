// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cart is the golang structure of table cart for DAO operations like Where/Data.
type Cart struct {
	g.Meta    `orm:"table:cart, do:true"`
	Id        any         // 购物车项ID
	UserId    any         // 用户ID
	GoodsId   any         // 商品ID
	Quantity  any         // 数量
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
