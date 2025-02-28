package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/model"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
)

type MarkOrderCanceledService struct {
	ctx context.Context
} // NewMarkOrderCanceledService new MarkOrderCanceledService
func NewMarkOrderCanceledService(ctx context.Context) *MarkOrderCanceledService {
	return &MarkOrderCanceledService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderCanceledService) Run(req *order.MarkOrderCanceledReq) (resp *order.MarkOrderCanceledResp, err error) {
	// Finish your business logic.
	err = model.MarkOrderCanceled(s.ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
