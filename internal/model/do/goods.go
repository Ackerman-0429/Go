// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta     `orm:"table:goods, do:true"`
	Id         any         // 商品ID
	CategoryId any         // 所属分类ID
	Name       any         // 商品名称
	Subtitle   any         // 副标题/卖点
	Price      any         // 售价，必须用DECIMAL
	Stock      any         // 库存
	Cover      any         // 封面图URL
	Images     any         // 多图URL，JSON数组
	Detail     any         // 商品详情
	Sales      any         // 销量
	Status     any         // 状态：1上架 0下架
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
