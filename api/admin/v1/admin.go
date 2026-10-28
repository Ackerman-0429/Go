// Package v1 定义后台管理模块 v1 版本的接口出入参结构。
// 所有接口都挂在 /admin 分组下，由 Auth + Admin 中间件保护。
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------------------------------------------------------------
// 概览
// ---------------------------------------------------------------------------

// StatsReq 后台首页概览
type StatsReq struct {
	g.Meta `path:"/stats" method:"get" summary:"后台概览" tags:"后台"`
}

// StatsRes 概览统计
type StatsRes struct {
	TodayOrderCount  int     `json:"todayOrderCount"  dc:"今日订单数"`
	TodaySales       float64 `json:"todaySales"       dc:"今日销售额"`
	PendingShipCount int     `json:"pendingShipCount" dc:"待发货订单数"`
	TotalOrderCount  int     `json:"totalOrderCount"  dc:"累计订单数"`
	OnSaleGoodsCount int     `json:"onSaleGoodsCount" dc:"在售商品数"`
	UserCount        int     `json:"userCount"        dc:"注册用户数"`
}

// ---------------------------------------------------------------------------
// 商品管理
// ---------------------------------------------------------------------------

// GoodsListReq 后台商品列表，与前台不同：能看到下架商品
type GoodsListReq struct {
	g.Meta     `path:"/goods/list" method:"get" summary:"商品列表（含下架）" tags:"后台-商品"`
	Page       int    `json:"page"       d:"1"  v:"min:1#页码最小为1"                  dc:"页码"`
	PageSize   int    `json:"pageSize"   d:"10" v:"between:1,100#每页数量需在1-100之间" dc:"每页数量"`
	Keyword    string `json:"keyword"                                              dc:"按商品名搜索"`
	CategoryID uint64 `json:"categoryId"                                           dc:"按分类筛选，0 为全部"`
	Status     int    `json:"status"     d:"-1" v:"in:-1,0,1#状态只能是-1/0/1"        dc:"按状态筛选，-1 全部"`
}

// GoodsItem 后台商品条目
type GoodsItem struct {
	ID         uint64  `json:"id"         dc:"商品ID"`
	CategoryID uint64  `json:"categoryId" dc:"分类ID"`
	Name       string  `json:"name"       dc:"商品名称"`
	Subtitle   string  `json:"subtitle"   dc:"副标题"`
	Price      float64 `json:"price"      dc:"售价"`
	Stock      uint    `json:"stock"      dc:"库存"`
	Cover      string  `json:"cover"      dc:"封面图"`
	Images     string  `json:"images"     dc:"多图，JSON数组"`
	Detail     string  `json:"detail"     dc:"商品详情"`
	Sales      uint    `json:"sales"      dc:"销量"`
	Status     uint    `json:"status"     dc:"状态：1上架 0下架"`
	CreatedAt  string  `json:"createdAt"  dc:"创建时间"`
}

// GoodsListRes 分页结果
type GoodsListRes struct {
	List     []GoodsItem `json:"list"     dc:"商品列表"`
	Total    int         `json:"total"    dc:"总条数"`
	Page     int         `json:"page"     dc:"当前页"`
	PageSize int         `json:"pageSize" dc:"每页数量"`
}

// GoodsSaveReq 新建或编辑商品，ID 为空表示新建
type GoodsSaveReq struct {
	g.Meta     `path:"/goods/save" method:"post" summary:"新建/编辑商品" tags:"后台-商品"`
	ID         uint64  `json:"id"         dc:"商品ID，留空表示新建"`
	CategoryID uint64  `json:"categoryId" v:"required|min:1#请选择分类|分类不合法"        dc:"分类ID"`
	Name       string  `json:"name"       v:"required|length:2,100#请输入商品名称|商品名称长度需在2-100之间" dc:"商品名称"`
	Subtitle   string  `json:"subtitle"   v:"max-length:150#副标题不能超过150个字符"    dc:"副标题"`
	Price      float64 `json:"price"      v:"required|min:0.01#请输入售价|售价必须大于0"  dc:"售价"`
	Stock      uint    `json:"stock"      dc:"库存"`
	Cover      string  `json:"cover"      v:"max-length:255#封面地址过长"            dc:"封面图URL"`
	Images     string  `json:"images"     dc:"多图JSON数组，如 [\"url1\",\"url2\"]"`
	Detail     string  `json:"detail"     dc:"商品详情"`
	Status     uint    `json:"status"     d:"1" v:"in:0,1#状态只能是0或1"             dc:"状态：1上架 0下架"`
}

// GoodsSaveRes 保存结果
type GoodsSaveRes struct {
	ID uint64 `json:"id" dc:"商品ID"`
}

