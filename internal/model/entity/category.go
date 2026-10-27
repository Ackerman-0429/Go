// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure for table category.
type Category struct {
	Id        uint        `json:"id"        orm:"id"         description:"分类ID"`        // 分类ID
	ParentId  uint        `json:"parentId"  orm:"parent_id"  description:"父分类ID，0表示顶级"` // 父分类ID，0表示顶级
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`        // 分类名称
	Sort      int         `json:"sort"      orm:"sort"       description:"排序值，越小越靠前"`   // 排序值，越小越靠前
	Status    uint        `json:"status"    orm:"status"     description:"状态：1启用 0停用"`  // 状态：1启用 0停用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`            //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`            //
}
