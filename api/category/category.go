// Package category 定义分类模块的接口契约，由 controller 层实现。
package category

import (
	"context"

	v1 "hello/api/category/v1"
)

// ICategoryV1 分类模块接口
type ICategoryV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
}
