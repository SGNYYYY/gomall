package service

import (
	"context"

	"github.com/SGNYYYY/gomall/app/product/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/product/biz/model"
	product "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	productMutation := model.NewProductMutation(s.ctx, mysql.DB)
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	categoryMutation := model.NewCategoryMutation(s.ctx, mysql.DB)
	var categories []model.Category

	// 遍历分类，看是否有新的类别
	for _, c := range req.Product.Categories {
		categoryResult, err := categoryQuery.GetCategoryByName(c)
		if err != nil {
			newCategory := &model.Category{
				Name:        c,
				Description: "",
			}
			// 如果有新的类别，则创建一个类别
			err1 := categoryMutation.Create(newCategory)
			if err1 != nil {
				return nil, err
			}
			categoryResult = *newCategory
		}
		categories = append(categories, categoryResult)
	}
	// 更新商品
	err = productMutation.Update(&model.Product{
		Base:        model.Base{ID: int(req.Product.Id)},
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Picture:     req.Product.Picture,
		Price:       req.Product.Price,
		Categoreis:  categories,
	})
	if err != nil {
		return nil, err
	}
	return
}
