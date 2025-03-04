package service

import (
	"context"
	"testing"

	auth "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/auth"
	"github.com/joho/godotenv"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	// 校验身份令牌
	req := &auth.VerifyTokenReq{
		Token: "",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 校验身份令牌
	req = &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjEsInJvbGUiOiJhZG1pbiIsImlzcyI6ImdvbWFsbC1kZW1vIiwiZXhwIjoxNzQxMTMwNTE3LCJuYmYiOjE3NDEwODczMTcsImlhdCI6MTc0MTA4NzMxN30.0R4oXN_Qu423RpFqCTRsnCSQUcycUJ-zibXFeDxHca4",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
