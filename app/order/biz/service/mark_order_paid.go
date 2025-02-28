package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/model"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	err = model.MarkOrderPaid(s.ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
