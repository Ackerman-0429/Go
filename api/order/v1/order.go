// Package v1 定义订单模块 v1 版本的接口出入参结构。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CreateReq 提交订单（把购物车中可结算的商品生成订单）
type CreateReq struct {
	g.Meta          `path:"/order/create" method:"post" summary:"提交订单" tags:"订单"`
	ReceiverName    string `json:"receiverName"    v:"required|length:2,50#请填写收货人|收货人姓名长度需在2-50之间" dc:"收货人"`
	ReceiverPhone   string `json:"receiverPhone"   v:"required|phone#请填写手机号|手机号格式不正确"             dc:"收货电话"`
	ReceiverAddress string `json:"receiverAddress" v:"required|length:5,255#请填写收货地址|收货地址长度需在5-255之间" dc:"收货地址"`
	Remark          string `json:"remark"          v:"max-length:255#备注不能超过255个字符"               dc:"订单备注"`
}

// CreateRes 下单结果
type CreateRes struct {
	OrderID     uint64  `json:"orderId"     dc:"订单ID"`
	OrderNo     string  `json:"orderNo"     dc:"订单号"`
	TotalAmount float64 `json:"totalAmount" dc:"订单金额"`
	Quantity    uint    `json:"quantity"    dc:"商品总件数"`
}

// ListReq 我的订单列表
type ListReq struct {
	g.Meta   `path:"/order/list" method:"get" summary:"我的订单" tags:"订单"`
	Page     int `json:"page"     d:"1"  v:"min:1#页码最小为1"                dc:"页码，默认1"`
	PageSize int `json:"pageSize" d:"10" v:"between:1,100#每页数量需在1-100之间" dc:"每页数量，默认10"`
	Status   int `json:"status"   d:"-1" v:"in:-1,0,1,2,3,4#状态不合法"      dc:"按状态筛选，-1 表示全部"`
}

// ListItem 订单列表项
type ListItem struct {
	ID          uint64  `json:"id"          dc:"订单ID"`
	OrderNo     string  `json:"orderNo"     dc:"订单号"`
	TotalAmount float64 `json:"totalAmount" dc:"订单金额"`
	Status      uint    `json:"status"      dc:"状态：0待付款 1已付款 2已发货 3已完成 4已取消"`
	StatusText  string  `json:"statusText"  dc:"状态文案"`
	Quantity    uint    `json:"quantity"    dc:"商品总件数"`
	ItemCount   int     `json:"itemCount"   dc:"商品种类数"`
	FirstName   string  `json:"firstName"   dc:"首个商品名称"`
	FirstCover  string  `json:"firstCover"  dc:"首个商品封面"`
	CreatedAt   string  `json:"createdAt"   dc:"下单时间"`
}

// ListRes 分页结果
type ListRes struct {
	List     []ListItem `json:"list"     dc:"订单列表"`
	Total    int        `json:"total"    dc:"总条数"`
	Page     int        `json:"page"     dc:"当前页"`
	PageSize int        `json:"pageSize" dc:"每页数量"`
}

// DetailReq 订单详情
type DetailReq struct {
	g.Meta `path:"/order/detail" method:"get" summary:"订单详情" tags:"订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// DetailItem 订单明细
type DetailItem struct {
	GoodsID   uint64  `json:"goodsId"   dc:"商品ID"`
	GoodsName string  `json:"goodsName" dc:"下单时商品名"`
	Cover     string  `json:"cover"     dc:"下单时封面图"`
	Price     float64 `json:"price"     dc:"下单时单价"`
	Quantity  uint    `json:"quantity"  dc:"数量"`
	Subtotal  float64 `json:"subtotal"  dc:"小计"`
}

// DetailRes 订单详情
type DetailRes struct {
	ID              uint64       `json:"id"              dc:"订单ID"`
	OrderNo         string       `json:"orderNo"         dc:"订单号"`
	TotalAmount     float64      `json:"totalAmount"     dc:"订单金额"`
	Status          uint         `json:"status"          dc:"状态"`
	StatusText      string       `json:"statusText"      dc:"状态文案"`
	ReceiverName    string       `json:"receiverName"    dc:"收货人"`
	ReceiverPhone   string       `json:"receiverPhone"   dc:"收货电话"`
	ReceiverAddress string       `json:"receiverAddress" dc:"收货地址"`
	Remark          string       `json:"remark"          dc:"备注"`
	PaidAt          string       `json:"paidAt"          dc:"付款时间，未付款为空"`
	CreatedAt       string       `json:"createdAt"       dc:"下单时间"`
	Items           []DetailItem `json:"items"           dc:"订单明细"`
}

// CancelReq 取消订单
type CancelReq struct {
	g.Meta `path:"/order/cancel" method:"post" summary:"取消订单" tags:"订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// CancelRes 取消结果
type CancelRes struct{}

// PayReq 模拟支付
type PayReq struct {
	g.Meta `path:"/order/pay" method:"post" summary:"模拟支付（演示用）" tags:"订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// PayRes 支付结果
type PayRes struct {
	Status     uint   `json:"status"     dc:"支付后状态"`
	StatusText string `json:"statusText" dc:"状态文案"`
	PaidAt     string `json:"paidAt"     dc:"付款时间"`
}
