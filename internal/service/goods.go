// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
	"goshop/internal/model/entity"
)

type (
	IGoods interface {
		// List 分页查询上架商品。用户端只能看到 status=1 的商品。
		List(ctx context.Context, in model.GoodsListInput) ([]entity.Goods, int, error)
		// GetById 查询上架商品详情
		GetById(ctx context.Context, id uint64) (*entity.Goods, error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
