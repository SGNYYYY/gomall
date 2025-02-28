package service

import (
	"context"

	payment "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/payment"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpcorder "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	rpcpayment "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type PaymentWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentWaitingService(Context context.Context, RequestContext *app.RequestContext) *PaymentWaitingService {
	return &PaymentWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentWaitingService) Run(req *payment.PaymentReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.PaymentClient.Charge(h.Context, &rpcpayment.ChargeReq{
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
		OrderId: req.OrderId,
		UserId:  userId,
		Amount:  req.Total,
	})
	if err != nil {
		return nil, err
	}
	_, err = rpc.OrderClient.MarkOrderPaid(h.Context, &rpcorder.MarkOrderPaidReq{OrderId: req.OrderId, UserId: userId})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/payment/result",
	}, nil
}
