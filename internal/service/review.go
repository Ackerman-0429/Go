// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
)

type (
	IReview interface {
		// Create 对已完成订单中的某个商品发表评价。
		// 校验顺序：明细存在 → 属于当前用户 → 订单已完成 → 未重复评价。
		Create(ctx context.Context, userID uint64, in model.ReviewCreateInput) (uint64, error)
		// List 商品评价分页列表。只返回显示中的评价，匿名评价替换昵称。
		List(ctx context.Context, in model.ReviewListInput) ([]model.ReviewItem, int, error)
		// Summary 商品评分汇总：平均分、总条数、各星级分布。
		// 一次 GROUP BY rating 拿到，不逐条累加。
		Summary(ctx context.Context, goodsID uint64) (*model.ReviewSummary, error)
		// AdminList 后台评价列表：可按状态/星级/商品筛选，支持内容、商品名、用户名搜索。
		// 同时返回「显示中 / 已隐藏」两个计数，供列表顶部的 tab 使用。
		AdminList(ctx context.Context, in model.ReviewAdminListInput) ([]model.ReviewAdminItem, int, model.ReviewAdminStats, error)
		// AdminSetStatus 显示或隐藏评价。前台列表与评分汇总都只算 status=1 的，
		// 所以隐藏后该条立刻从前台消失，重新显示则恢复。
		AdminSetStatus(ctx context.Context, id uint64, status uint) error
	}
)

var (
	localReview IReview
)

func Review() IReview {
	if localReview == nil {
		panic("implement not found for interface IReview, forgot register?")
	}
	return localReview
}

func RegisterReview(i IReview) {
	localReview = i
}
