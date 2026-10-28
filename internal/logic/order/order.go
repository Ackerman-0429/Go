// Package order 实现订单模块的业务逻辑：下单、列表、详情、取消、模拟支付。
package order

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"

	"hello/internal/consts"
	"hello/internal/dao"
	"hello/internal/model"
	"hello/internal/model/entity"
	"hello/internal/service"
	"hello/utility/money"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(&sOrder{})
}

// genOrderNo 生成订单号：14 位时间 + 6 位随机数
func genOrderNo() string {
	return gtime.Now().Format("YmdHis") + fmt.Sprintf("%06d", grand.N(100000, 999999))
}

// Create 把购物车中可结算的商品生成订单。
// 扣库存、写订单、写明细、清购物车在同一个事务里，任何一步失败整体回滚。
func (s *sOrder) Create(ctx context.Context, in model.OrderCreateInput) (*model.OrderCreateOutput, error) {
	var out *model.OrderCreateOutput

	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		orderCols := dao.Orders.Columns()
		itemCols := dao.OrderItem.Columns()
		goodsCols := dao.Goods.Columns()
		cartCols := dao.Cart.Columns()

		cartOut, err := service.Cart().List(ctx, in.UserID)
		if err != nil {
			return err
		}
		if len(cartOut.List) == 0 {
			return gerror.NewCode(gcode.CodeNotFound, "购物车是空的")
		}

		var (
			totalCents int64
			totalQty   uint
			cartIDs    = make([]uint64, 0, len(cartOut.List))
			itemRows   = make([]g.Map, 0, len(cartOut.List))
		)

		for _, item := range cartOut.List {
			// 下架、售罄的条目直接跳过，不阻断其他商品的结算
			if !item.Available {
				continue
			}

			// 条件更新：只有库存足够时才扣减，并发下不会超卖
			res, err := dao.Goods.Ctx(ctx).
				Where(goodsCols.Id, item.GoodsID).
				Where(goodsCols.Status, consts.GoodsStatusOnSale).
				Where("stock >= ?", item.Quantity).
				Decrement(goodsCols.Stock, item.Quantity)
			if err != nil {
				return err
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return err
			}
			if affected == 0 {
				return gerror.Newf("「%s」库存不足，请返回购物车调整数量", item.Name)
			}

			cents := money.ToCents(item.Price) * int64(item.Quantity)
			totalCents += cents
			totalQty += item.Quantity
			cartIDs = append(cartIDs, item.ID)

			// 商品名、价格、封面存快照，商品后续改价改名不影响历史订单
			itemRows = append(itemRows, g.Map{
				itemCols.GoodsId:    item.GoodsID,
				itemCols.GoodsName:  item.Name,
				itemCols.GoodsCover: item.Cover,
				itemCols.GoodsPrice: money.FromCents(money.ToCents(item.Price)),
				itemCols.Quantity:   item.Quantity,
				itemCols.Subtotal:   money.FromCents(cents),
				itemCols.CreatedAt:  gtime.Now(),
			})
		}

		if len(itemRows) == 0 {
			return gerror.NewCode(gcode.CodeNotFound, "购物车中没有可结算的商品")
		}

		orderNo := genOrderNo()
		orderID, err := dao.Orders.Ctx(ctx).Data(g.Map{
			orderCols.OrderNo:         orderNo,
			orderCols.UserId:          in.UserID,
			orderCols.TotalAmount:     money.FromCents(totalCents),
			orderCols.Status:          consts.OrderStatusPending,
			orderCols.ReceiverName:    in.ReceiverName,
			orderCols.ReceiverPhone:   in.ReceiverPhone,
			orderCols.ReceiverAddress: in.ReceiverAddress,
			orderCols.Remark:          in.Remark,
		}).InsertAndGetId()
		if err != nil {
			return err
		}

		for _, row := range itemRows {
			row[itemCols.OrderId] = uint64(orderID)
		}
		if _, err = dao.OrderItem.Ctx(ctx).Data(itemRows).Insert(); err != nil {
			return err
		}

		// 删除已下单的购物车项。影响行数少于预期说明购物车被并发修改（例如重复提交），
		// 此时回滚，避免同一批购物车生成两笔订单。
		res, err := dao.Cart.Ctx(ctx).
			Where(cartCols.UserId, in.UserID).
			WhereIn(cartCols.Id, cartIDs).
			Delete()
		if err != nil {
			return err
		}
		deleted, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if int(deleted) != len(cartIDs) {
			return gerror.New("购物车已发生变化，请刷新后重试")
		}

		out = &model.OrderCreateOutput{
			OrderID:     uint64(orderID),
			OrderNo:     orderNo,
			TotalAmount: money.FromCents(totalCents),
			Quantity:    totalQty,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// List 分页查询当前用户的订单，商品信息由明细聚合
func (s *sOrder) List(ctx context.Context, in model.OrderListInput) ([]model.OrderListItem, int, error) {
	orderCols := dao.Orders.Columns()
	itemCols := dao.OrderItem.Columns()

	m := dao.Orders.Ctx(ctx).Where(orderCols.UserId, in.UserID)
	if in.Status >= 0 {
		m = m.Where(orderCols.Status, in.Status)
	}

	var (
		orders []entity.Orders
		total  int
	)
	if err := m.Page(in.Page, in.PageSize).
		OrderDesc(orderCols.Id).
		ScanAndCount(&orders, &total, false); err != nil {
		return nil, 0, err
	}
	if len(orders) == 0 {
		return []model.OrderListItem{}, total, nil
	}

	ids := make([]uint64, 0, len(orders))
	for _, o := range orders {
		ids = append(ids, o.Id)
	}

	var items []entity.OrderItem
	if err := dao.OrderItem.Ctx(ctx).
		WhereIn(itemCols.OrderId, ids).
		OrderAsc(itemCols.Id).
		Scan(&items); err != nil {
		return nil, 0, err
	}

	agg := make(map[uint64]*model.OrderListItem, len(orders))
	for _, o := range orders {
		agg[o.Id] = &model.OrderListItem{
			ID:          o.Id,
			OrderNo:     o.OrderNo,
			TotalAmount: o.TotalAmount,
			Status:      o.Status,
			CreatedAt:   o.CreatedAt,
		}
	}
	for _, it := range items {
		a, ok := agg[it.OrderId]
		if !ok {
			continue
		}
		if a.ItemCount == 0 {
			a.FirstName = it.GoodsName
			a.FirstCover = it.GoodsCover
		}
		a.ItemCount++
		a.Quantity += it.Quantity
	}

	out := make([]model.OrderListItem, 0, len(orders))
	for _, o := range orders {
		out = append(out, *agg[o.Id])
	}
	return out, total, nil
}

// Detail 查询订单详情，订单与明细都限定在当前用户范围内
func (s *sOrder) Detail(ctx context.Context, userID, orderID uint64) (*model.OrderDetail, error) {
	orderCols := dao.Orders.Columns()
	itemCols := dao.OrderItem.Columns()

	var order *entity.Orders
	if err := dao.Orders.Ctx(ctx).
		Where(orderCols.Id, orderID).
		Where(orderCols.UserId, userID).
		Scan(&order); err != nil {
		return nil, err
	}
	if order == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "订单不存在")
	}

	var items []entity.OrderItem
	if err := dao.OrderItem.Ctx(ctx).
		Where(itemCols.OrderId, orderID).
		OrderAsc(itemCols.Id).
		Scan(&items); err != nil {
		return nil, err
	}
	return &model.OrderDetail{Order: order, Items: items}, nil
}

