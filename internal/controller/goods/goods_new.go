// Package goods 实现商品模块的控制器，只负责参数接收与响应组装。
package goods

import (
	goodsapi "hello/api/goods"
)

// ControllerV1 商品模块控制器
type ControllerV1 struct{}

// NewV1 由 cmd 层调用，完成接口注入
func NewV1() goodsapi.IGoodsV1 {
	return &ControllerV1{}
}
