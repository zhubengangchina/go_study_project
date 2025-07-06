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

	var user models.User
	// // 查询主键最小的（默认按主键升序）
	// db.First(&user)
	// fmt.Println("First:", user)

	// // 不保证顺序（效率稍高）
	db.Take(&user)
	fmt.Println("Take:", user)

	// 查询主键最大的（默认按主键降序）
	// db.Last(&user)
	// fmt.Println("Last:", user)
}
