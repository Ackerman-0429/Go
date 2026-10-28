package admin

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// formatTime 时间格式化，空值返回空字符串，避免前端显示 null
func formatTime(t *gtime.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("Y-m-d H:i:s")
}