// GoodsStatusReq 上下架
type GoodsStatusReq struct {
	g.Meta `path:"/goods/status" method:"post" summary:"商品上下架" tags:"后台-商品"`
	ID     uint64 `json:"id"     v:"required|min:1#请提供商品ID|商品ID不合法" dc:"商品ID"`
	Status uint   `json:"status" v:"in:0,1#状态只能是0或1"            dc:"目标状态：1上架 0下架"`
}

// GoodsStatusRes 操作结果
type GoodsStatusRes struct {
	ID     uint64 `json:"id"     dc:"商品ID"`
	Status uint   `json:"status" dc:"当前状态"`
}

// GoodsDeleteReq 删除商品
type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"post" summary:"删除商品" tags:"后台-商品"`
	ID     uint64 `json:"id" v:"required|min:1#请提供商品ID|商品ID不合法" dc:"商品ID"`
}

// GoodsDeleteRes 删除结果
type GoodsDeleteRes struct{}

// ---------------------------------------------------------------------------
// 订单管理
// ---------------------------------------------------------------------------

// OrderListReq 全站订单列表。userId 非 0 时只看该用户的订单，
// 供用户管理页「查看该用户订单」复用同一个接口，不必再开一个平行接口。
type OrderListReq struct {
	g.Meta   `path:"/order/list" method:"get" summary:"全站订单列表" tags:"后台-订单"`
	Page     int    `json:"page"     d:"1"  v:"min:1#页码最小为1"                  dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"between:1,100#每页数量需在1-100之间" dc:"每页数量"`
	Status   int    `json:"status"   d:"-1" v:"in:-1,0,1,2,3,4#状态不合法"        dc:"按状态筛选，-1 全部"`
	Keyword  string `json:"keyword"                                              dc:"按订单号/收货人搜索"`
	UserId   uint64 `json:"userId"   d:"0"                                       dc:"按下单用户筛选，0 为全部"`
}

// OrderItem 后台订单条目
type OrderItem struct {
	ID              uint64  `json:"id"              dc:"订单ID"`
	OrderNo         string  `json:"orderNo"         dc:"订单号"`
	UserID          uint64  `json:"userId"          dc:"下单用户ID"`
	Username        string  `json:"username"        dc:"下单用户"`
	TotalAmount     float64 `json:"totalAmount"     dc:"订单金额"`
	Status          uint    `json:"status"          dc:"状态：0待付款 1已付款 2已发货 3已完成 4已取消"`
	StatusText      string  `json:"statusText"      dc:"状态文案"`
	Quantity        uint    `json:"quantity"        dc:"商品总件数"`
	ItemCount       int     `json:"itemCount"       dc:"商品种类数"`
	FirstName       string  `json:"firstName"       dc:"首个商品名称"`
	ReceiverName    string  `json:"receiverName"    dc:"收货人"`
	ReceiverPhone   string  `json:"receiverPhone"   dc:"收货电话"`
	ReceiverAddress string  `json:"receiverAddress" dc:"收货地址"`
	CreatedAt       string  `json:"createdAt"       dc:"下单时间"`
	PaidAt          string  `json:"paidAt"          dc:"付款时间"`
}

// OrderListRes 分页结果
type OrderListRes struct {
	List     []OrderItem `json:"list"     dc:"订单列表"`
	Total    int         `json:"total"    dc:"总条数"`
	Page     int         `json:"page"     dc:"当前页"`
	PageSize int         `json:"pageSize" dc:"每页数量"`
}

// OrderDetailReq 后台查看单笔订单详情，不受下单用户限制
type OrderDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" summary:"订单详情（后台）" tags:"后台-订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// OrderGoodsItem 订单内单个商品的快照信息
type OrderGoodsItem struct {
	GoodsID   uint64  `json:"goodsId"   dc:"商品ID"`
	GoodsName string  `json:"goodsName" dc:"下单时商品名"`
	Cover     string  `json:"cover"     dc:"下单时封面图"`
	Price     float64 `json:"price"     dc:"下单时单价"`
	Quantity  uint    `json:"quantity"  dc:"数量"`
	Subtotal  float64 `json:"subtotal"  dc:"小计"`
}

// OrderDetailRes 订单详情，含明细与下单人，客服核对时一眼看清买了什么
type OrderDetailRes struct {
	ID              uint64           `json:"id"              dc:"订单ID"`
	OrderNo         string           `json:"orderNo"         dc:"订单号"`
	UserID          uint64           `json:"userId"          dc:"下单用户ID"`
	Username        string           `json:"username"        dc:"下单用户"`
	TotalAmount     float64          `json:"totalAmount"     dc:"订单金额"`
	Status          uint             `json:"status"          dc:"状态"`
	StatusText      string           `json:"statusText"      dc:"状态文案"`
	Quantity        uint             `json:"quantity"        dc:"商品总件数"`
	ReceiverName    string           `json:"receiverName"    dc:"收货人"`
	ReceiverPhone   string           `json:"receiverPhone"   dc:"收货电话"`
	ReceiverAddress string           `json:"receiverAddress" dc:"收货地址"`
	Remark          string           `json:"remark"          dc:"买家备注"`
	CreatedAt       string           `json:"createdAt"       dc:"下单时间"`
	PaidAt          string           `json:"paidAt"          dc:"付款时间"`
	Items           []OrderGoodsItem `json:"items"           dc:"订单明细"`
}

// OrderShipReq 发货，已付款 → 已发货
type OrderShipReq struct {
	g.Meta `path:"/order/ship" method:"post" summary:"订单发货" tags:"后台-订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// OrderShipRes 操作结果
