package params

// binding 是 gin 框架使用 validator 做参数校验
// form 是前端传入参数匹配结构体解析

type CountArticleRequest struct {
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1" example:"1"`
	ArticleId uint64 `json:"article_id" form:"article_id" binding:"required"` // 标签 ID
}

// ListArticleRequest 查询文章
type ListArticleRequest struct {
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1" example:"1"`
	ArticleId uint64 `json:"article_id" form:"article_id" binding:"required"` // 标签 ID
}

// CreateArticleRequest
// 创建文章
type CreateArticleRequest struct {
	State         uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1" example:"1"`
	ArticleId     uint64 `json:"article_id" form:"article_id" swaggerignore:"true"`
	TagId         uint64 `json:"tag_id,string" from:"tag_id"`                                    // 标签 ID,创建文章可以不设置标签
	CreatedBy     uint64 `json:"created_by" form:"created_by" swaggerignore:"true"`              // 创建人；以后从 token 中获取；min 和 max 限制的是长度 2-100s
	Title         string `json:"title" form:"title" binding:"required,min=2,max=100"`            // 文章标题
	Desc          string `json:"desc" form:"desc" binding:"required,min=2,max=255"`              // 文章描述
	Content       string `json:"content" form:"content" binding:"required,min=2,max=4294967295"` // 文章内容
	CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"required,url"`  // 文章封面
}