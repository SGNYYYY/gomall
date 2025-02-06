package service

import (
	"context"

	cart "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/cart"
	"github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpccart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
