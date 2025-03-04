package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
)

func TestMarkOrderPaid_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)
	// init req and assert value

	req := &order.MarkOrderPaidReq{
		UserId:  1,
		OrderId: "fecfc600-f916-11ef-b7c1-4a2bf72c5fbe",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
