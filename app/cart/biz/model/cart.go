// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserId    uint32 `json:"user_id"`
	ProductId uint32 `json:"product_id"`
	Qty       uint32 `json:"qty"`
}

func (c Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	return db.WithContext(ctx).Model(&Cart{}).Create(item).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("需要user id")
	}

	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	if userId == 0 {
		return nil, errors.New("需要user id")
	}
	var cartList []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}
