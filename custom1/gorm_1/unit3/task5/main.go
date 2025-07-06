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

	//硬删除（Delete）
	// db.Delete(&models.User{}, 1)
	// //如果模型没有 DeletedAt 字段，则直接删除  如果模型有 DeletedAt，则默认是软删除

	// //软删除（推荐）
	// var user models.User
	// db.Find(&user, 2)
	// db.Debug().Delete(&user)

	// //如何查询已删除的数据？
	// var users []models.User
	// db.Unscoped().Where("deleted_at is not null").Find(&users)
	// for _, v := range users {
	// 	fmt.Println(v.Name)
	// }

	// //恢复被软删除的记录
	// db.Unscoped().Model(&models.User{}).Where("id = ? ", 1).Update("deleted_at", nil)

	// //强制硬删除（跳过软删除）
	// db.Unscoped().Delete(&models.User{}, 1)

	//练习
	/*
		删除 ID 为 2 的用户（使用软删除）；

		再次查询该用户，验证是否还能查到（普通查询）；

		使用 Unscoped() 查询被删除的用户；

		恢复该用户；

		强制删除 ID 为 3 的用户（真正从数据库移除）；
	*/

	//var user models.User

	// db.Find(&user, 30)
	// //db.Delete(&user)
	// fmt.Println(user.Name)

	var users []models.User
	db.Unscoped().Where("deleted_at is not null").Find(&users)
	fmt.Println(len(users))
	db.Unscoped().Model(&models.User{}).Where("id = ?", 30).Update("deleted_at", nil)
	fmt.Println(len(users))

	db.Unscoped().Delete(&models.User{}, 30)
}
