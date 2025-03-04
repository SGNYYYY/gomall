package service

import (
	"context"
	"testing"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestCreateProduct_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewCreateProductService(ctx)
	// init req and assert value

	// 创建一个带新建的分类的商品
	req := &product.CreateProductReq{
		Product: &product.ProductCreate{
			Name:        "hamberger",
			Description: "food",
			Picture:     "/static/image/hamberger.jpg",
			Price:       6,
			Categories:  []string{"food"},
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 创建一个已有分类的商品
	req = &product.CreateProductReq{
		Product: &product.ProductCreate{
			Name:        "hamberger toy",
			Description: "Sticker",
			Picture:     "/static/image/hamberger.jpg",
			Price:       6,
			Categories:  []string{"Sticker"},
		},
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
