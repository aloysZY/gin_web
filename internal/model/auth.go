package model

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	UserId    uint64 `json:"user_id"`
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	*Model
}

func (a Auth) TableName() string { return "web_auth" }

func (a Auth) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	// 这里要初始化是要根据 a的参数去查询，结果写入到 auth
	var auth Auth
	// db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	db = db.Where("app_key = ? AND is_del = ?", a.AppKey, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		/*传入接收结果集的变量只能为Struct类型或Slice类型
		当传入变量为Struct类型时，如果检索出来的数据为0条，会抛出ErrRecordNotFound错误
		当传入变量为Slice类型时，任何条件下均不会抛出ErrRecordNotFound错误*/
		return auth, err
	}
	return auth, nil
}
