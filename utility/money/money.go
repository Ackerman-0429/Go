// Package money 提供以「分」为单位的金额换算，避免 float64 直接参与金额运算。
package money

import "math"

// ToCents 把元换算成分。商品表的 decimal 被 gf gen dao 映射为 float64，
// 直接累加会出现 0.1+0.2 这类精度误差，统一转成整数分再算。
func ToCents(yuan float64) int64 {
	return int64(math.Round(yuan * 100))
}

// FromCents 把分换算回元
func FromCents(cents int64) float64 {
	return float64(cents) / 100
}
