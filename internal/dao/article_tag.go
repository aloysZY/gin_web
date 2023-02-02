package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateArticleTag(articleId, tagId uint64) error {
	articleTag := &model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return articleTag.Create(d.Engine)
}
