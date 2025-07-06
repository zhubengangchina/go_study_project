package models

type Product struct {
	ID         uint
	Name       string
	CategoryID uint
	Category   Category
}
