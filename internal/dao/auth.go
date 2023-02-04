package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateAuth(userId uint64, username, appKey, appSecret string) error {
	auth := model.Auth{UserId: userId, UserName: username, AppKey: appKey, AppSecret: appSecret}
	return auth.Create(d.Engine)
}

// func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
// 	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
// 	return auth.Get(d.Engine)
// }

func (d *Dao) GetAuth(appKey string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey}
	return auth.Get(d.Engine)
}

// GetUserNameByArticleCreatedBy 根据创建人 ID，查询创建人名称（名称在用户表不唯一）
func (d *Dao) GetUserNameByArticleCreatedBy(createdBy uint64) ([]string, error) {
	auth := &model.Auth{UserId: createdBy}
	return auth.GetUserNameByArticleCreatedBy(d.Engine)
}
