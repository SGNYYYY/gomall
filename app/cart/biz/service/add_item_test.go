package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/cart/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/cart/infra/rpc"
	cart "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
)

func TestAddItem_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")

	mysql.Init()
	rpc.InitClient()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	// 加入一个不存在的商品
	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 99,
			Quantity:  1,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 加入商品
	req = &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
