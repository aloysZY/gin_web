package dao

import "github.com/aloysZy/gin_web/internal/model"

func (d *Dao) CreateAuth(userId uint64, appKey, appSecret string) error {
	auth := model.Auth{UserID: userId, AppKey: appKey, AppSecret: appSecret}
	return auth.Create(d.engine)
}

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
