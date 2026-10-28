// Package category 实现分类模块的业务逻辑。
package category

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
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

// AdminList 后台分类列表：包含停用分类，按「顶级 → 子级」排序，
// 并补上父分类名称、商品数与子分类数，供后台表格直接展示。
func (s *sCategory) AdminList(ctx context.Context) ([]model.CategoryAdminItem, error) {
	cols := dao.Category.Columns()

	var all []entity.Category
	if err := dao.Category.Ctx(ctx).
		OrderAsc(cols.Sort).
		OrderAsc(cols.Id).
		Scan(&all); err != nil {
		return nil, err
	}
	// 复用前台那套深度优先排序，保证后台看到的顺序和商城一致
	ordered := sortAsTree(all)

	counts, err := goodsCountByCategory(ctx)
	if err != nil {
		return nil, err
	}

	names := make(map[uint]string, len(all))
	for _, c := range all {
		names[c.Id] = c.Name
	}

	// 顶级分类的商品数累加子分类，否则父分类永远显示 0 件，容易被误判为统计错误
	totals := make(map[uint]int, len(all))
	children := make(map[uint]int, len(all))
	for _, c := range all {
		totals[c.Id] += counts[c.Id]
		if c.ParentId != 0 {
			children[c.ParentId]++
			totals[c.ParentId] += counts[c.Id]
		}
	}

	out := make([]model.CategoryAdminItem, 0, len(ordered))
	for _, c := range ordered {
		item := model.CategoryAdminItem{
			ID:         c.Id,
			ParentID:   c.ParentId,
			ParentName: names[c.ParentId],
			Name:       c.Name,
			Sort:       c.Sort,
			Status:     c.Status,
			GoodsCount: totals[c.Id],
			Children:   children[c.Id],
			CreatedAt:  c.CreatedAt,
		}
		out = append(out, item)
	}
	return out, nil
}

// goodsCountByCategory 一次分组查询拿到各分类的商品数，避免逐个分类 Count。
func goodsCountByCategory(ctx context.Context) (map[uint]int, error) {
	var rows []struct {
		CategoryId uint `orm:"category_id"`
		Total      int  `orm:"total"`
	}
	cols := dao.Goods.Columns()
	if err := dao.Goods.Ctx(ctx).
		Fields(cols.CategoryId, "COUNT(*) AS total").
		Group(cols.CategoryId).
		Scan(&rows); err != nil {
		return nil, err
	}

	out := make(map[uint]int, len(rows))
	for _, r := range rows {
		out[r.CategoryId] = r.Total
	}
	return out, nil
}

// AdminSave 新建或编辑分类。仅支持两级，父分类必须是顶级分类。
func (s *sCategory) AdminSave(ctx context.Context, in model.CategorySaveInput) (uint64, error) {
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return 0, gerror.NewCode(gcode.CodeInvalidParameter, "分类名称不能为空")
	}

	cols := dao.Category.Columns()
	parentID := in.ParentID
	if parentID != 0 {
		if in.ID != 0 && in.ID == parentID {
			return 0, gerror.NewCode(gcode.CodeInvalidParameter, "不能把分类挂到自己下面")
		}

		var parent *entity.Category
		if err := dao.Category.Ctx(ctx).WherePri(parentID).Scan(&parent); err != nil {
			return 0, err
		}
		if parent == nil {
			return 0, gerror.NewCode(gcode.CodeNotFound, "父分类不存在")
		}
		if parent.ParentId != 0 {
			return 0, gerror.NewCode(gcode.CodeInvalidParameter, "分类最多两级，不能挂在子分类下面")
		}

		// 自身已经有子分类时不能再变成别人的子分类，否则会出现第三级
		if in.ID != 0 {
			n, err := dao.Category.Ctx(ctx).Where(cols.ParentId, in.ID).Count()
			if err != nil {
				return 0, err
			}
			if n > 0 {
				return 0, gerror.NewCode(gcode.CodeInvalidParameter, "该分类下还有子分类，不能改成子分类")
			}
		}
	}

	// 同一父级下不允许重名，避免前台出现两个一模一样的筛选项
	dup, err := dao.Category.Ctx(ctx).
		Where(cols.ParentId, parentID).
		Where(cols.Name, name).
		WhereNot(cols.Id, in.ID).
		Count()
	if err != nil {
		return 0, err
	}
	if dup > 0 {
		return 0, gerror.NewCode(gcode.CodeInvalidParameter, "同级下已存在同名分类")
	}

	data := g.Map{
		cols.ParentId: parentID,
		cols.Name:     name,
		cols.Sort:     in.Sort,
		cols.Status:   in.Status,
	}

	if in.ID == 0 {
		id, err := dao.Category.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return 0, err
		}
		return uint64(id), nil
	}

	// 不能靠 RowsAffected 判断存在性：提交值与库中完全一致时 MySQL 返回 0
	count, err := dao.Category.Ctx(ctx).WherePri(in.ID).Count()
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, gerror.NewCode(gcode.CodeNotFound, "分类不存在")
	}
	if _, err = dao.Category.Ctx(ctx).WherePri(in.ID).Data(data).Update(); err != nil {
		return 0, err
	}
	return in.ID, nil
}

// AdminDelete 删除分类。有子分类或仍挂着商品时拒绝删除，避免出现孤儿数据。
func (s *sCategory) AdminDelete(ctx context.Context, id uint64) error {
	cols := dao.Category.Columns()

	count, err := dao.Category.Ctx(ctx).WherePri(id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "分类不存在")
	}

	childCount, err := dao.Category.Ctx(ctx).Where(cols.ParentId, id).Count()
	if err != nil {
		return err
	}
	if childCount > 0 {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "该分类下还有 %d 个子分类，请先删除子分类", childCount)
	}

	goodsCount, err := dao.Goods.Ctx(ctx).Where(dao.Goods.Columns().CategoryId, id).Count()
	if err != nil {
		return err
	}
	if goodsCount > 0 {
		return gerror.NewCodef(gcode.CodeInvalidParameter, "该分类下还有 %d 件商品，请先把商品移到其他分类", goodsCount)
	}

	_, err = dao.Category.Ctx(ctx).WherePri(id).Delete()
	return err
}
