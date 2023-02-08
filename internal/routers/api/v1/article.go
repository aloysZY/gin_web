// Package v1 文章接口，根据创建文章的标签来查询对于的文章
package v1

import (
	"strconv"

	"gin_web/internal/routers/api/params"
	"gin_web/internal/service"
	"gin_web/pkg/app"
	"gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// NewArticle 路由初始化使用，空结构体，使用指针和值都没关系
func NewArticle() Article { return Article{} }

// List 查询文章
// @Summary 查询文章
// @Description 查询文章 支持文章名称模糊查找
// @Tags 文章
// @Produce  json
// @Param title query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.SwaggerArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
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
	// 这里应该还有一步，根据文章 ID，去查找文章标签表中的文章对于的标签，用标签 ID 查找标签表，返回标签名称,在svc.ListArticleS里面做了

	// 3.设置每页返回的内容数量，返回数据
	response.ToResponseList(articleList, totalRows)
}

// Create 创建文章
// @Summary 创建文章
// @Description 创建文章接口
// @Tags 文章
// @Produce  json
// @Param object body params.CreateArticleRequest true "创建文章"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
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
	response.ToResponse(gin.H{})
}

// Get 根据文章 ID 获取单个文章
// @Summary 根据文章 ID 获取单个文章
// @Description 根据文章 ID 获取单个文章
// @Tags 文章
// @Produce  json
// @Param id path string false "文章ID" number
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Security ApiKeyAuth
// @Success 200 {object} third_party.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	response := app.NewResponse(c)
	// 1.解析参数
	articleIdStr := c.Param("id")
	articleId, err := strconv.ParseUint(articleIdStr, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 业务处理
	svc := service.New(c.Request.Context())
	article, err := svc.GetArticleByArticleId(articleId)
	if err != nil {
		if err == errcode.ErrorNotArticle {
			response.ToErrorResponse(errcode.ErrorNotArticle)
			return
		}
		response.ToErrorResponse(errcode.ErrorNotArticle.WithDetails(err.Error()))
		return
	}
	// 返回数据
	response.ToResponse(article)
}

// 更新文章
func (a Article) Update(c *gin.Context) {

}

// 删除文章
func (a Article) Delete(c *gin.Context) {

}
