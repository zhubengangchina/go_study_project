package main

import (
	"fmt"
	"go_study_project/custom1/gorm_1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	db.AutoMigrate(&models.Student{}, &models.Course{})

	//创建课程
	// course1 := models.Course{Title: "数学"}
	// course2 := models.Course{Title: "语文"}
	// db.Create(&course1)
	// db.Create(&course2)

	// //创建学生并关联课程
	// student1 := models.Student{
	// 	Name:    "小三",
	// 	Courses: []models.Course{course1, course2},
	// }
	// db.Create(&student1)

	var student models.Student

	db.Preload("Courses").First(&student)
	fmt.Printf("学生：%s\n", student.Name)
	for _, course := range student.Courses {
		fmt.Println("课程：", course.Title)
	}
}
