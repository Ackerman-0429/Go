// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint64      `json:"id"        orm:"id"         description:"用户ID"`               // 用户ID
	Username  string      `json:"username"  orm:"username"   description:"用户名，登录用"`            // 用户名，登录用
	Password  string      `json:"password"  orm:"password"   description:"密码哈希(bcrypt)，禁止存明文"` // 密码哈希(bcrypt)，禁止存明文
	Nickname  string      `json:"nickname"  orm:"nickname"   description:"昵称"`                 // 昵称
	Email     string      `json:"email"     orm:"email"      description:"邮箱，可为空但不可重复"`        // 邮箱，可为空但不可重复
	Phone     string      `json:"phone"     orm:"phone"      description:"手机号，可为空但不可重复"`       // 手机号，可为空但不可重复
	Status    uint        `json:"status"    orm:"status"     description:"状态：1正常 0禁用"`         // 状态：1正常 0禁用
	Role      uint        `json:"role"      orm:"role"       description:"角色：0普通用户 1管理员"`      // 角色：0普通用户 1管理员
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                   //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`                   //
}
