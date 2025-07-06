package models

import (
	"gorm.io/gorm"
)

type Post struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	Body          string
	UserID        uint      //外键
	User          User      // 属于哪个用户
	CommentStatus string    //"有评论" or "无评论"
	Comments      []Comment // 一对多：Post → Comments
}

// 钩子函数：文章创建后，更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	//每次创建文章后，更新对应用户的 PostCount 字段
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}
