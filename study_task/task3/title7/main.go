package main

import (
	"go_study_project/study_task/task3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
*/
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	// 创建用户
	user := models.User{Name: "张三"}
	db.Create(&user)

	// 创建文章（会触发 AfterCreate）
	post := models.Post{Title: "GORM 教程", UserID: user.ID}
	db.Create(&post)

	// 添加评论
	comment := models.Comment{Body: "写得不错", PostID: post.ID}
	db.Create(&comment)

	// 删除评论（会触发 AfterDelete）
	db.Delete(&comment)

}
