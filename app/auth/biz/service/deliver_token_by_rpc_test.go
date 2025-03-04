package service

import (
	"context"
	"testing"

	auth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
	"github.com/joho/godotenv"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	// 生成令牌
	req := &auth.DeliverTokenReq{
		UserId: 1,
		Role:   "admin",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
