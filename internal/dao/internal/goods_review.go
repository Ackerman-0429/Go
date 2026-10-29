// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsReviewDao is the data access object for the table goods_review.
type GoodsReviewDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GoodsReviewColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GoodsReviewColumns defines and stores column names for the table goods_review.
type GoodsReviewColumns struct {
	Id          string // 评价ID
	GoodsId     string // 商品ID
	OrderId     string // 订单ID
	OrderItemId string // 订单明细ID，一条明细只能评一次
	UserId      string // 评价用户ID
	Rating      string // 评分 1-5 星
	Content     string // 评价内容
	Anonymous   string // 是否匿名：1 是 0 否
	Status      string // 状态：1 显示 0 隐藏
	CreatedAt   string //
}

// goodsReviewColumns holds the columns for the table goods_review.
var goodsReviewColumns = GoodsReviewColumns{
	Id:          "id",
	GoodsId:     "goods_id",
	OrderId:     "order_id",
	OrderItemId: "order_item_id",
	UserId:      "user_id",
	Rating:      "rating",
	Content:     "content",
	Anonymous:   "anonymous",
	Status:      "status",
	CreatedAt:   "created_at",
}

// NewGoodsReviewDao creates and returns a new DAO object for table data access.
func NewGoodsReviewDao(handlers ...gdb.ModelHandler) *GoodsReviewDao {
	return &GoodsReviewDao{
		group:    "default",
		table:    "goods_review",
		columns:  goodsReviewColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GoodsReviewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GoodsReviewDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GoodsReviewDao) Columns() GoodsReviewColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GoodsReviewDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GoodsReviewDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *GoodsReviewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
