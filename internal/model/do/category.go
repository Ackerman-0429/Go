// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table category for DAO operations like Where/Data.
type Category struct {
	g.Meta    `orm:"table:category, do:true"`
	Id        any         // 分类ID
	ParentId  any         // 父分类ID，0表示顶级
	Name      any         // 分类名称
	Sort      any         // 排序值，越小越靠前
	Status    any         // 状态：1启用 0停用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
