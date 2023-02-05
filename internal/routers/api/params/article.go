package params

// binding 是 gin 框架使用 validator 做参数校验
// form 是前端传入参数匹配结构体解析

// CountArticleRequest 查询文章数量
type CountArticleRequest struct {
	State uint8 `json:"state" form:"state"`
	// TagId uint64 `json:"tag_id" form:"tag_id" binding:"required"` // 文章标签，根据标签查询数量
	Title string `json:"title" form:"title"` // 这里想要额外实现一个根据文章名称查询（创建文章的时候可以没有设置标签）
}

// ListArticleRequest 查询文章,全部文章和根据标题模糊查找接口
type ListArticleRequest struct {
	State uint8 `json:"state" form:"state,default=1" binding:"required_with_all=TagId Title,oneof=0 1" example:"1"`
	// TagId uint64 `json:"tag_id,staring" form:"tag_id"` // 文章标签，根据标签查询文章
	Title string `json:"title" form:"title"` // 这里想要额外实现一个根据文章名称查询（创建文章的时候可以没有设置标签）
}

// CreateArticleRequest 创建文章
type CreateArticleRequest struct {
	State         uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1" example:"1"`
	ArticleId     uint64 `json:"article_id" form:"article_id" swaggerignore:"true"`
	TagId         uint64 `json:"tag_id,string,omitempty" from:"tag_id" binding:"number"`         // 标签 ID,创建文章可以不设置标签 omitempty 为空不显示
	CreatedBy     uint64 `json:"created_by" form:"created_by" swaggerignore:"true"`              // 创建人；以后从 token 中获取；min 和 max 限制的是长度 2-100s
	Title         string `json:"title" form:"title" binding:"required,min=2,max=100"`            // 文章标题
	Desc          string `json:"desc" form:"desc" binding:"required,min=2,max=255"`              // 文章描述
	Content       string `json:"content" form:"content" binding:"required,min=2,max=4294967295"` // 文章内容
	CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"required,url"`  // 文章封面
}

// GetArticleRequest 根据 ID 获取单个文章
type GetArticleRequest struct {
	ArticleId uint64 `json:"article_id" form:"article_id"`
}
