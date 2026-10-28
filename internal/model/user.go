package model

// UserProfileInput 用户自助修改资料。
// 只含用户自己能决定的三项；用户名、角色、状态都不在这里，避免越权。
type UserProfileInput struct {
	ID       uint64
	Nickname string
	Email    string
	Phone    string
}

// UserPasswordInput 用户自助修改密码，必须带上原密码做校验
type UserPasswordInput struct {
	ID          uint64
	OldPassword string
	NewPassword string
}
