package goods

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// formatTime 时间格式化，空值返回空字符串，避免前端显示 null。
// 各控制器包里各留一份六行的小工具，比为一个格式化函数引入共享包更省事。
func formatTime(t *gtime.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("Y-m-d H:i:s")
}