type OrderShipRes struct {
	ID         uint64 `json:"id"         dc:"订单ID"`
	Status     uint   `json:"status"     dc:"当前状态"`
	StatusText string `json:"statusText" dc:"状态文案"`
}

// OrderFinishReq 确认完成，已发货 → 已完成
type OrderFinishReq struct {
	g.Meta `path:"/order/finish" method:"post" summary:"订单完成" tags:"后台-订单"`
	ID     uint64 `json:"id" v:"required|min:1#请提供订单ID|订单ID不合法" dc:"订单ID"`
}

// OrderFinishRes 操作结果
type OrderFinishRes struct {
	ID         uint64 `json:"id"         dc:"订单ID"`
	Status     uint   `json:"status"     dc:"当前状态"`
	StatusText string `json:"statusText" dc:"状态文案"`
}

// ---------------------------------------------------------------------------
// 分类管理
// ---------------------------------------------------------------------------

// CategoryListReq 后台分类列表
type CategoryListReq struct {
	g.Meta `path:"/category/list" method:"get" summary:"分类列表（含停用）" tags:"后台-分类"`
}

// CategoryItem 后台分类条目
type CategoryItem struct {
	ID         uint64 `json:"id"         dc:"分类ID"`
	ParentID   uint64 `json:"parentId"   dc:"父分类ID，0为顶级"`
	ParentName string `json:"parentName" dc:"父分类名称，顶级为空"`
	Name       string `json:"name"       dc:"分类名称"`
	Sort       int    `json:"sort"       dc:"排序值，越小越靠前"`
	Status     uint   `json:"status"     dc:"状态：1启用 0停用"`
	GoodsCount int    `json:"goodsCount" dc:"商品数，顶级分类含其子分类"`
	Children   int    `json:"children"   dc:"直接子分类数"`
	CreatedAt  string `json:"createdAt"  dc:"创建时间"`
}

// CategoryListRes 平铺返回，已按「顶级 → 子级」排好序
type CategoryListRes struct {
	List []CategoryItem `json:"list" dc:"分类列表"`
}

// CategorySaveReq 新建或编辑分类。仅支持两级，所以父分类必须是顶级分类。
type CategorySaveReq struct {
	g.Meta   `path:"/category/save" method:"post" summary:"保存分类" tags:"后台-分类"`
	ID       uint64 `json:"id"       dc:"分类ID，0 表示新建"`
	ParentID uint64 `json:"parentId" d:"0" dc:"父分类ID，0 为顶级"`
	Name     string `json:"name"     v:"required|length:2,30#请输入分类名称|分类名称长度需在2-30之间" dc:"分类名称"`
	Sort     int    `json:"sort"     d:"0" v:"between:0,9999#排序值需在0-9999之间"                  dc:"排序值，越小越靠前"`
	Status   uint   `json:"status"   d:"1" v:"in:0,1#状态只能是0或1"                             dc:"状态：1启用 0停用"`
}

// CategorySaveRes 保存结果
type CategorySaveRes struct {
	ID uint64 `json:"id" dc:"分类ID"`
}

