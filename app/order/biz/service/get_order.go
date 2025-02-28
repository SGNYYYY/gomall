package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/model"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	// Finish your business logic.
	if req.OrderId == "" {
		return nil, kerrors.NewBizStatusError(40000, "order id is required")
	}
	o, err := model.GetById(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500001, err.Error())
	}
	var items []*order.OrderItem
	for _, oi := range o.OrderItems {
		items = append(items, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: oi.ProductId,
				Quantity:  oi.Quantity,
			},
			Cost: oi.Cost,
		})
	}
	orderResp := &order.Order{
		OrderId:      o.OrderId,
		UserId:       o.UserId,
		UserCurrency: o.UserCurrency,
		Email:        o.Consignee.Email,
		Address: &order.Address{
			StreetAddress: o.Consignee.StreetAddress,
			City:          o.Consignee.City,
			State:         o.Consignee.State,
			Country:       o.Consignee.Country,
			ZipCode:       o.Consignee.ZipCode,
		},
		OrderItems: items,
		CreatedAt:  int32(o.CreatedAt.Unix()),
		OrderState: string(o.OrderState),
	}

	return &order.GetOrderResp{Order: orderResp}, nil
}
