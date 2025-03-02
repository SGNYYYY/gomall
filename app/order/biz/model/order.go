package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState
}

func (o Order) TableName() string {
	return "order"
}

func GetById(ctx context.Context, db *gorm.DB, orderId string) (Order, error) {
	var order Order
	err := db.WithContext(ctx).Model(&Order{}).Where("order_id = ?", orderId).Preload("OrderItems").Find(&order).Error
	return order, err
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Model(&Order{}).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func MarkOrderPaid(ctx context.Context, db *gorm.DB, userId uint32, orderId string) error {
	return db.WithContext(ctx).Model(&Order{}).Where("user_id = ? AND order_id = ? ", userId, orderId).Update("order_state", OrderStatePaid).Error
}

func MarkOrderCanceled(ctx context.Context, db *gorm.DB, userId uint32, orderId string) error {
	return db.WithContext(ctx).Model(&Order{}).Where("user_id = ? AND order_id = ? ", userId, orderId).Update("order_state", OrderStateCanceled).Error
}

func CanceledByOrderId(ctx context.Context, db *gorm.DB, orderId string) error {
	return db.WithContext(ctx).Model(&Order{}).Where("order_id = ? ", orderId).Update("order_state", OrderStateCanceled).Error
}

func GetPendingOrders(ctx context.Context, db *gorm.DB) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Model(&Order{}).Where("order_state = ? AND created_at < ?", OrderStatePlaced, time.Now().Add(-30*time.Minute)).Find(&orders).Error
	return orders, err
}
