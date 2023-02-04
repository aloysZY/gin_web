package params

// SignUpRequest 注册信息
type SignUpRequest struct {
	UserId    uint64 `json:"user_id,string" from:"user" swaggerignore:"true"` // 注册后生成写入到数据库
	UserName  string `json:"user_name" from:"user_name" binding:"required,min=2,max=20"`
	AppKey    string `json:"app_key" from:"app_key" binding:"required,email"`
	AppSecret string `json:"app_secret" from:"app_secret" binding:"required,min=5,max=20"`
	// ReAppSecret string `json:"re_app_secret" from:"re_app_secret" binding:"required,eqfield=AppSecret"`
}
