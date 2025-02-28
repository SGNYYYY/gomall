package service

import (
	"context"
	"time"

	"github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/payment"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	"github.com/SGNYYYY/gomall/app/frontend/types"
	rpcorder "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type PaymentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentService(Context context.Context, RequestContext *app.RequestContext) *PaymentService {
	return &PaymentService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentService) Run(req *payment.PaymentPageReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	orderResp, err := rpc.OrderClient.GetOrder(h.Context, &rpcorder.GetOrderReq{OrderId: req.OrderId})
	if err != nil {
		return nil, err
	}
	var items []types.OrderItem
	var total float32
	var ids []uint32
	if len(orderResp.Order.OrderItems) > 0 {
		for _, vv := range orderResp.Order.OrderItems {
			total += vv.Cost
			i := vv.Item
			ids = append(ids, vv.Item.ProductId)
			items = append(items, types.OrderItem{
				ProductId: i.ProductId,
				Qty:       uint32(i.Quantity),
				Cost:      vv.Cost,
			})
		}
	}
	productList, err := rpc.ProductClient.GetProducts(h.Context, &rpcproduct.GetProductsReq{Ids: ids})
	if err != nil {
		return nil, err
	}
	for index, p := range productList.Products {
		i := &items[index]
		i.ProductId = p.Id
		i.ProductName = p.Name
		i.Picture = p.Picture
	}

	timeObj := time.Unix(int64(orderResp.Order.CreatedAt), 0)
	order := &types.Order{
		Cost:        total,
		Items:       items,
		CreatedDate: timeObj.Format("2006-01-02 15:04:05"),
		OrderId:     orderResp.Order.OrderId,
		Consignee:   types.Consignee{Email: orderResp.Order.Email},
		OrderState:  orderResp.Order.OrderState,
	}
	return utils.H{
		"title":    "Payment",
		"order_id": order.OrderId,
		"order":    order,
		"items":    order.Items,
		"total":    total,
	}, nil
}
