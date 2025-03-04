package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/product/biz/dal/redis"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestGetProducts_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	redis.Init()
	ctx := context.Background()
	s := NewGetProductsService(ctx)
	// init req and assert value

	req := &product.GetProductsReq{
		Ids: []uint32{1, 2, 3},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
