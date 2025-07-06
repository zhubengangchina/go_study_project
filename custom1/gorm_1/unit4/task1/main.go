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
	// 查询参数（模拟从接口传来的参数）
	searchName := "A"   // 模糊搜索关键词
	minAge := 20        // 最小年龄
	maxAge := 30        // 最大年龄
	page := 1           // 页码
	pageSize := 5       // 每页条数
	sortField := "age"  // 排序字段
	sortOrder := "desc" // 排序方式：asc / desc

	//分页计算
	offset := (page - 1) * pageSize
	//查询构建
	var users []models.User

	query := db.Model(&models.User{})

	//条件模糊查询
	if searchName != "" {
		query = query.Where("name like ?", "%"+searchName+"%")
	}

	//条件 年龄范围
	query = query.Where("age > ? and age < ?", minAge, maxAge)

	//排序
	query = query.Order(fmt.Sprintf("%s %s", sortField, sortOrder))

	//分页
	query = query.Limit(pageSize).Offset(offset)

	//执行查询
	if err := query.Find(&users).Error; err != nil {
		panic(err)
	}

	//打印结果
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", u.ID, u.Name, u.Age, u.Email)
	}
}

func GetUsersByPage() {

}
