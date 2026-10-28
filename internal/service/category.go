// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hello/internal/model/entity"
)

type (
	ICategory interface {
		// List 查询全部启用分类，按「顶级 -> 其子级」的树形顺序返回。
		List(ctx context.Context) ([]entity.Category, error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
