package service

import (
	"context"

	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type PaymentResultService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentResultService(Context context.Context, RequestContext *app.RequestContext) *PaymentResultService {
	return &PaymentResultService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentResultService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return utils.H{}, nil
}
