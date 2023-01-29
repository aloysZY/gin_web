// Package errcode 定义错误码
package errcode

// code 是根据需求来定义的
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotLogin                  = NewError(10000002, "未登录")
	RegistrationFailed        = NewError(10000003, "创建账号失败,用户名已存在")
	UnauthorizedAuthNotExist  = NewError(10000004, "鉴权失败，找不到对应的用户名和密码")
	UnauthorizedTokenError    = NewError(10000005, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000006, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000007, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000008, "请求过多")
	GatewayTimeout            = NewError(10000009, "请求超时")
)
