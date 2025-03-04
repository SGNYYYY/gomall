package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
)

func TestUpdateOrder_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewUpdateOrderService(ctx)
	// init req and assert value

	req := &order.UpdateOrderReq{
		OrderId:      "65a00156-f5d4-11ef-a7dc-4a2bf72c5fbe",
		UserCurrency: "CNY",
		Address: &order.Address{
			StreetAddress: "新港东路",
			State:         "广东省",
			City:          "广州市",
			Country:       "中国",
			ZipCode:       364000,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
