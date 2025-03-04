package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/model"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// Finish your business logic.
	err = model.UpdateOrder(s.ctx, mysql.DB, model.Order{
		OrderId:      req.OrderId,
		UserCurrency: req.UserCurrency,
		Consignee: model.Consignee{
			StreetAddress: req.Address.StreetAddress,
			State:         req.Address.State,
			City:          req.Address.City,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
	})
	return
}
