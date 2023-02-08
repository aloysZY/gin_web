package model

import (
	"github.com/jinzhu/gorm"
)

// ArticleIdTagId 文章和标签关联表
type ArticleIdTagId struct {
	ArticleId uint64 `json:"article_id,string"`
	TagId     uint64 `json:"tag_id,string"`
	*Model
}

func (at ArticleIdTagId) TableName() string { return "web_articleId_tagId" }

// Create 建立文章和标签的关联
func (at ArticleIdTagId) Create(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

// ListTagNameByArticleId 根据文章 ID，查询文章名称
func (at ArticleIdTagId) ListTagNameByArticleId(db *gorm.DB, state uint8) ([]string, error) {
	// 使用联合查询，在文章和标签关联表中查询到标签 ID，在标签表中根据标签ID，查询标签name,返回到一个列表中
	var articleTag []string
	db.Table(ArticleIdTagId{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON t.tag_id = at.tag_id").
		Where("at.article_id = ? AND state = ? AND t.is_del = ?", at.ArticleId, state, 0).Pluck("name", &articleTag) // Pluck("name")这个name是要查询的字段列名称,聚合的表中有其他重复的列字段，是查询不到的
	return articleTag, nil
}
