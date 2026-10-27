// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryDao is the data access object for the table category.
type CategoryDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CategoryColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CategoryColumns defines and stores column names for the table category.
type CategoryColumns struct {
	Id        string // 分类ID
	ParentId  string // 父分类ID，0表示顶级
	Name      string // 分类名称
	Sort      string // 排序值，越小越靠前
	Status    string // 状态：1启用 0停用
	CreatedAt string //
	UpdatedAt string //
}

// categoryColumns holds the columns for the table category.
var categoryColumns = CategoryColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCategoryDao creates and returns a new DAO object for table data access.
func NewCategoryDao(handlers ...gdb.ModelHandler) *CategoryDao {
	return &CategoryDao{
		group:    "default",
		table:    "category",
		columns:  categoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CategoryDao) Columns() CategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CategoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
