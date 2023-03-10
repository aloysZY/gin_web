// Package v1 对于的相关请求 请求参数处理
package v1

import (
	"strconv"

	"gin_web/internal/routers/api/params"
	"gin_web/internal/service"
	"gin_web/pkg/app"
	"gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Tag 类型别名
type Tag struct{}

// NewTag Tag 这里的作用是在 router 的时候初始化，路由调用方法
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

// http://127.0.0.1:8080/api/v1/tags?page=2&page_size=1,展示第2 夜的，每页是一条数据，所以展示数据库第二条数据

// List 查询标签
// @Summary 查询标签
// @Description 查询标签
// @Tags 标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.SwaggerTage "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	response := app.NewResponse(c)
	// 1.解析参数
	// params.ListTagRequest{} 有一个问题，初始化后，没传入state参数，解析后 state 是 1，有问题
	param := params.ListTagRequest{} // state 怎么解析后就是 1 了？、
	// State uint8  `form:"state,default=1"  是这个default的问题，只有 get 请求才生效
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 2.业务逻辑处理
	svc := service.New(c.Request.Context())
	// 解析 URL 传入的页码和每页展示数量
	pager := app.Pager{
		Page:     app.GetPage(c),     // 第几页
		PageSize: app.GetPageSize(c), // 每页显示多少数据
	}
	// 根据page 和 param 等参数去查询数据相关数据
	tagList, err := svc.ListTag(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	// 查询符合条件的标签总数，返回前端
	totalRows, err := svc.CountTag(&params.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	// panic("failed to get tag list test email")
	response.ToResponseList(tagList, totalRows)
	return
}

// Create 添加标签
// @Summary 添加标签
// @Description 添加标签接口
// @Tags 标签
// @Produce  json
// @Param object body params.CreateTagRequest true "创建标签"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	response := app.NewResponse(c)
	// 1. 解析参数
	// param := params.CreateTagRequest{} // 也有一个问题，没传入 state，解析后 state 是 0，这是正常的
	param := params.CreateTagRequest{State: 1}
	userID, err := service.GetUserID(c)
	if err != nil {
		response.ToErrorResponse(errcode.NotLogin)
		return
	}
	param.CreatedBy = userID
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 2.处理业务逻辑 ,c.Request.Context() 请求行下文传入
	// 创建 tagID
	svc := service.New(c.Request.Context())
	if err := svc.CreateTag(&param); err != nil {
		if err == errcode.ErrorTagExists {
			response.ToErrorResponse(errcode.ErrorTagExists)
			return
		}
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	// 返回请求
	response.ToResponse(gin.H{})
	return
}

// Update 更新标签
// @Summary 更新标签
// @Description 更新标签接口
// @Tags 标签
// @Produce  json
// @Param id path uint64 true "标签ID"
// @Param object body params.UpdateTagRequest true "更新标签"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := c.Param("id")
	parseUInt64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// 如果现在是启用状态1，我这样初始化后，更新的是标签的名称数据库修改了，就变成 0 了，？存在问题
	// 初始化赋值，默认为修改名称
	param := params.UpdateTagRequest{TagId: parseUInt64, State: 1}
	userID, err := service.GetUserID(c)
	if err != nil {
		response.ToErrorResponse(errcode.NotLogin)
		return
	}
	param.ModifiedBy = userID
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.UpdateTag(&param)
	if err != nil {
		if err == errcode.ErrorNoDataModified {
			response.ToErrorResponse(errcode.ErrorNoDataModified)
			return
		}
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// Delete 删除标签
// @Summary 删除标签
// @Description 删除标签
// @Tags 标签
// @Produce  json
// @Param id path uint64 true "标签ID"
// @Security ApiKeyAuth
// @Success 200 {object} third_party.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	idStr := c.Param("id")
	parseUInt64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	param := params.DeleteTagRequest{TagId: parseUInt64}
	userID, err := service.GetUserID(c)
	if err != nil {
		response.ToErrorResponse(errcode.NotLogin)
		return
	}
	param.ModifiedBy = userID
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err = svc.DeleteTag(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
