package models

type Identity struct {
	ID     uint
	UserID uint   //外键 指向user
	Number string //身份证
}
