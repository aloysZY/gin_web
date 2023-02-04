package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateArticleTag(articleId, tagId, createdBy uint64) error {
	articleTag := &model.ArticleIdTagId{
		ArticleId: articleId,
		TagId:     tagId,
		Model:     &model.Model{CreatedBy: createdBy},
	}
	return articleTag.Create(d.Engine)
}

func (d *Dao) ListTagNameByArticleId(articleId uint64) ([]string, error) {
	articleTag := &model.ArticleIdTagId{ArticleId: articleId}
	ListTagName, err := articleTag.ListTagNameByArticleId(d.Engine)
	if err != nil {
		return nil, err
	}
	return ListTagName, nil
}
