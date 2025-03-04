package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/product/biz/dal/redis"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestListProducts_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	redis.Init()
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
