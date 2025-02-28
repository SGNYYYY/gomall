package payment

import (
	"context"

	"github.com/SGNYYYY/gomall/app/frontend/biz/service"
	"github.com/SGNYYYY/gomall/app/frontend/biz/utils"
	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	payment "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Payment .
// @router /payment [GET]
func Payment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.PaymentPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPaymentService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "payment", utils.WrapResponse(ctx, c, resp))
}

// PaymentWaiting .
// @router /payment/waiting [POST]
func PaymentWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.PaymentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewPaymentWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WrapResponse(ctx, c, resp))
}

// PaymentResult .
// @router /payment/result [GET]
func PaymentResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewPaymentResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WrapResponse(ctx, c, resp))
}
