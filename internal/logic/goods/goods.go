// Package goods 实现商品模块的业务逻辑：分页列表与详情查询。
package goods

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"hello/internal/dao"
	"hello/internal/model/entity"
	"hello/internal/service"
)

// goodsStatusOnSale 上架状态
const goodsStatusOnSale = 1

type sGoods struct{}

func init() {
	service.RegisterGoods(&sGoods{})
}

// List 分页查询上架商品。用户端只能看到 status=1 的商品。
func (s *sGoods) List(ctx context.Context, in service.GoodsListInput) ([]entity.Goods, int, error) {
	m := dao.Goods.Ctx(ctx).Where(dao.Goods.Columns().Status, goodsStatusOnSale)

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
		Where(dao.Goods.Columns().Status, goodsStatusOnSale).
		Scan(&goods); err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "商品不存在或已下架")
	}
	return goods, nil
}
