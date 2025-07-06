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

	db.AutoMigrate(&models.User{}, &models.Order{})

	//创建用户
	user := models.User{Name: "zhoudan", Email: "alice2@example.com", Age: 25}
	db.Create(&user)

	//创建订单
	orders := []models.Order{{UserID: user.ID, Amount: 100.5}, {UserID: user.ID, Amount: 250.0}}
	db.Create(&orders)

	//查询用户及其订单
	var result models.User
	db.Preload("Orders").Find(&result, user.ID)
	fmt.Printf("用户：%s\n", result.Name)
	for _, order := range result.Orders {
		fmt.Printf("订单金额：%.2f\n", order.Amount)
	}
}
