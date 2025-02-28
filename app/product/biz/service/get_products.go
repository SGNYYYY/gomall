package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/product/biz/model"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductsService struct {
	ctx context.Context
} // NewGetProductsService new GetProductsService
func NewGetProductsService(ctx context.Context) *GetProductsService {
	return &GetProductsService{ctx: ctx}
}

// Run create note info
func (s *GetProductsService) Run(req *product.GetProductsReq) (resp *product.GetProductsResp, err error) {
	// Finish your business logic.
	if req.Ids == nil {
		return nil, kerrors.NewBizStatusError(40000, "product id list is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	ids := make([]int, len(req.Ids))
	for i, id := range req.Ids {
		ids[i] = int(id)
	}
	var results []*product.Product
	productList, err := productQuery.GetByIds(ids)
	for _, v := range productList {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	return &product.GetProductsResp{Products: results}, err
}
