package service

import (
	"context"

	common "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)
	items := []map[string]any{
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
		{"Name": "Hamberger", "Price": 36, "Picture": "/static/image/hamberger.jpg"},
	}
	resp["title"] = "Hot Sales"
	resp["items"] = items
	return resp, nil
}
