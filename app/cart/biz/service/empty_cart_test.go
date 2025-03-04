package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/cart/biz/dal/mysql"
	cart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
)

func TestEmptyCart_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")

	mysql.Init()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
