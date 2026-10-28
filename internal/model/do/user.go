// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        any         // 用户ID
	Username  any         // 用户名，登录用
	Password  any         // 密码哈希(bcrypt)，禁止存明文
	Nickname  any         // 昵称
	Email     any         // 邮箱，可为空但不可重复
	Phone     any         // 手机号，可为空但不可重复
	Status    any         // 状态：1正常 0禁用
	Role      any         // 角色：0普通用户 1管理员
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
