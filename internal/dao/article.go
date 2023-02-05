package dao

import (
	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/pkg/app"
)

// Article 文章结构体，tag 使用形参传参，但是文章内容太大了，使用结构体指针
// 这个传入的时候使用 model.Article了。这个也没有用了
// type Article struct {
// 	State     uint8  `json:"state"`
// 	ArticleId uint64 `json:"article_id"`
// 	// TagId         uint64 `json:"tag_id"` 建立关联表
// 	CreatedBy     uint64 `json:"created_by"` // 创建人；以后从 token 中获取；min 和 max 限制的是长度 2-100s
// 	Title         string `json:"title"`
// 	Desc          string `json:"desc"`
// 	Content       string `json:"content"`
// 	CoverImageUrl string `json:"cover_image_url"`
// }

// // CreateArticle 创建文章
// func (d *Dao) CreateArticle(param *Article) error {
// 	article := &model.Article{
// 		State:     param.State,
// 		ArticleId: param.ArticleId,
// 		// TagId:         param.TagId,
// 		Title:         param.Title,
// 		Desc:          param.Desc,
// 		CoverImageUrl: param.CoverImageUrl,
// 		Content:       param.Content,
// 		Model:         &model.Model{CreatedBy: param.CreatedBy},
// 	}
// 	return article.Create(d.Engine)
// }

// CreateArticle 创建文章
func (d *Dao) CreateArticle(param *model.Article) error {
	return param.Create(d.Engine)
}

// CountArticle 查询文章数量
func (d *Dao) CountArticle(state uint8) (int, error) {
	article := &model.Article{
		State: state,
	}
	return article.CountArticle(d.Engine)
}

// CountArticleByTitle 根据标题查询文章数量
func (d *Dao) CountArticleByTitle(title string, state uint8) (int, error) {
	article := &model.Article{Title: title, State: state}
	return article.CountArticleByTitle(d.Engine)
}

// ListArticle 查询文章列表
func (d *Dao) ListArticle(state uint8, page, pageSize int) ([]*model.Article, error) {
	article := &model.Article{State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.ListArticle(d.Engine, pageOffset, pageSize)
}

// ListArticleByTitle 根据标题查询文章列表
func (d *Dao) ListArticleByTitle(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := &model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.ListArticleByTitle(d.Engine, pageOffset, pageSize)
}

// GetArticleByArticleId 根据文章 ID 获取单个文章
func (d *Dao) GetArticleByArticleId(id uint64) (*model.Article, error) {
	article := &model.Article{ArticleId: id}
	return article.GetArticleByArticleId(d.Engine)
}
