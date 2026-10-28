// Package cart 实现购物车模块的业务逻辑：加购、列表、改数量、删除、清空。
package cart

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"hello/internal/dao"
	"hello/internal/model/entity"
	"hello/internal/service"
	"hello/utility/money"
)

// goodsStatusOnSale 上架状态，与商品模块保持一致
const goodsStatusOnSale = 1

type sCart struct{}

func init() {
	service.RegisterCart(&sCart{})
}

// cartJoinRow 购物车联查商品的结果行
type cartJoinRow struct {
	CartId      uint64  `orm:"cart_id"`
	GoodsId     uint64  `orm:"goods_id"`
	Quantity    uint    `orm:"quantity"`
	Name        string  `orm:"name"`
	Subtitle    string  `orm:"subtitle"`
	Cover       string  `orm:"cover"`
	Price       float64 `orm:"price"`
	Stock       uint    `orm:"stock"`
	GoodsStatus uint    `orm:"goods_status"`
}

// Add 加入购物车。同一商品已存在时累加数量，累加后不得超过库存。
func (s *sCart) Add(ctx context.Context, in service.CartAddInput) (uint64, uint, error) {
	// 商品必须存在且在售
	goods, err := service.Goods().GetById(ctx, in.GoodsID)
	if err != nil {
		return 0, 0, err
	}
	if goods.Stock == 0 {
		return 0, 0, gerror.NewCode(gcode.CodeNotFound, "商品已售罄")
	}

	cols := dao.Cart.Columns()
	exist, err := s.getByGoods(ctx, in.UserID, in.GoodsID)
	if err != nil {
		return 0, 0, err
	}

	quantity := in.Quantity
	if exist != nil {
		quantity += exist.Quantity
	}
	if quantity > goods.Stock {
		return 0, 0, gerror.Newf("库存不足，最多可购买 %d 件", goods.Stock)
	}

	if exist != nil {
		if _, err = dao.Cart.Ctx(ctx).
			Where(cols.Id, exist.Id).
			Data(g.Map{cols.Quantity: quantity}).
			Update(); err != nil {
			return 0, 0, err
		}
		return exist.Id, quantity, nil
	}

	id, err := dao.Cart.Ctx(ctx).Data(g.Map{
		cols.UserId:   in.UserID,
		cols.GoodsId:  in.GoodsID,
		cols.Quantity: quantity,
	}).InsertAndGetId()
	if err != nil {
		// 唯一键 uk_user_goods 冲突说明并发下刚插入同商品记录，退化为累加更新
		if row, e := s.getByGoods(ctx, in.UserID, in.GoodsID); e == nil && row != nil {
			merged := row.Quantity + in.Quantity
			if merged > goods.Stock {
				return 0, 0, gerror.Newf("库存不足，最多可购买 %d 件", goods.Stock)
			}
			if _, e2 := dao.Cart.Ctx(ctx).
				Where(cols.Id, row.Id).
				Data(g.Map{cols.Quantity: merged}).
				Update(); e2 != nil {
				return 0, 0, e2
			}
			return row.Id, merged, nil
		}
		return 0, 0, err
	}
	return uint64(id), quantity, nil
}

// List 查询购物车。左联商品表一次取全，避免逐条查询产生 N+1。
func (s *sCart) List(ctx context.Context, userID uint64) (*service.CartListOutput, error) {
	var rows []cartJoinRow
	err := dao.Cart.Ctx(ctx).As("c").
		LeftJoin("goods g", "g.id = c.goods_id").
		Fields("c.id AS cart_id, c.goods_id, c.quantity, " +
			"g.name, g.subtitle, g.cover, g.price, g.stock, g.status AS goods_status").
		Where("c.user_id", userID).
		OrderAsc("c.id").
		Scan(&rows)
	if err != nil {
		return nil, err
	}

	out := &service.CartListOutput{List: make([]service.CartItem, 0, len(rows))}
	var totalCents int64
	for _, row := range rows {
		item := service.CartItem{
			ID:        row.CartId,
			GoodsID:   row.GoodsId,
			Name:      row.Name,
			Subtitle:  row.Subtitle,
			Cover:     row.Cover,
			Price:     row.Price,
			Quantity:  row.Quantity,
			Stock:     row.Stock,
			Available: true,
		}

		// 下架、售罄、数量超库存都不能结算，但保留在列表里让用户自己处理
		switch {
		case row.GoodsStatus != goodsStatusOnSale:
			item.Available = false
			item.Reason = "商品已下架"
		case row.Stock == 0:
			item.Available = false
			item.Reason = "商品已售罄"
		case row.Quantity > row.Stock:
			item.Available = false
			item.Reason = "库存不足，请调整数量"
		}

		// 单价先转成分再乘数量，全程整数运算
		cents := money.ToCents(row.Price) * int64(row.Quantity)
		item.Subtotal = money.FromCents(cents)
		if item.Available {
			totalCents += cents
			out.TotalQuantity += row.Quantity
		} else {
			out.InvalidCount++
		}
		out.List = append(out.List, item)
	}

	out.TotalAmount = money.FromCents(totalCents)
	return out, nil
}

// UpdateQuantity 修改购物车项数量
func (s *sCart) UpdateQuantity(ctx context.Context, userID, cartID uint64, quantity uint) error {
	row, err := s.getOwned(ctx, userID, cartID)
	if err != nil {
		return err
	}

	goods, err := service.Goods().GetById(ctx, row.GoodsId)
	if err != nil {
		return err
	}
	if quantity > goods.Stock {
		return gerror.Newf("库存不足，最多可购买 %d 件", goods.Stock)
	}

	cols := dao.Cart.Columns()
	_, err = dao.Cart.Ctx(ctx).
		Where(cols.Id, cartID).
		Where(cols.UserId, userID).
		Data(g.Map{cols.Quantity: quantity}).
		Update()
	return err
}

// Remove 删除购物车项
func (s *sCart) Remove(ctx context.Context, userID, cartID uint64) error {
	cols := dao.Cart.Columns()
	res, err := dao.Cart.Ctx(ctx).
		Where(cols.Id, cartID).
		Where(cols.UserId, userID).
		Delete()
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return gerror.NewCode(gcode.CodeNotFound, "购物车项不存在")
	}
	return nil
}

// Clear 清空当前用户的购物车
func (s *sCart) Clear(ctx context.Context, userID uint64) (int, error) {
	cols := dao.Cart.Columns()
	res, err := dao.Cart.Ctx(ctx).Where(cols.UserId, userID).Delete()
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// getByGoods 查询用户购物车中指定商品的记录，不存在返回 nil
func (s *sCart) getByGoods(ctx context.Context, userID, goodsID uint64) (*entity.Cart, error) {
	cols := dao.Cart.Columns()
	var row *entity.Cart
	err := dao.Cart.Ctx(ctx).
		Where(cols.UserId, userID).
		Where(cols.GoodsId, goodsID).
		Scan(&row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// getOwned 查询属于该用户的购物车项。
// 不存在与不属于该用户返回同一条提示，避免他人借此探测数据是否存在。
func (s *sCart) getOwned(ctx context.Context, userID, cartID uint64) (*entity.Cart, error) {
	cols := dao.Cart.Columns()
	var row *entity.Cart
	err := dao.Cart.Ctx(ctx).
		Where(cols.Id, cartID).
		Where(cols.UserId, userID).
		Scan(&row)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "购物车项不存在")
	}
	return row, nil
}