// CategoryDeleteReq 删除分类。存在子分类或商品时会被拒绝。
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"post" summary:"删除分类" tags:"后台-分类"`
	ID     uint64 `json:"id" v:"required|min:1#请提供分类ID|分类ID不合法" dc:"分类ID"`
}

// CategoryDeleteRes 删除结果
type CategoryDeleteRes struct{}

// ---------------------------------------------------------------------------
// 用户管理
// ---------------------------------------------------------------------------

// UserListReq 后台用户列表
type UserListReq struct {
	g.Meta   `path:"/user/list" method:"get" summary:"用户列表" tags:"后台-用户"`
	Page     int    `json:"page"     d:"1"  v:"min:1#页码最小为1"                  dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"between:1,100#每页数量需在1-100之间" dc:"每页数量"`
	Keyword  string `json:"keyword"                                            dc:"按用户名/昵称/邮箱/手机号搜索"`
	Role     int    `json:"role"     d:"-1" v:"in:-1,0,1#角色不合法"              dc:"按角色筛选，-1 全部"`
	Status   int    `json:"status"   d:"-1" v:"in:-1,0,1#状态不合法"              dc:"按状态筛选，-1 全部"`
}

// UserItem 后台用户条目
type UserItem struct {
	ID         uint64  `json:"id"         dc:"用户ID"`
	Username   string  `json:"username"   dc:"用户名"`
	Nickname   string  `json:"nickname"   dc:"昵称"`
	Email      string  `json:"email"      dc:"邮箱"`
	Phone      string  `json:"phone"      dc:"手机号"`
	Role       uint    `json:"role"       dc:"角色：0普通用户 1管理员"`
	Status     uint    `json:"status"     dc:"状态：1正常 0禁用"`
	OrderCount int     `json:"orderCount" dc:"订单总数"`
	TotalSpent float64 `json:"totalSpent" dc:"累计消费（已付款及之后，不含待付款与已取消）"`
	CreatedAt  string  `json:"createdAt"  dc:"注册时间"`
}

// UserListRes 分页结果
type UserListRes struct {
	List     []UserItem `json:"list"     dc:"用户列表"`
	Total    int        `json:"total"    dc:"总条数"`
	Page     int        `json:"page"     dc:"当前页"`
	PageSize int        `json:"pageSize" dc:"每页数量"`
}

// UserSaveReq 编辑用户资料 / 角色 / 状态。
// 不能降级或禁用当前登录账号，也不能让系统失去最后一名可用管理员。
type UserSaveReq struct {
	g.Meta   `path:"/user/save" method:"post" summary:"编辑用户" tags:"后台-用户"`
	ID       uint64 `json:"id"       v:"required|min:1#请提供用户ID|用户ID不合法"          dc:"用户ID"`
	Nickname string `json:"nickname" v:"max-length:20#昵称不能超过20个字符"             dc:"昵称，留空则沿用用户名"`
	Email    string `json:"email"    v:"email#邮箱格式不正确"                        dc:"邮箱，留空表示不填"`
	Phone    string `json:"phone"    v:"phone-loose#手机号格式不正确"                  dc:"手机号，留空表示不填"`
	Role     uint   `json:"role"     d:"0" v:"in:0,1#角色只能是0或1"                  dc:"角色：0普通用户 1管理员"`
	Status   uint   `json:"status"   d:"1" v:"in:0,1#状态只能是0或1"                  dc:"状态：1正常 0禁用"`
}

// UserCreateReq 后台直接新建用户。管理员设置初始密码，用户后续可自行重置。
type UserCreateReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"新建用户" tags:"后台-用户"`
	Username string `json:"username" v:"required|length:3,20#请输入用户名|用户名长度需在3-20之间" dc:"用户名，登录用，不可重复"`
	Password string `json:"password" v:"required|length:6,32#请输入初始密码|密码长度需在6-32之间" dc:"初始密码"`
	Nickname string `json:"nickname" v:"max-length:20#昵称不能超过20个字符"                  dc:"昵称，留空则沿用用户名"`
	Email    string `json:"email"    v:"email#邮箱格式不正确"                             dc:"邮箱，可为空但不可重复"`
	Phone    string `json:"phone"    v:"phone-loose#手机号格式不正确"                       dc:"手机号，可为空但不可重复"`
	Role     uint   `json:"role"     d:"0" v:"in:0,1#角色只能是0或1"                       dc:"角色：0普通用户 1管理员"`
	Status   uint   `json:"status"   d:"1" v:"in:0,1#状态只能是0或1"                       dc:"状态：1正常 0禁用"`
}

// UserCreateRes 新建结果
type UserCreateRes struct {
	ID uint64 `json:"id" dc:"用户ID"`
}

// UserSaveRes 保存结果
type UserSaveRes struct {
	ID     uint64 `json:"id"     dc:"用户ID"`
	Role   uint   `json:"role"   dc:"当前角色"`
	Status uint   `json:"status" dc:"当前状态"`
}

// UserResetPasswordReq 重置用户密码。管理员无需知道原密码。
type UserResetPasswordReq struct {
	g.Meta   `path:"/user/reset-password" method:"post" summary:"重置密码" tags:"后台-用户"`
	ID       uint64 `json:"id"       v:"required|min:1#请提供用户ID|用户ID不合法"           dc:"用户ID"`
	Password string `json:"password" v:"required|length:6,32#请输入新密码|密码长度需在6-32之间" dc:"新密码"`
}

// UserResetPasswordRes 重置结果
type UserResetPasswordRes struct{}
