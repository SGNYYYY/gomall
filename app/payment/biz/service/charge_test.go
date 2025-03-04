package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/payment/biz/dal/mysql"
	payment "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/payment"
	"github.com/joho/godotenv"
)

func TestCharge_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Amount: 10,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "424242424242424242",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2030,
			CreditCardExpirationMonth: 12,
		},
		OrderId: "4a640642-f919-11ef-af0d-4a2bf72c5fbe",
		UserId:  1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
