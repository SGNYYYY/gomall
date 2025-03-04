package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string   `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string   `gorm:"not null;type:varchar(255)"`
	Role           UserRole `gorm:"default:user"`
	Status         State
}
type State string

const (
	Enabled  State = "enabled"
	Disabled State = "disabled"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Model(&User{}).Create(&user).Error
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func GetById(db *gorm.DB, ctx context.Context, userId uint) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).First(&user, userId).Error
	return
}

func DeleteUser(db *gorm.DB, ctx context.Context, userId uint) error {
	return db.WithContext(ctx).Model(&User{}).Delete(&User{}, userId).Error
}

func GetUsers(db *gorm.DB, ctx context.Context) (users []*User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Find(&users).Error
	return
}

func UpdatePassword(db *gorm.DB, ctx context.Context, userId uint, passwordHashed string) error {
	return db.WithContext(ctx).Model(&User{}).Where("ID = ?", userId).Update("password_hashed", passwordHashed).Error
}

func UpdateRole(db *gorm.DB, ctx context.Context, userId uint, role UserRole) error {
	return db.WithContext(ctx).Model(&User{}).Where("ID = ?", userId).Update("role", role).Error
}

func ActivateUser(db *gorm.DB, ctx context.Context, userId uint) error {
	return db.WithContext(ctx).Model(&User{}).Where("ID = ?", userId).Update("status", Enabled).Error
}

func DeactivateUser(db *gorm.DB, ctx context.Context, userId uint) error {
	return db.WithContext(ctx).Model(&User{}).Where("ID = ?", userId).Update("status", Disabled).Error
}
