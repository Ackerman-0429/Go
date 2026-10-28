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
		// AdminList 后台商品列表。与用户端不同，这里不限制上架状态，可筛选出下架商品。
		AdminList(ctx context.Context, in model.GoodsAdminListInput) ([]entity.Goods, int, error)
		// AdminSave 新建或编辑商品。ID 为 0 走新建，否则按主键更新。
		AdminSave(ctx context.Context, in model.GoodsSaveInput) (uint64, error)
		// AdminSetStatus 商品上下架
		AdminSetStatus(ctx context.Context, id uint64, status uint) error
		// AdminDelete 删除商品。
		// 同时清理购物车里的引用，否则用户购物车联查会出现空商品；
		// 历史订单明细存的是快照，不受影响，所以可以物理删除。
		AdminDelete(ctx context.Context, id uint64) error
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