// Cancel 取消待付款订单并回滚库存
func (s *sOrder) Cancel(ctx context.Context, userID, orderID uint64) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		orderCols := dao.Orders.Columns()
		itemCols := dao.OrderItem.Columns()
		goodsCols := dao.Goods.Columns()

		// 带上 status 条件，天然防止重复取消（第二次更新影响 0 行）
		res, err := dao.Orders.Ctx(ctx).
			Where(orderCols.Id, orderID).
			Where(orderCols.UserId, userID).
			Where(orderCols.Status, consts.OrderStatusPending).
			Data(g.Map{orderCols.Status: consts.OrderStatusCancelled}).
			Update()
		if err != nil {
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return gerror.NewCode(gcode.CodeNotFound, "订单不存在或当前状态不可取消")
		}

		var items []entity.OrderItem
		if err = dao.OrderItem.Ctx(ctx).Where(itemCols.OrderId, orderID).Scan(&items); err != nil {
			return err
		}
		for _, it := range items {
			if _, err = dao.Goods.Ctx(ctx).
				Where(goodsCols.Id, it.GoodsId).
				Increment(goodsCols.Stock, it.Quantity); err != nil {
				return err
			}
		}
		return nil
	})
}

// Pay 模拟支付：待付款改为已付款，并把明细数量累加进商品销量。
// 真实项目这里要接支付网关与异步回调，此处仅演示状态流转。
func (s *sOrder) Pay(ctx context.Context, userID, orderID uint64) (*gtime.Time, error) {
	paidAt := gtime.Now()

	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		orderCols := dao.Orders.Columns()
		itemCols := dao.OrderItem.Columns()
		goodsCols := dao.Goods.Columns()

		res, err := dao.Orders.Ctx(ctx).
			Where(orderCols.Id, orderID).
			Where(orderCols.UserId, userID).
			Where(orderCols.Status, consts.OrderStatusPending).
			Data(g.Map{
				orderCols.Status: consts.OrderStatusPaid,
				orderCols.PaidAt: paidAt,
			}).
			Update()
		if err != nil {
			return err
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if affected == 0 {
			return gerror.NewCode(gcode.CodeNotFound, "订单不存在或已支付")
		}

		var items []entity.OrderItem
		if err = dao.OrderItem.Ctx(ctx).Where(itemCols.OrderId, orderID).Scan(&items); err != nil {
			return err
		}
		for _, it := range items {
			if _, err = dao.Goods.Ctx(ctx).
				Where(goodsCols.Id, it.GoodsId).
				Increment(goodsCols.Sales, it.Quantity); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paidAt, nil
}
