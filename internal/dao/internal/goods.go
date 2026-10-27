// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsDao is the data access object for the table goods.
type GoodsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GoodsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GoodsColumns defines and stores column names for the table goods.
type GoodsColumns struct {
	Id         string // 商品ID
	CategoryId string // 所属分类ID
	Name       string // 商品名称
	Subtitle   string // 副标题/卖点
	Price      string // 售价，必须用DECIMAL
	Stock      string // 库存
	Cover      string // 封面图URL
	Images     string // 多图URL，JSON数组
	Detail     string // 商品详情
	Sales      string // 销量
	Status     string // 状态：1上架 0下架
	CreatedAt  string //
	UpdatedAt  string //
}

// goodsColumns holds the columns for the table goods.
var goodsColumns = GoodsColumns{
	Id:         "id",
	CategoryId: "category_id",
	Name:       "name",
	Subtitle:   "subtitle",
	Price:      "price",
	Stock:      "stock",
	Cover:      "cover",
	Images:     "images",
	Detail:     "detail",
	Sales:      "sales",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewGoodsDao creates and returns a new DAO object for table data access.
func NewGoodsDao(handlers ...gdb.ModelHandler) *GoodsDao {
	return &GoodsDao{
		group:    "default",
		table:    "goods",
		columns:  goodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GoodsDao) Columns() GoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
