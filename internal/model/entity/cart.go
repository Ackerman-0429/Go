// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cart is the golang structure for table cart.
type Cart struct {
	Id        uint64      `json:"id"        orm:"id"         description:"购物车项ID"` // 购物车项ID
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"用户ID"`   // 用户ID
	GoodsId   uint64      `json:"goodsId"   orm:"goods_id"   description:"商品ID"`   // 商品ID
	Quantity  uint        `json:"quantity"  orm:"quantity"   description:"数量"`     // 数量
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`       //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`       //
}
