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
	ICategory interface {
		// List 查询全部启用分类，按「顶级 -> 其子级」的树形顺序返回。
		List(ctx context.Context) ([]entity.Category, error)
		// AdminList 后台分类列表：包含停用分类，按「顶级 → 子级」排序，
		// 并补上父分类名称、商品数与子分类数，供后台表格直接展示。
		AdminList(ctx context.Context) ([]model.CategoryAdminItem, error)
		// AdminSave 新建或编辑分类。仅支持两级，父分类必须是顶级分类。
		AdminSave(ctx context.Context, in model.CategorySaveInput) (uint64, error)
		// AdminDelete 删除分类。有子分类或仍挂着商品时拒绝删除，避免出现孤儿数据。
		AdminDelete(ctx context.Context, id uint64) error
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
