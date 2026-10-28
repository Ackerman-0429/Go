// Package v1 定义购物车模块 v1 版本的接口出入参结构。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AddReq 加入购物车请求
type AddReq struct {
	g.Meta   `path:"/cart/add" method:"post" summary:"加入购物车" tags:"购物车"`
	GoodsID  uint64 `json:"goodsId"  v:"required|min:1#请选择商品|商品ID不合法"        dc:"商品ID"`
	Quantity uint   `json:"quantity" d:"1" v:"between:1,999#数量需在1-999之间" dc:"购买数量，默认1"`
}

// AddRes 加购结果
type AddRes struct {
	ID       uint64 `json:"id"       dc:"购物车项ID"`
	Quantity uint   `json:"quantity" dc:"加入后该商品在购物车中的总数量"`
}

// ListReq 购物车列表请求
type ListReq struct {
	g.Meta `path:"/cart/list" method:"get" summary:"购物车列表" tags:"购物车"`
}

// ListItem 购物车项，带商品快照与可购买状态
type ListItem struct {
	ID        uint64  `json:"id"        dc:"购物车项ID"`
	GoodsID   uint64  `json:"goodsId"   dc:"商品ID"`
	Name      string  `json:"name"      dc:"商品名称"`
	Subtitle  string  `json:"subtitle"  dc:"副标题"`
	Cover     string  `json:"cover"     dc:"封面图"`
	Price     float64 `json:"price"     dc:"单价"`
	Quantity  uint    `json:"quantity"  dc:"数量"`
	Stock     uint    `json:"stock"     dc:"当前库存"`
	Subtotal  float64 `json:"subtotal"  dc:"小计"`
	Available bool    `json:"available" dc:"是否可结算（下架、售罄或超过库存时为 false）"`
	Reason    string  `json:"reason"    dc:"不可结算的原因"`
}

// ListRes 购物车汇总。总件数与总金额只统计可结算的商品。
type ListRes struct {
	List          []ListItem `json:"list"          dc:"购物车列表"`
	TotalQuantity uint       `json:"totalQuantity" dc:"可结算商品总件数"`
	TotalAmount   float64    `json:"totalAmount"   dc:"可结算商品总金额"`
	InvalidCount  int        `json:"invalidCount"  dc:"不可结算的商品条数"`
}

// UpdateReq 修改数量请求
type UpdateReq struct {
	g.Meta   `path:"/cart/update" method:"post" summary:"修改购物车数量" tags:"购物车"`
	ID       uint64 `json:"id"       v:"required|min:1#请提供购物车项ID|购物车项ID不合法" dc:"购物车项ID"`
	Quantity uint   `json:"quantity" v:"between:1,999#数量需在1-999之间"              dc:"修改后的数量"`
}

// UpdateRes 修改结果
type UpdateRes struct {
	Quantity uint `json:"quantity" dc:"修改后的数量"`
}

// DeleteReq 删除购物车项请求
type DeleteReq struct {
	g.Meta `path:"/cart/delete" method:"post" summary:"删除购物车项" tags:"购物车"`
	ID     uint64 `json:"id" v:"required|min:1#请提供购物车项ID|购物车项ID不合法" dc:"购物车项ID"`
}

// DeleteRes 删除结果
type DeleteRes struct{}

// ClearReq 清空购物车请求
type ClearReq struct {
	g.Meta `path:"/cart/clear" method:"post" summary:"清空购物车" tags:"购物车"`
}

// ClearRes 清空结果
type ClearRes struct {
	Deleted int `json:"deleted" dc:"删除的条数"`
}
