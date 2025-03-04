package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/checkout/infra/rpc"
	checkout "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/joho/godotenv"
)

func TestCheckout_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value
	rpc.InitClient()
	req := &checkout.CheckoutReq{
		UserId:    1,
		Firstname: "上官",
		Lastname:  "Noris",
		Email:     "sgny@demo.com",
		Address: &checkout.Address{
			StreetAddress: "和平路",
			Country:       "中国",
			State:         "福建省",
			City:          "龙岩市",
			ZipCode:       "364000",
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
