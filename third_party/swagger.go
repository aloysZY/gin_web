package third_party

// swagger返回数据用的结构体，返回的还是可能要多个结构体嵌套，在这里组合

// https://blog.csdn.net/qq_38371367/article/details/123005909  swagger 文档

import (
	"gin_web/internal/model"
	"gin_web/pkg/app"
)

// SwaggerTage swagger接口返回数据用的结构体,将需要Tag以外的字段的时候，使用这个，这样就将字段数据包含进去了，这个就是额外添加了页码信息
type SwaggerTage struct {
	List []*model.Tag
	Page *app.Pager
}

// SwaggerArticle 文章
type SwaggerArticle struct {
	List []*model.ArticleTag
	Page *app.Pager
}

// Swagger 空结构体
type Swagger struct{}

// SwaggerAuth 认证
type SwaggerAuth struct {
	token string
}
