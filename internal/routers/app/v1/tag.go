// Package v1 对于的相关请求
package v1

import (
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/param"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag { return Tag{} }

/*注解	描述
@Summary	摘要
@Tags 标签  分组，一组的使用一个
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释*/

// 测试的时候 post 不能用三个 body 参数，不然 swagger 生成命令出错
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)

// List 查询标签
// @Summary 获取多个标签
// @Description 获取多个标签
// @Tags 标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TageSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	response := app.NewResponse(c)
	params := param.TagListRequest{}
	// 解析参数
	b, v := app.BindAndValid(c, &params)
	if !b {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(v.Errors()...))
		return
	}
	response.ToResponse(gin.H{"name": params.Name, "state": params.State})
}

// Create 新增标签
// @Summary 新增标签
// @Description 添加标签接口
// @Tags 标签
// @Produce  json
// @Param object body param.CreateTagRequest true "创建标签"
// @Success 200 {object} model.TageSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	response := app.NewResponse(c)
	params := param.CreateTagRequest{}
	// 解析参数
	b, v := app.BindAndValid(c, &params)
	if !b {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(v.Errors()...))
		return
	}
	response.ToResponse(gin.H{"name": params.Name, "createdBy": params.CreatedBy, "state": params.State})
}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}

// 这层应该做的是返回数据到前端的，基本不做业务处理
