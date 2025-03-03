package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categoreis []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return product, err
}

func (p ProductQuery) GetByIds(productIdList []int) (product []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&product, productIdList).Error
	return product, err
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%").Error
	return products, err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cacheResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)

	err = func() error {
		if err1 := cacheResult.Err(); err1 != nil {
			return err1
		}
		cachedResultByte, err2 := cacheResult.Bytes()
		if err2 != nil {
			return err2
		}
		err3 := json.Unmarshal(cachedResultByte, &product)
		if err3 != nil {
			return err3
		}
		return nil
	}()
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func (c CachedProductQuery) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q)
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

type ProductMutation struct {
	ctx context.Context
	db  *gorm.DB
}
