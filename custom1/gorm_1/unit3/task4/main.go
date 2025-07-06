package main

import (
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

	//使用 Save() 更新整条记录
	var user models.User
	db.Find(&user, 1)

	user.Age = 99
	user.Name = "Alice Updated"
	db.Save(&user)

	//使用 Update() 更新单个字段
	db.Model(&user).Where("id = ?", 1).Update("age", 35)

	//使用 Updates() 更新多个字段
	// 使用 map（零值不会被忽略）
	db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{
		"name": "new name",
		"age":  40,
	})
	//使用 struct（零值会被忽略！）
	db.Model(&user).Where("id = ?", 1).Updates(models.User{Name: "n", Age: 0})
	//// age=0 不会被更新，因为是零值

	//如何更新零值字段？（强制更新）
	//db.Model(&user).Where("id = ?", 1).Select("age").Updates(models.User{Age: 0})

	db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{
		"age": 0,
	})

	//使用原生 SQL 更新
	db.Exec("update my_users set name = ? , age = ?", "bbb", 12)
	//不会自动更新 UpdatedAt  适合复杂或批量处理

	//练习
	/*
		将 ID=1 的用户 Name 改为 "Alice Updated"，并打印更新后的结果；

		将 ID=2 的用户 Age 改为 0，验证是否成功更新；

		批量将年龄小于 30 的用户的 Phone 设置为 "unknown"；

		使用原生 SQL 将 ID=3 的用户 Email 改为 "deleted@example.com"；
	*/

	var myUser models.User
	// db.Find(&myUser, 1)

	// myUser.Name = "Alice Updated"
	// db.Save(&myUser)
	// db.Find(&user, user.ID)
	// fmt.Printf("名称：%s\n", myUser.Name)

	db.Model(&myUser).Where("id = ?", 2).Updates(map[string]interface{}{
		"age": 0,
	})

	db.Exec("update my_users set phone = ? where age < ?", "unknown", 30)

	db.Exec("update my_users set email = ? where id = ?", "deleted@example.com", 30)
}
