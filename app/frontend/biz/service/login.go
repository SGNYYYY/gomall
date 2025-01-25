package service

import (
	"context"

	auth "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// TODO User SVC API
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return "", err
	}
	if req.Next != "" {
		redirect = req.Next
	} else {
		redirect = "/"
	}
	return redirect, nil
}
