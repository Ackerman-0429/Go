package service

import (
	"context"
)

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

// ICart 购物车业务接口
type ICart interface {
	// Add 加入购物车，已存在则累加数量，返回购物车项ID与累加后的数量
	Add(ctx context.Context, in CartAddInput) (id uint64, quantity uint, err error)
	// List 查询当前用户的购物车
	List(ctx context.Context, userID uint64) (out *CartListOutput, err error)
	// UpdateQuantity 修改数量，需校验归属与库存
	UpdateQuantity(ctx context.Context, userID, cartID uint64, quantity uint) (err error)
	// Remove 删除购物车项
	Remove(ctx context.Context, userID, cartID uint64) (err error)
	// Clear 清空购物车，返回删除条数
	Clear(ctx context.Context, userID uint64) (deleted int, err error)
}

var localCart ICart

// Cart 获取购物车业务实例
func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

// RegisterCart 由 logic 层在 init 中调用
func RegisterCart(i ICart) {
	localCart = i
}
