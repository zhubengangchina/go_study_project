package models

type Course struct {
	ID       uint
	Title    string
	Students []Student `gorm:"many2many:student_courses"`
}
