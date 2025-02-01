package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/cart/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/cart/biz/model"
	cart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	cartList, err := model.GetCartByUserId(s.ctx, mysql.DB, req.GetUserId())
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*cart.CartItem
	for _, item := range cartList {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Qty),
		})
	}
	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.GetUserId(),
			Items:  items,
		},
	}, nil
}
