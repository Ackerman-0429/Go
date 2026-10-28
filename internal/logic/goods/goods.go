// Package goods 实现商品模块的业务逻辑：分页列表与详情查询。
package goods

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(&sGoods{})
}

// List 分页查询上架商品。用户端只能看到 status=1 的商品。
func (s *sGoods) List(ctx context.Context, in model.GoodsListInput) ([]entity.Goods, int, error) {
	m := dao.Goods.Ctx(ctx).Where(dao.Goods.Columns().Status, consts.GoodsStatusOnSale)

	if in.CategoryID > 0 {
		m = m.Where(dao.Goods.Columns().CategoryId, in.CategoryID)
	}
	if in.Keyword != "" {
		m = m.WhereLike(dao.Goods.Columns().Name, "%"+in.Keyword+"%")
	}

	var (
		list  []entity.Goods
		total int
	)
	if err := m.Page(in.Page, in.PageSize).
		OrderDesc(dao.Goods.Columns().Sales).
		OrderDesc(dao.Goods.Columns().Id).
		ScanAndCount(&list, &total, false); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// GetById 查询上架商品详情
func (s *sGoods) GetById(ctx context.Context, id uint64) (*entity.Goods, error) {
	var goods *entity.Goods
	if err := dao.Goods.Ctx(ctx).
		Where(dao.Goods.Columns().Id, id).
		Where(dao.Goods.Columns().Status, consts.GoodsStatusOnSale).
		Scan(&goods); err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "商品不存在或已下架")
	}
	return goods, nil
}

// AdminList 后台商品列表。与用户端不同，这里不限制上架状态，可筛选出下架商品。
func (s *sGoods) AdminList(ctx context.Context, in model.GoodsAdminListInput) ([]entity.Goods, int, error) {
	cols := dao.Goods.Columns()
	m := dao.Goods.Ctx(ctx)

	if in.Status >= 0 {
		m = m.Where(cols.Status, in.Status)
	}
	if in.CategoryID > 0 {
		m = m.Where(cols.CategoryId, in.CategoryID)
	}
	if in.Keyword != "" {
		m = m.WhereLike(cols.Name, "%"+in.Keyword+"%")
	}

	var (
		list  []entity.Goods
		total int
	)
	// 后台按 ID 倒序：最近录入的排前面，便于核对
	if err := m.Page(in.Page, in.PageSize).
		OrderDesc(cols.Id).
		ScanAndCount(&list, &total, false); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// AdminSave 新建或编辑商品。ID 为 0 走新建，否则按主键更新。
func (s *sGoods) AdminSave(ctx context.Context, in model.GoodsSaveInput) (uint64, error) {
	cols := dao.Goods.Columns()
	data := g.Map{
		cols.CategoryId: in.CategoryID,
		cols.Name:       in.Name,
		cols.Subtitle:   in.Subtitle,
		cols.Price:      in.Price,
		cols.Stock:      in.Stock,
		cols.Cover:      in.Cover,
		cols.Images:     in.Images,
		cols.Detail:     in.Detail,
		cols.Status:     in.Status,
	}

	if in.ID == 0 {
		id, err := dao.Goods.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return 0, err
		}
		return uint64(id), nil
	}

	// 先确认记录存在。MySQL 在「提交的值与库里完全一致」时受影响行数为 0，
	// 不能拿 RowsAffected 判断商品是否存在，否则原样保存会被误报为不存在。
	count, err := dao.Goods.Ctx(ctx).WherePri(in.ID).Count()
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	if _, err = dao.Goods.Ctx(ctx).WherePri(in.ID).Data(data).Update(); err != nil {
		return 0, err
	}
	return in.ID, nil
}

// AdminSetStatus 商品上下架
func (s *sGoods) AdminSetStatus(ctx context.Context, id uint64, status uint) error {
	count, err := dao.Goods.Ctx(ctx).WherePri(id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "商品不存在")
	}
	_, err = dao.Goods.Ctx(ctx).
		WherePri(id).
		Data(g.Map{dao.Goods.Columns().Status: status}).
		Update()
	return err
}

// AdminDelete 删除商品。
// 同时清理购物车里的引用，否则用户购物车联查会出现空商品；
// 历史订单明细存的是快照，不受影响，所以可以物理删除。
func (s *sGoods) AdminDelete(ctx context.Context, id uint64) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		count, err := dao.Goods.Ctx(ctx).WherePri(id).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			return gerror.NewCode(gcode.CodeNotFound, "商品不存在")
		}

		if _, err = dao.Cart.Ctx(ctx).
			Where(dao.Cart.Columns().GoodsId, id).
			Delete(); err != nil {
			return err
		}
		_, err = dao.Goods.Ctx(ctx).WherePri(id).Delete()
		return err
	})
}
