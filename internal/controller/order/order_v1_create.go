package order

import (
	"context"

	v1 "goshop/api/order/v1"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Create 处理 POST /order/create
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	out, err := service.Order().Create(ctx, model.OrderCreateInput{
		UserID:          currentUserID(ctx),
		ReceiverName:    req.ReceiverName,
		ReceiverPhone:   req.ReceiverPhone,
		ReceiverAddress: req.ReceiverAddress,
		Remark:          req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{
		OrderID:     out.OrderID,
		OrderNo:     out.OrderNo,
		TotalAmount: out.TotalAmount,
		Quantity:    out.Quantity,
	}, nil
}
