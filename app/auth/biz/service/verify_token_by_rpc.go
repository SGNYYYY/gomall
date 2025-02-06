package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/auth/intra/jwt"
	auth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	jwtResp, err := jwt.VerifyToken(req.Token)
	if err != nil {
		return nil, err
	}
	if jwtResp.Token == "" {
		return &auth.VerifyResp{
			UserId: jwtResp.UserId,
			Role:   jwtResp.Role,
			Res:    false,
			Token:  jwtResp.Token,
		}, nil
	}
	return &auth.VerifyResp{
		UserId: jwtResp.UserId,
		Role:   jwtResp.Role,
		Res:    true,
		Token:  jwtResp.Token,
	}, nil
}
