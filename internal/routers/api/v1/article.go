// Package v1 文章接口，根据创建文章的标签来查询对于的文章
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

// List 查询多个文章接口
func (a Article) List(c *gin.Context) {
	response := app.NewResponse(c)

	// 1.解析请求
	param := params.ListArticleRequest{State: 1}
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 2.业务解析
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	// 根据page 和 param 等参数去查询数据相关数据
	// 解析 URL 传入的页码和每页展示数量
	articleList, totalRows, err := svc.ListArticleS(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetArticlesFail.WithDetails(err.Error()))
		return
	}

	// 这里应该还有一步，根据文章 ID，去查找文章标签表中的文章对于的标签，用标签 ID 查找标签表，返回标签名称
	articleTagList, err := svc.ListTagNameByArticleId(articleList)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetArticlesFail.WithDetails(err.Error()))
		return
	}

	// 返回的应该是另一个结构体
	response.ToResponseList(articleTagList, totalRows)

	// 3.设置每页返回的内容数量，返回数据

}

// Create 创建文章
func (a Article) Create(c *gin.Context) {
	response := app.NewResponse(c)
	// 1.解析参数
	param := params.CreateArticleRequest{State: 1}
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取登录用户 ID
	userID, err := service.GetUserID(c)
	if err != nil {
		response.ToErrorResponse(errcode.NotLogin.WithDetails(err.Error()))
		return
	}
	param.CreatedBy = userID

	svc := service.New(c.Request.Context())
	// 创建文章
	if err = svc.CreateArticle(&param); err != nil {
		if err == errcode.ErrorGetTagFail {
			response.ToErrorResponse(errcode.ErrorGetTagFail)
			return
		}
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
