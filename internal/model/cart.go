// Package model 定义跨层复用的业务数据结构（输入输出对象）。
// 表结构映射见子包 do/entity，由 gf gen dao 自动生成。
package model

// CartAddInput 加购参数
type CartAddInput struct {
	UserID   uint64
	GoodsID  uint64
	Quantity uint
}

// CartItem 购物车项，已联查商品信息
type CartItem struct {
	ID        uint64
	GoodsID   uint64
	Name      string
	Subtitle  string
	Cover     string
	Price     float64
	Quantity  uint
	Stock     uint
	Subtotal  float64
	Available bool
	Reason    string
}

// CartListOutput 购物车汇总，Total 系列只统计可结算的商品
type CartListOutput struct {
	List          []CartItem
	TotalQuantity uint
	TotalAmount   float64
	InvalidCount  int
}
