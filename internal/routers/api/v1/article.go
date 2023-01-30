package v1

import (
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/internal/service"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// NewArticle 路由初始化使用，空结构体，使用指针和值都没关系
func NewArticle() Article { return Article{} }

// List 查询文章接口
func (a Article) List(c *gin.Context) {

}

// Create 创建文章
func (a Article) Create(c *gin.Context) {
	response := app.NewResponse(c)
	param := params.CreateArticleRequest{State: 1}
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID, err := service.GetUserID(c)
	if err != nil {
		response.ToErrorResponse(errcode.NotLogin.WithDetails(err.Error()))
		return
	}
	param.CreatedBy = userID

	svc := service.New(c)
	err = svc.CreateArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(param)
}

// 更新文章
func (a Article) Update(c *gin.Context) {

}

// 删除文章
func (a Article) Delete(c *gin.Context) {

}
