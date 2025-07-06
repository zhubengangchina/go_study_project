package models

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	PostCount int64
	Posts     []Post // 一对多：User → Posts
}
