package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"` //主键
	Name      string `gorm:"size:100; not null"`
	Email     string `gorm:"unique;not null"`
	Age       int
	Phone     string `gorm:"size:10;default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` //启用软删除
	Orders    []Order
	Identity  Identity
	Balance   float64
}

func (User) TableName() string {
	return "my_users"
}
