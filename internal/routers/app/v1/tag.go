package v1

import (
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
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
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}
