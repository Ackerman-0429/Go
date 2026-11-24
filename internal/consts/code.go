package consts

import "github.com/gogf/gf/v2/errors/gcode"

// 业务错误码。
// 区间约定：40xxx 权限类，50xxx 业务校验类，避免与 gcode 内置码（0~1000）冲突。
var (
	// CodeNotAdmin 非管理员访问后台接口
	CodeNotAdmin = gcode.New(40301, "需要管理员权限", nil)

	// CodeGoodsOffSale 商品已下架，不可操作
	CodeGoodsOffSale = gcode.New(50001, "商品已下架", nil)

	// CodeStockNotEnough 商品库存不足，加购或下单被拒
	CodeStockNotEnough = gcode.New(50002, "库存不足", nil)

	// CodeCartChanged 下单时购物车内容已变化，需刷新后重试
	CodeCartChanged = gcode.New(50003, "购物车已发生变化", nil)
)
