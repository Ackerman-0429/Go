package service

import (
	"context"

	"hello/internal/model/entity"
)

// GoodsListInput 商品列表查询条件
type GoodsListInput struct {
	Page       int
	PageSize   int
	CategoryID uint64
	Keyword    string
}

// IGoods 商品业务接口
type IGoods interface {
	// List 分页查询上架商品
	List(ctx context.Context, in GoodsListInput) (list []entity.Goods, total int, err error)
	// GetById 查询上架商品详情
	GetById(ctx context.Context, id uint64) (goods *entity.Goods, err error)
}

var localGoods IGoods

// Goods 获取商品业务实例
func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

// RegisterGoods 由 logic 层在 init 中调用
func RegisterGoods(i IGoods) {
	localGoods = i
}
