package dao

import "github.com/aloysZy/gin_web/internal/model"

// Article 文章结构体，tag 使用形参传参，但是文章内容太大了，使用结构体指针
type Article struct {
	State     uint8  `json:"state"`
	ArticleId uint64 `json:"article_id"`
	// TagId         uint64 `json:"tag_id"` 建立关联表
	CreatedBy     uint64 `json:"created_by"` // 创建人；以后从 token 中获取；min 和 max 限制的是长度 2-100s
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}

func (d *Dao) CreateArticle(param *Article) error {
	article := &model.Article{
		State:     param.State,
		ArticleId: param.ArticleId,
		// TagId:         param.TagId,
		Title:         param.Title,
		Desc:          param.Desc,
		CoverImageUrl: param.CoverImageUrl,
		Content:       param.Content,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.Engine)
}
