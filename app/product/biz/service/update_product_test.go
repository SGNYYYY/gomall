package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestUpdateProduct_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &product.UpdateProductReq{
		Product: &product.Product{
			Id:          10,
			Name:        "hamberger",
			Description: "food",
			Picture:     "/static/image/hamberger.jpg",
			Price:       7,
			Categories:  []string{"food"},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
