package main

import (
	"go_study_project/study_task/task3/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	err = db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Post{})
	if err != nil {
		log.Fatal("迁移失败: ", err)
	}

	log.Println("数据库表创建完成 ✅")

	// user := models.User{Name: "Alice", Email: "alice@example.com"}
	// db.Create(&user)

	post := models.Post{Title: "第一篇文章2", Body: "内容内容2", UserID: 1}
	db.Create(&post)

	comment := models.Comment{Body: "不错的文章2！", PostID: post.ID}
	db.Create(&comment)

}
