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
