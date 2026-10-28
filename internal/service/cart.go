// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hello/internal/model"
)

type (
	ICart interface {
		// Add 加入购物车。同一商品已存在时累加数量，累加后不得超过库存。
		Add(ctx context.Context, in model.CartAddInput) (uint64, uint, error)
		// List 查询购物车。左联商品表一次取全，避免逐条查询产生 N+1。
		List(ctx context.Context, userID uint64) (*model.CartListOutput, error)
		// UpdateQuantity 修改购物车项数量
		UpdateQuantity(ctx context.Context, userID uint64, cartID uint64, quantity uint) error
		// Remove 删除购物车项
		Remove(ctx context.Context, userID uint64, cartID uint64) error
		// Clear 清空当前用户的购物车
		Clear(ctx context.Context, userID uint64) (int, error)
	}
)

var (
	localCart ICart
)

func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

func RegisterCart(i ICart) {
	localCart = i
}
