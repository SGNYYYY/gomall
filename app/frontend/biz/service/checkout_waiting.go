package service

import (
	"context"
	"strconv"

	"github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpccheckout "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	r, err := rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    userId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		// CreditCard: &rpcpayment.CreditCardInfo{
		// 	CreditCardNumber:          req.CardNum,
		// 	CreditCardExpirationYear:  req.ExpirationYear,
		// 	CreditCardExpirationMonth: req.ExpirationMonth,
		// 	CreditCardCvv:             req.Cvv,
		// },
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "Payment",
		"order_id": r.OrderId,
		"total":    strconv.FormatFloat(float64(r.Amount), 'f', 2, 64),
	}, nil
}
