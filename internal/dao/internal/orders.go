// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrdersDao is the data access object for the table orders.
type OrdersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OrdersColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OrdersColumns defines and stores column names for the table orders.
type OrdersColumns struct {
	Id              string // 订单ID
	OrderNo         string // 订单号，对外展示
	UserId          string // 下单用户ID
	TotalAmount     string // 订单总金额
	Status          string // 状态：0待付款 1已付款 2已发货 3已完成 4已取消
	ReceiverName    string // 收货人
	ReceiverPhone   string // 收货电话
	ReceiverAddress string // 收货地址
	Remark          string // 订单备注
	PaidAt          string // 付款时间
	CreatedAt       string //
	UpdatedAt       string //
}

// ordersColumns holds the columns for the table orders.
var ordersColumns = OrdersColumns{
	Id:              "id",
	OrderNo:         "order_no",
	UserId:          "user_id",
	TotalAmount:     "total_amount",
	Status:          "status",
	ReceiverName:    "receiver_name",
	ReceiverPhone:   "receiver_phone",
	ReceiverAddress: "receiver_address",
	Remark:          "remark",
	PaidAt:          "paid_at",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao(handlers ...gdb.ModelHandler) *OrdersDao {
	return &OrdersDao{
		group:    "default",
		table:    "orders",
		columns:  ordersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrdersDao) Columns() OrdersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
