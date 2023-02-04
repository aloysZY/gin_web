package params

// https://segmentfault.com/a/1190000023725115

// AuthRequest 登录解析信息
type AuthRequest struct {
	UserId uint64 `json:"user_id,string" from:"user" swaggerignore:"true"` // 注册用户 ID，在登录的时候去数据库查找，赋值到上下文，之后的所有关于用户的记录都是记录的这个 ID
	// UserName  string `json:"user_name" from:"user_name" binding:"required,min=2,max=20"`
	AppKey    string `json:"app_key" from:"app_key" binding:"required,email"`
	AppSecret string `json:"app_secret" from:"app_secret" binding:"required,min=5,max=20"`
	// ReAppSecret string `json:"re_app_secret" from:"re_app_secret" binding:"required,eqfield=AppSecret"`
}
