// Package v1 对于的相关请求
package v1

import (
	"github.com/aloysZy/gin_web/internal/service"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Tag struct{}

func NewTag() Tag { return Tag{} }

func (t Tag) Get(c *gin.Context) {}

func (t Tag) List(c *gin.Context) {}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TageSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	// app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	// return
	response := app.NewResponse(c)
	param := service.CreateTagRequest{}
	// 解析参数
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		zap.L().Error("app.BindAndValid errs:", zap.Error(errs))
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}
