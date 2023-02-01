package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateArticleTag(articleId, tagId uint64) error {
	ArticleTag := &model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return ArticleTag.Create(d.Engine)
}
