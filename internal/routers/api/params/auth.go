package params

// https://segmentfault.com/a/1190000023725115

// AuthRequest 登录解析信息
type AuthRequest struct {
	UserId    uint64 `json:"user_id,string" from:"user" swaggerignore:"true"`
	AppKey    string `json:"app_key" from:"app_key" binding:"required,min=5,max=20"`
	AppSecret string `json:"app_secret" from:"app_secret" binding:"required,min=5,max=20"`
	// ReAppSecret string `json:"re_app_secret" from:"re_app_secret" binding:"required,eqfield=AppSecret"`
}
