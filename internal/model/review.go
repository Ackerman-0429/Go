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

// ReviewAdminListInput 后台评价列表查询条件。Status 为 -1、Rating 为 0 表示不限。
type ReviewAdminListInput struct {
	Page     int
	PageSize int
	Status   int
	Rating   int
	GoodsID  uint64
	Keyword  string
}

// ReviewAdminItem 后台评价列表项。
// 带商品名与用户信息，匿名评价不替换昵称——管理员需要知道是谁。
type ReviewAdminItem struct {
	ID        uint64
	GoodsID   uint64
	GoodsName string
	UserID    uint64
	Username  string
	Nickname  string
	Rating    uint
	Content   string
	Anonymous bool
	Status    uint
	CreatedAt *gtime.Time
}

// ReviewAdminStats 后台评价计数。
// 跟随除状态以外的筛选条件，用于列表顶部「显示中 / 已隐藏」两个 tab 的数字。
type ReviewAdminStats struct {
	Visible int
	Hidden  int
}
