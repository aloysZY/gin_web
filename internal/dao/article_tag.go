package dao

import "github.com/aloysZy/gin_web/internal/model"

// CreateArticleTag 创建标签和文章关联
func (d *Dao) CreateArticleTag(articleId, tagId, createdBy uint64) error {
	articleTag := &model.ArticleIdTagId{
		ArticleId: articleId,
		TagId:     tagId,
		Model:     &model.Model{CreatedBy: createdBy},
	}
	return articleTag.Create(d.Engine)
}

// ListTagNameByArticleId 根据文章 ID 获取标签名称
func (d *Dao) ListTagNameByArticleId(articleId uint64, state uint8) ([]string, error) {
	articleTag := &model.ArticleIdTagId{ArticleId: articleId}
	ListTagName, err := articleTag.ListTagNameByArticleId(d.Engine, state)
	if err != nil {
		return nil, err
	}
	return ListTagName, nil
}
