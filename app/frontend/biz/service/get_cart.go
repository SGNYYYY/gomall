package service

import (
	"context"
	"strconv"

	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpccart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	var items []map[string]string
	var total float64
	for _, v := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.ProductId})
		if err != nil {
			continue
		}
		items = append(items, map[string]string{
			"Name":        productResp.Product.Name,
			"Description": productResp.Product.Description,
			"Price":       strconv.FormatFloat(float64(productResp.Product.Price), 'f', 2, 64),
			"Picture":     productResp.Product.Picture,
			"Qty":         strconv.Itoa(int(v.Quantity)),
		})
		total += float64(productResp.Product.Price) * float64(v.Quantity)
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
