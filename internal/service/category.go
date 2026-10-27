package service

import (
	"context"

	"hello/internal/model/entity"
)

// ICategory 分类业务接口
type ICategory interface {
	// List 查询全部启用分类
	List(ctx context.Context) (list []entity.Category, err error)
}

var localCategory ICategory

// Category 获取分类业务实例
func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

// RegisterCategory 由 logic 层在 init 中调用
func RegisterCategory(i ICategory) {
	localCategory = i
}
