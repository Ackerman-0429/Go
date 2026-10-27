// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CartDao is the data access object for the table cart.
type CartDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CartColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CartColumns defines and stores column names for the table cart.
type CartColumns struct {
	Id        string // 购物车项ID
	UserId    string // 用户ID
	GoodsId   string // 商品ID
	Quantity  string // 数量
	CreatedAt string //
	UpdatedAt string //
}

// cartColumns holds the columns for the table cart.
var cartColumns = CartColumns{
	Id:        "id",
	UserId:    "user_id",
	GoodsId:   "goods_id",
	Quantity:  "quantity",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCartDao creates and returns a new DAO object for table data access.
func NewCartDao(handlers ...gdb.ModelHandler) *CartDao {
	return &CartDao{
		group:    "default",
		table:    "cart",
		columns:  cartColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CartDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CartDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CartDao) Columns() CartColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CartDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CartDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CartDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
