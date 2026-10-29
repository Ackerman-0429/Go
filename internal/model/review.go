package model

import "github.com/gogf/gf/v2/os/gtime"

// ReviewCreateInput 发表评价。用订单明细 ID 定位商品，
// 前端不需要传 goodsId/orderId，少一个可被伪造的参数。
type ReviewCreateInput struct {
	OrderItemID uint64
	Rating      uint
	Content     string
	Anonymous   bool
}

// ReviewListInput 商品评价列表查询条件。Rating 为 0 表示全部星级。
type ReviewListInput struct {
	GoodsID  uint64
	Page     int
	PageSize int
	Rating   int
}

// ReviewItem 评价列表项，昵称已按匿名规则处理过
type ReviewItem struct {
	ID        uint64
	Nickname  string
	Rating    uint
	Content   string
	Anonymous bool
	CreatedAt *gtime.Time
}

// ReviewSummary 商品评分汇总
type ReviewSummary struct {
	AvgRating float64
	Count     int
	Counts    [5]int // Counts[0] 对应 1 星，依次到 5 星
}
