package model

import (
	"github.com/jinzhu/gorm"
)

// ArticleTag 文章和标签关联表
type ArticleTag struct {
	ArticleId uint64 `json:"article_id"`
	TagId     uint64 `json:"tag_id"`
	*Model
}

func (at ArticleTag) TableName() string { return "web_article_tag" }

// Create 建立文章和标签的关联
func (at ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
