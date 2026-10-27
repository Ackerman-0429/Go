// Package category 实现分类模块的业务逻辑。
package category

import (
	"context"

	"hello/internal/dao"
	"hello/internal/model/entity"
	"hello/internal/service"
)

// categoryStatusEnabled 启用状态
const categoryStatusEnabled = 1

type sCategory struct{}

func init() {
	service.RegisterCategory(&sCategory{})
}

// List 查询全部启用分类，按「顶级 -> 其子级」的树形顺序返回。
func (s *sCategory) List(ctx context.Context) ([]entity.Category, error) {
	var all []entity.Category
	err := dao.Category.Ctx(ctx).
		Where(dao.Category.Columns().Status, categoryStatusEnabled).
		OrderAsc(dao.Category.Columns().Sort).
		OrderAsc(dao.Category.Columns().Id).
		Scan(&all)
	if err != nil {
		return nil, err
	}
	return sortAsTree(all), nil
}

// sortAsTree 把平铺列表按父子关系重排。
// 仅按 sort 排序会让子分类插到别的顶级分类之间，这里改成深度优先。
func sortAsTree(list []entity.Category) []entity.Category {
	children := make(map[uint][]entity.Category, len(list))
	for _, item := range list {
		children[item.ParentId] = append(children[item.ParentId], item)
	}

	ordered := make([]entity.Category, 0, len(list))
	var walk func(parentID uint)
	walk = func(parentID uint) {
		for _, item := range children[parentID] {
			ordered = append(ordered, item)
			walk(item.Id)
		}
	}
	walk(0)

	// 兜底：若存在 parent 已被删除成为孤儿的数据，追加到末尾避免丢失
	if len(ordered) < len(list) {
		seen := make(map[uint]bool, len(ordered))
		for _, item := range ordered {
			seen[item.Id] = true
		}
		for _, item := range list {
			if !seen[item.Id] {
				ordered = append(ordered, item)
			}
		}
	}
	return ordered
}
