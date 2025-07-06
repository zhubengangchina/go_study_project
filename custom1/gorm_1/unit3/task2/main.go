package main

import (
	"fmt"
	"go_study_project/custom1/gorm_1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//3.2.2：根据条件查询单个用户（Where + First）

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	var user models.User
	// result := db.Where("email = ? ", "alice@example.com").First(&user)
	// if result.Error != nil {
	// 	fmt.Println("❌ 用户不存在")
	// } else {
	// 	fmt.Println("✅ 找到用户：", user)
	// }

	//也可以用 struct 或 map 作为条件

	// result1 := db.Where(&models.User{Name: "ZBG"}).First(&user)
	// if result1.Error != nil {
	// 	fmt.Println("❌ 用户不存在")
	// } else {
	// 	fmt.Println("✅ 找到用户：", user)
	// }

	result2 := db.Where(map[string]interface{}{"name": "ZBG1"}).First(&user)
	if result2.Error != nil {
		fmt.Println("❌ 用户不存在")
	} else {
		fmt.Println("✅ 找到用户：", user)
	}

}
