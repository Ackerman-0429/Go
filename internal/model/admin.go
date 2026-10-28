package model

import "github.com/gogf/gf/v2/os/gtime"

// GoodsAdminListInput 后台商品列表查询条件。与前台不同，Status 可筛选下架商品。
type GoodsAdminListInput struct {
	Page       int
	PageSize   int
	Keyword    string
	CategoryID uint64
	Status     int // -1 表示全部
}

// GoodsSaveInput 后台新建/编辑商品。ID 为 0 表示新建。
type GoodsSaveInput struct {
	ID         uint64
	CategoryID uint64
	Name       string
	Subtitle   string
	Price      float64
	Stock      uint
	Cover      string
	Images     string
	Detail     string
	Status     uint
}

// OrderAdminListInput 后台订单列表查询条件，跨用户查询
type OrderAdminListInput struct {
	Page     int
	PageSize int
	Status   int // -1 表示全部
	Keyword  string
}

// OrderAdminItem 后台订单条目，比用户端多了下单人信息
type OrderAdminItem struct {
	ID              uint64
	OrderNo         string
	UserID          uint64
	Username        string
	TotalAmount     float64
	Status          uint
	Quantity        uint
	ItemCount       int
	FirstName       string
	ReceiverName    string
	ReceiverPhone   string
	ReceiverAddress string
	CreatedAt       *gtime.Time
	PaidAt          *gtime.Time
}

// AdminStats 后台概览统计
type AdminStats struct {
	TodayOrderCount  int
	TodaySales       float64
	PendingShipCount int
	TotalOrderCount  int
	OnSaleGoodsCount int
	UserCount        int
}

// CategoryAdminItem 后台分类条目。GoodsCount 对顶级分类会累加其子分类的商品数。
type CategoryAdminItem struct {
	ID         uint
	ParentID   uint
	ParentName string
	Name       string
	Sort       int
	Status     uint
	GoodsCount int
	Children   int
	CreatedAt  *gtime.Time
}

// CategorySaveInput 后台新建/编辑分类。ID 为 0 表示新建。
type CategorySaveInput struct {
	ID       uint64
	ParentID uint64
	Name     string
	Sort     int
	Status   uint
}
