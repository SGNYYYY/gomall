package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/auth/intra/jwt"
	auth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	stringToken, err := jwt.GenerateToken(req.UserId, req.Role)
	if err != nil {
		return nil, err
	}
	return &auth.DeliveryResp{
		Token: stringToken,
	}, nil
}
