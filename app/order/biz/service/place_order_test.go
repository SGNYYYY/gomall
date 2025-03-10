package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/order"
	"github.com/joho/godotenv"
)

func TestPlaceOrder_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &order.PlaceOrderReq{
		UserId:       1,
		UserCurrency: "CNY",
		Address: &order.Address{
			StreetAddress: "和平路",
			Country:       "中国",
			State:         "福建省",
			City:          "龙岩市",
			ZipCode:       364000,
		},
		Email: "sgny@demo.com",
		OrderItems: []*order.OrderItem{
			{
				Item: &cart.CartItem{
					ProductId: 1,
					Quantity:  1,
				},
				Cost: 5,
			},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
