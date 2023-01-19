// Package model 数据库相关模型
package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 公共的字段
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"-"` // 主键,"-"反序列化的时候不返回 ID
	CreatedOn  uint32 `json:"created_on"`           // 创建时间 ，自动获取提交时间
	ModifiedOn uint32 `json:"modified_on"`          // 修改时间，自动获取提交时间
	DeletedOn  uint32 `json:"deleted_on"`           // 删除时间，自动获取提交时间
	CreatedBy  string `json:"created_by"`           // 创建人
	ModifiedBy string `json:"modified_by"`          // 修改人
	IsDel      uint8  `json:"is_del"`               // 是否删除 0为删除，1 已删除
}
