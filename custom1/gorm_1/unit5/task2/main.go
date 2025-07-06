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

	db.AutoMigrate(&models.User{}, &models.Identity{})

	user := models.User{Name: "zhoudan", Email: "alice5@example.com", Age: 25}
	db.Create(&user)

	identity := models.Identity{UserID: user.ID, Number: "1234567890"}
	db.Create(&identity)

	var result models.User
	db.Preload("Identity").Find(&result, user.ID)
	fmt.Printf("用户：%s，身份证号：%s\n", result.Name, result.Identity.Number)
}
