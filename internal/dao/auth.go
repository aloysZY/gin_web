package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateAuth(userId uint64, appKey, appSecret string) error {
	auth := model.Auth{UserId: userId, AppKey: appKey, AppSecret: appSecret}
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
