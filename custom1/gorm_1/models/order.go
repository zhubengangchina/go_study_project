package models

type Order struct {
	ID        uint
	UserID    uint //外键字段
	Amount    float64
	ProductID uint
	Product   Product
}
