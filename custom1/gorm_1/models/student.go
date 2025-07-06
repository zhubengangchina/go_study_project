package models

type Student struct {
	ID      uint
	Name    string
	Courses []Course `gorm:"many2many:student_courses"`
}
