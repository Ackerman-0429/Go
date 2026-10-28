package model

// GoodsListInput 商品列表查询条件
type GoodsListInput struct {
	Page       int
	PageSize   int
	CategoryID uint64
	Keyword    string
}
