package service

import (
	"context"

	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/SGNYYYY/gomall/app/frontend/utils"
	rpccart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.EmptyCart(h.Context, &rpccart.EmptyCartReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	return
}
