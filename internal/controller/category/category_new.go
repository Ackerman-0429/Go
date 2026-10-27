// Package category 实现分类模块的控制器。
package category

import (
	categoryapi "hello/api/category"
)

// ControllerV1 分类模块控制器
type ControllerV1 struct{}

// NewV1 由 cmd 层调用，完成接口注入
func NewV1() categoryapi.ICategoryV1 {
	return &ControllerV1{}
}
