// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure for table goods.
type Goods struct {
	Id         uint64      `json:"id"         orm:"id"          description:"商品ID"`          // 商品ID
	CategoryId uint        `json:"categoryId" orm:"category_id" description:"所属分类ID"`        // 所属分类ID
	Name       string      `json:"name"       orm:"name"        description:"商品名称"`          // 商品名称
	Subtitle   string      `json:"subtitle"   orm:"subtitle"    description:"副标题/卖点"`        // 副标题/卖点
	Price      float64     `json:"price"      orm:"price"       description:"售价，必须用DECIMAL"` // 售价，必须用DECIMAL
	Stock      uint        `json:"stock"      orm:"stock"       description:"库存"`            // 库存
	Cover      string      `json:"cover"      orm:"cover"       description:"封面图URL"`        // 封面图URL
	Images     string      `json:"images"     orm:"images"      description:"多图URL，JSON数组"`  // 多图URL，JSON数组
	Detail     string      `json:"detail"     orm:"detail"      description:"商品详情"`          // 商品详情
	Sales      uint        `json:"sales"      orm:"sales"       description:"销量"`            // 销量
	Status     uint        `json:"status"     orm:"status"      description:"状态：1上架 0下架"`    // 状态：1上架 0下架
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`              //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`              //
}
