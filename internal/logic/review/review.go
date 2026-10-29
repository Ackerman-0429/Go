// Package review 实现商品评价：发表、列表、评分汇总。
// 写入侧挂在订单模块（一条订单明细评一次），读取侧给商品详情页用，
// 所以逻辑集中在这里，goods / order 两个控制器都调它。
package review

import (
	"context"
	"math"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
)

type sReview struct{}

func init() {
	service.RegisterReview(&sReview{})
}

// 匿名评价对外展示的固定昵称，避免前端各自拼一个
const anonymousName = "匿名用户"

// Create 对已完成订单中的某个商品发表评价。
// 校验顺序：明细存在 → 属于当前用户 → 订单已完成 → 未重复评价。
func (s *sReview) Create(ctx context.Context, userID uint64, in model.ReviewCreateInput) (uint64, error) {
	cols := dao.GoodsReview.Columns()

	var item *entity.OrderItem
	if err := dao.OrderItem.Ctx(ctx).WherePri(in.OrderItemID).Scan(&item); err != nil {
		return 0, err
	}
	if item == nil {
		return 0, gerror.NewCode(gcode.CodeNotFound, "订单明细不存在")
	}

	var order *entity.Orders
	if err := dao.Orders.Ctx(ctx).WherePri(item.OrderId).Scan(&order); err != nil {
		return 0, err
	}
	// 不是自己的订单统一按「不存在」处理，不透露别人的订单有没有这个商品
	if order == nil || order.UserId != userID {
		return 0, gerror.NewCode(gcode.CodeNotFound, "订单明细不存在")
	}
	// 只有签收完成的订单才能评价，否则下单当天就能刷好评
	if order.Status != consts.OrderStatusFinished {
		return 0, gerror.NewCode(gcode.CodeInvalidParameter, "订单完成后才能评价")
	}

	exists, err := dao.GoodsReview.Ctx(ctx).Where(cols.OrderItemId, in.OrderItemID).Count()
	if err != nil {
		return 0, err
	}
	if exists > 0 {
		return 0, gerror.NewCode(gcode.CodeInvalidParameter, "该商品已经评价过了")
	}

	id, err := dao.GoodsReview.Ctx(ctx).Data(g.Map{
		cols.GoodsId:     item.GoodsId,
		cols.OrderId:     item.OrderId,
		cols.OrderItemId: item.Id,
		cols.UserId:      userID,
		cols.Rating:      in.Rating,
		cols.Content:     strings.TrimSpace(in.Content),
		cols.Anonymous:   boolToTiny(in.Anonymous),
		cols.Status:      1,
	}).InsertAndGetId()
	if err != nil {
		// 并发下两个请求可能同时通过上面的 Count，唯一键兜底，转成同样的提示
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, gerror.NewCode(gcode.CodeInvalidParameter, "该商品已经评价过了")
		}
		return 0, err
	}
	return uint64(id), nil
}

// List 商品评价分页列表。只返回显示中的评价，匿名评价替换昵称。
func (s *sReview) List(ctx context.Context, in model.ReviewListInput) ([]model.ReviewItem, int, error) {
	cols := dao.GoodsReview.Columns()

	m := dao.GoodsReview.Ctx(ctx).
		Where(cols.GoodsId, in.GoodsID).
		Where(cols.Status, 1)
	if in.Rating > 0 {
		m = m.Where(cols.Rating, in.Rating)
	}

	var (
		rows  []entity.GoodsReview
		total int
	)
	if err := m.Page(in.Page, in.PageSize).
		OrderDesc(cols.Id).
		ScanAndCount(&rows, &total, false); err != nil {
		return nil, 0, err
	}
	if len(rows) == 0 {
		return []model.ReviewItem{}, total, nil
	}

	// 批量取昵称，避免逐条查用户
	userIDs := make([]uint64, 0, len(rows))
	for _, r := range rows {
		if r.Anonymous == 0 {
			userIDs = append(userIDs, r.UserId)
		}
	}
	names, err := nicknamesByID(ctx, userIDs)
	if err != nil {
		return nil, 0, err
	}

	out := make([]model.ReviewItem, 0, len(rows))
	for _, r := range rows {
		nickname := names[r.UserId]
		anonymous := r.Anonymous == 1
		if anonymous {
			nickname = anonymousName
		} else if nickname == "" {
			// 用户已被删除，别显示空白
			nickname = "已注销用户"
		}
		out = append(out, model.ReviewItem{
			ID:        r.Id,
			Nickname:  nickname,
			Rating:    r.Rating,
			Content:   r.Content,
			Anonymous: anonymous,
			CreatedAt: r.CreatedAt,
		})
	}
	return out, total, nil
}

// Summary 商品评分汇总：平均分、总条数、各星级分布。
// 一次 GROUP BY rating 拿到，不逐条累加。
func (s *sReview) Summary(ctx context.Context, goodsID uint64) (*model.ReviewSummary, error) {
	cols := dao.GoodsReview.Columns()

	var rows []struct {
		Rating uint `orm:"rating"`
		Total  int  `orm:"total"`
	}
	if err := dao.GoodsReview.Ctx(ctx).
		Fields(cols.Rating, "COUNT(*) AS total").
		Where(cols.GoodsId, goodsID).
		Where(cols.Status, 1).
		Group(cols.Rating).
		Scan(&rows); err != nil {
		return nil, err
	}

	var (
		sum   int
		count int
		out   = &model.ReviewSummary{}
	)
	for _, r := range rows {
		if r.Rating < 1 || r.Rating > 5 {
			continue
		}
		out.Counts[r.Rating-1] = r.Total
		sum += int(r.Rating) * r.Total
		count += r.Total
	}
	out.Count = count
	if count > 0 {
		// 保留一位小数，4.6666 显示成 4.7
		out.AvgRating = math.Round(float64(sum)/float64(count)*10) / 10
	}
	return out, nil
}

// nicknamesByID 批量查昵称，昵称为空时回退到用户名
func nicknamesByID(ctx context.Context, ids []uint64) (map[uint64]string, error) {
	out := make(map[uint64]string, len(ids))
	if len(ids) == 0 {
		return out, nil
	}

	var users []entity.User
	if err := dao.User.Ctx(ctx).
		WhereIn(dao.User.Columns().Id, ids).
		Scan(&users); err != nil {
		return nil, err
	}
	for _, u := range users {
		if u.Nickname != "" {
			out[u.Id] = u.Nickname
		} else {
			out[u.Id] = u.Username
		}
	}
	return out, nil
}

// boolToTiny 布尔转 tinyint，库里存 1/0 而不是 Go 的 bool
func boolToTiny(b bool) uint {
	if b {
		return 1
	}
	return 0
}
