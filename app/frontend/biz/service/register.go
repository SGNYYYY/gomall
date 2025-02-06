package service

import (
	"context"

	auth "github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/SGNYYYY/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/SGNYYYY/gomall/app/frontend/infra/rpc"
	rpcauth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
	rpcuser "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	registerResp, err := rpc.UserClient.Register(h.Context, &rpcuser.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return nil, err
	}

	stringToken, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{
		UserId: registerResp.UserId,
		Role:   "user",
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("token", stringToken.Token)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
