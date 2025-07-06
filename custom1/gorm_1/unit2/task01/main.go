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
		panic("连接数据库失败:" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层连接失败")
	}
	err = sqlDB.Ping()
	if err != nil {
		panic("数据库 Ping 失败")
	}
	fmt.Println("数据库连接成功！")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}
	fmt.Println("User 表创建成功！")

	//插入一条数据
	// user := models.User{Name: "zhoudan", Email: "alice@example.com", Age: 25}
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	panic("插入失败: " + result.Error.Error())
	// }
	// fmt.Printf("✅ 插入成功：新用户 ID = %d\n", user.ID)

	// db.Model(&user).Update("Age", 20)
	// fmt.Printf("🕒 用户更新时间：%s\n", user.UpdatedAt)
	// db.First(&user, user.ID)
	// fmt.Println("🕒 再次查询后更新时间:", user.UpdatedAt)

	// bod := models.User{Name: "Bod", Age: 20, Email: "12141@qq.com"}
	// result := db.Create(&bod)
	// if result.Error != nil {
	// 	panic("插入bod失败")
	// }
	// fmt.Printf("插入成功用户ID = %d\n", bod.ID)

	//更新用户字段信息
	// bod := models.User{}
	// db.First(&bod, 1)

	// bod.Age = 17
	// db.Save(bod)
	// fmt.Printf("🕒 用户更新时间：%s\n", bod.UpdatedAt)
	// //需要再次查询更新时间
	// db.First(&bod, bod.ID)
	// fmt.Println("🕒 再次查询后更新时间:", bod.UpdatedAt)

	//email 不设置值
	// bod := models.User{Age: 10, Name: "12414", Phone: "12414"}
	// db.Create(&bod)

	//批量插入数据

	// bods1 := []models.User{{Name: "ZBG", Age: 19, Email: "14112@qq.com", Phone: "24141"}, {Name: "ZBG1", Age: 19, Email: "141112@qq.com", Phone: "24141"}}

	// result := db.Create(&bods1)
	// if result.Error != nil {
	// 	panic("批量插入失败")
	// }
	// fmt.Print("批量更新成功：", result.Row())
}
