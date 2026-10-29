// Package v1 定义商品模块 v1 版本的接口出入参结构。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 商品列表请求
type ListReq struct {
	g.Meta     `path:"/goods/list" method:"get" summary:"商品列表" tags:"商品"`
	Page       int    `json:"page"       d:"1"  v:"min:1#页码最小为1" dc:"页码，默认1"`
	PageSize   int    `json:"pageSize"   d:"10" v:"between:1,100#每页数量需在1-100之间" dc:"每页数量，默认10"`
	CategoryID uint64 `json:"categoryId" dc:"按分类筛选，0或不传表示全部"`
	Keyword    string `json:"keyword"    dc:"按商品名模糊搜索"`
}

// ListItem 列表项。只返回浏览所需字段，不含 detail 大文本。
type ListItem struct {
	ID         uint64  `json:"id"         dc:"商品ID"`
	Name       string  `json:"name"       dc:"商品名称"`
	Subtitle   string  `json:"subtitle"   dc:"副标题"`
	Price      float64 `json:"price"      dc:"售价"`
	Cover      string  `json:"cover"      dc:"封面图"`
	Sales      uint    `json:"sales"      dc:"销量"`
	Stock      uint    `json:"stock"      dc:"库存"`
	CategoryID uint    `json:"categoryId" dc:"分类ID"`
}

// ListRes 分页结果
type ListRes struct {
	List     []ListItem `json:"list"     dc:"商品列表"`
	Total    int        `json:"total"    dc:"总条数"`
	Page     int        `json:"page"     dc:"当前页"`
	PageSize int        `json:"pageSize" dc:"每页数量"`
}

// DetailReq 商品详情请求
type DetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" summary:"商品详情" tags:"商品"`
	ID     uint64 `json:"id" v:"required|min:1#请提供商品ID|商品ID不合法" dc:"商品ID"`
}

// DetailRes 商品详情
type DetailRes struct {
	ID          uint64   `json:"id"          dc:"商品ID"`
	CategoryID  uint     `json:"categoryId"  dc:"分类ID"`
	Name        string   `json:"name"        dc:"商品名称"`
	Subtitle    string   `json:"subtitle"    dc:"副标题"`
	Price       float64  `json:"price"       dc:"售价"`
	Stock       uint     `json:"stock"       dc:"库存"`
	Cover       string   `json:"cover"       dc:"封面图"`
	Images      []string `json:"images"      dc:"多图列表"`
	Detail      string   `json:"detail"      dc:"商品详情"`
	Sales       uint     `json:"sales"       dc:"销量"`
	Rating      float64  `json:"rating"      dc:"平均评分，无评价时为 0"`
	ReviewCount int      `json:"reviewCount" dc:"评价条数"`
}

// ReviewListReq 商品评价列表
type ReviewListReq struct {
	g.Meta   `path:"/goods/reviews" method:"get" summary:"商品评价列表" tags:"商品"`
	GoodsID  uint64 `json:"goodsId"  v:"required|min:1#请提供商品ID|商品ID不合法"               dc:"商品ID"`
	Page     int    `json:"page"     d:"1"  v:"min:1#页码最小为1"                       dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"between:1,50#每页数量需在1-50之间"           dc:"每页数量"`
	Rating   int    `json:"rating"   d:"0"  v:"in:0,1,2,3,4,5#评分筛选只能是0-5"          dc:"按星级筛选，0 表示全部"`
}

// ReviewItem 单条评价。匿名评价不返回真实昵称，连用户 ID 也不给。
type ReviewItem struct {
	ID        uint64 `json:"id"        dc:"评价ID"`
	Nickname  string `json:"nickname"  dc:"评价人昵称，匿名时固定为「匿名用户」"`
	Rating    uint   `json:"rating"    dc:"评分 1-5 星"`
	Content   string `json:"content"   dc:"评价内容"`
	Anonymous bool   `json:"anonymous" dc:"是否匿名"`
	CreatedAt string `json:"createdAt" dc:"评价时间"`
}

// ReviewListRes 评价列表 + 评分汇总，一次往返同时喂给列表和星级筛选条
type ReviewListRes struct {
	List      []ReviewItem `json:"list"      dc:"评价列表"`
	Total     int          `json:"total"     dc:"当前筛选下的条数"`
	Page      int          `json:"page"      dc:"当前页"`
	PageSize  int          `json:"pageSize"  dc:"每页数量"`
	AvgRating float64      `json:"avgRating" dc:"平均评分，保留一位小数"`
	Count     int          `json:"count"     dc:"全部评价条数（不受星级筛选影响）"`
	Counts    []int        `json:"counts"    dc:"各星级条数，下标 0 对应 1 星"`
}
