// Package dao 具体数据库操作
package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
