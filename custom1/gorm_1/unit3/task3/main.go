package main

import (
	"fmt"
	"go_study_project/custom1/gorm_1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//3.2.3：查询多条记录（Find）
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}
	//var users []models.User
	// db.Where("age > ?", 19).Find(&users)
	// for _, u := range users {
	// 	fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)
	// }

	// 3.2.4：使用 Select, Order, Limit, Offset
	// 只查询 name 和 age 字段，按 age 降序
	// db.Select("name", "age").Order("age desc").Find(&users)
	// fmt.Println(users)
	// for _, u := range users {
	// 	fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)
	// }

	// 查询前 2 个用户（分页）
	// db.Limit(2).Offset(1).Find(&users)
	// for _, u := range users {
	// 	fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)
	// }

	//3.2.5：原生 SQL 查询（Raw + Scan）

	// type Result struct {
	// 	Name string
	// 	Age  int
	// }

	// var results []Result

	// db.Raw("select name ,age from my_users where age > ?", 17).Scan(&results)
	// for _, r := range results {
	// 	fmt.Println(r.Name, r.Age)
	// }

	/*
		查询年龄大于 25 的所有用户

		查询 email 为 "bob@example.com" 的用户，如果不存在打印提示

		查询数据库中前 3 个年龄最大用户（Order + Limit）

		使用原生 SQL 查询所有用户的 name 和 phone 字段
	*/
	// var users []models.User
	// db.Where("age > ? ", 20).Find(&users)
	// for _, u := range users {
	// 	fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)
	// }

	// var u models.User
	// db.Where("email = ?", "alice@example.com").First(&u)
	// fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)

	// var users []models.User
	// db.Limit(3).Offset(0).Order("age desc").Find(&users)
	// for _, u := range users {
	// 	fmt.Printf("用户：%s，年龄：%d\n", u.Name, u.Age)
	// }

	var users []models.User

	db.Raw("select name,phone from my_users").Scan(&users)

	for _, u := range users {
		fmt.Printf("用户：%s，年龄：%s\n", u.Name, u.Phone)
	}
}
