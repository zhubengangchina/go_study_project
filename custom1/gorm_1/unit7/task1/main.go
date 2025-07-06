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
	initData(db)

	queryWithPreload(db)
}

func initData(db *gorm.DB) {
	// 添加用户及订单（使用 Create 插入一对多数据）

	users := []models.User{
		{Name: "111", Email: "12411@qq.com", Age: 15, Orders: []models.Order{{Amount: 100}, {Amount: 1000}}},
		{Name: "1111", Email: "112411@qq.com", Age: 19, Orders: []models.Order{{Amount: 200}, {Amount: 2000}}},
	}

	for _, u := range users {
		db.Create(&u)
	}
}

func queryWithPreload(db *gorm.DB) {
	var users []models.User
	//核心语句 预加载order
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println("用户:", u.Name)
		for _, o := range u.Orders {
			fmt.Println("  订单金额:", o.Amount)
		}
	}

}
