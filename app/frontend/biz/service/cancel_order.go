package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	order "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/order"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpcorder "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CancelOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCancelOrderService(Context context.Context, RequestContext *app.RequestContext) *CancelOrderService {
	return &CancelOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *CancelOrderService) Run(req *order.CancelOrderReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.OrderClient.MarkOrderCanceled(h.Context, &rpcorder.MarkOrderCanceledReq{OrderId: req.OrderId, UserId: userId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
