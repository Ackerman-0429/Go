package cart

import (
	"context"

	v1 "hello/api/cart/v1"
	"hello/internal/service"
)

// Delete 处理 POST /cart/delete
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	if err = service.Cart().Remove(ctx, currentUserID(ctx), req.ID); err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
