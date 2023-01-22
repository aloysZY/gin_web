package api

import (
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/internal/service"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

func NewAuth() Auth { return Auth{} }

// SignUp 注册
// @Summary 注册
// @Description 注册接口
// @Tags 用户
// @Produce  json
// @Param object body params.AuthRequest true "注册用户"
// @Success 200 {object} app.Swagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /signup [post]
func (a *Auth) SignUp(c *gin.Context) {
	response := app.NewResponse(c)
	param := params.AuthRequest{}
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	// 后期添加创建用户的时候验证用户名是否存在，其实数据库uesr_name是索引不能重复
	if err := svc.CreateAuth(&param); err != nil {
		response.ToErrorResponse(errcode.RegistrationFailed)
		return
	}
	response.ToResponse(gin.H{})
}

// Auth 登录
// @Summary 登录
// @Description 登录接口
// @Tags 用户
// @Produce  json
// @Param object body params.AuthRequest true "用户登录"
// @Success 200 {object} app.SwaggerAuth "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /auth [post]
func (a *Auth) Auth(c *gin.Context) {
	response := app.NewResponse(c)
	param := params.AuthRequest{}
	if valid, errs := app.BindAndValid(c, &param); !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	// 先检查账号密码是否存在
	if err := svc.CheckAuth(&param); err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	// 生成 token，传入userID
	token, err := app.GenerateToken(param.UserId)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token": token})
	return
}
