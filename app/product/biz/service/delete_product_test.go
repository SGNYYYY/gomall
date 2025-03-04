package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestDeleteProduct_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewDeleteProductService(ctx)
	// init req and assert value

	// 删除商品
	req := &product.DeleteProductReq{
		Id: 11,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
