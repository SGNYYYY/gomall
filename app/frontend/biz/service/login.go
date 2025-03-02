package service

import (
	"context"
	"fmt"

	auth "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	rpcauth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
	rpcuser "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
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

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	loginrResp, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}

	stringToken, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{
		UserId: loginrResp.UserId,
		Role:   loginrResp.Role,
	})
	if err != nil {
		return "", err
	}
	fmt.Println("token:", stringToken.Token)
	session := sessions.Default(h.RequestContext)
	session.Set("token", stringToken.Token)
	err = session.Save()
	if err != nil {
		return "", err
	}
	var redirect string
	if req.Next != "" {
		redirect = req.Next
	} else {
		redirect = "/"
	}
	return redirect, nil
}
