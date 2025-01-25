package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex"`
	PasswordHashed string `gorm:"not null;type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}
