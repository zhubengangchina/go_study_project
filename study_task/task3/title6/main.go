package main

import (
	"fmt"
	"go_study_project/study_task/task3/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	//查询用户发布的文章
	posts, err := GetUserPostsWithComments(db, 1)
	if err != nil {
		log.Fatal("查询失败：", err)
	}

	for _, post := range posts {
		fmt.Printf("用户：%s 的 文章标题: %s\n", post.User.Name, post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("  评论内容: %s\n", comment.Body)
		}
	}

	//查询评论最多的文章信息
	postWithCount, err := GetMostCommentedPost(db)
	if err != nil {
		log.Fatal("查询失败：", err)
	}

	fmt.Printf("评论最多的文章：%s（共 %d 条评论）\n", postWithCount.Title, postWithCount.CommentCount)
}

func GetUserPostsWithComments(db *gorm.DB, userId uint) ([]models.Post, error) {
	var posts []models.Post
	err := db.
		Preload("Comments").
		Preload("User").
		Where("user_id = ?", userId).
		Find(&posts).Error
	return posts, err

}

type PostWithCount struct {
	models.Post
	CommentCount uint64 `gorm:"column:comment_count"`
}

func GetMostCommentedPost(db *gorm.DB) (PostWithCount, error) {
	var postWithCount PostWithCount
	err := db.Table("posts").
		Select("posts.*,count(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&postWithCount).Error
	return postWithCount, err
}
