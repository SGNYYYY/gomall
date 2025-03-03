package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/checkout/biz/dal/redis"
	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/product/biz/model"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewBizStatusError(40000, "product id is required")
	}
	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)
	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
		},
	}, nil
}
