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
	db.AutoMigrate(&models.Order{}, &models.Product{}, &models.Category{})
	initData(db)

	queryNestedPreload(db)
}

func initData(db *gorm.DB) {

	// 分类
	electronics := models.Category{Name: "电子产品"}
	clothes := models.Category{Name: "服装"}
	db.Create(&electronics)
	db.Create(&clothes)

	// 商品
	phone := models.Product{Name: "手机", CategoryID: electronics.ID}
	shirt := models.Product{Name: "衬衫", CategoryID: clothes.ID}
	db.Create(&phone)
	db.Create(&shirt)

	// 用户 & 订单（绑定商品）
	users := []models.User{
		{
			Name:  "Alice",
			Age:   50,
			Email: "2234112@qq.com",
			Orders: []models.Order{
				{Amount: 100, ProductID: phone.ID},
				{Amount: 150, ProductID: shirt.ID},
			},
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}

func queryNestedPreload(db *gorm.DB) {
	var users []models.User

	// 多层级预加载
	err := db.
		Preload("Orders.Product.Category").
		Find(&users).Error
	if err != nil {
		panic(err)
	}

	for _, u := range users {
		fmt.Println("用户:", u.Name)
		for _, o := range u.Orders {
			fmt.Printf("  订单金额: %.2f, 商品: %s, 分类: %s\n",
				o.Amount, o.Product.Name, o.Product.Category.Name)
		}
	}
}
