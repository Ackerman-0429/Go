// Package v1 定义分类模块 v1 版本的接口出入参结构。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 分类列表请求
type ListReq struct {
	g.Meta `path:"/category/list" method:"get" summary:"分类列表" tags:"分类"`
}

// ListItem 分类项
type ListItem struct {
	ID       uint   `json:"id"       dc:"分类ID"`
	ParentID uint   `json:"parentId" dc:"父分类ID，0为顶级"`
	Name     string `json:"name"     dc:"分类名称"`
	Sort     int    `json:"sort"     dc:"排序值"`
}

// ListRes 分类列表。平铺返回，由前端按 parentId 自行组装成树。
type ListRes struct {
	List []ListItem `json:"list" dc:"分类列表"`
}
