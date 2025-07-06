package models

import "gorm.io/gorm"

type Comment struct {
	ID     uint `gorm:"primaryKey"`
	Body   string
	PostID uint // 外键
	Post   Post // 属于哪个文章
}

// 删除评论后，检查是否为最后一条评论，如果是，则修改文章状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}
