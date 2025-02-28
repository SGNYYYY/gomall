package order

import (
	"context"

	"github.com/SGNYYYY/gomall/app/frontend/biz/service"
	"github.com/SGNYYYY/gomall/app/frontend/biz/utils"
	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "order", utils.WrapResponse(ctx, c, resp))
}

// CancelOrder .
// @router /order/cancel [POST]
func CancelOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CancelOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewCancelOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	var req1 common.Empty
	resp, err := service.NewOrderListService(ctx, c).Run(&req1)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "order", utils.WrapResponse(ctx, c, resp))
}
