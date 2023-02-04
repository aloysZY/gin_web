// Package model 数据库相关模型，整合数据，调用 dao
package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 公共的字段
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"-"` // 主键,"-"反序列化的时候不返回这个ID，新建雪花算法创建 ID
	CreatedOn  uint32 `json:"created_on"`           // 创建时间 ，自动获取提交时间
	ModifiedOn uint32 `json:"modified_on"`          // 修改时间，自动获取提交时间
	DeletedOn  uint32 `json:"-"`                    // `json:"deleted_on"`           // 删除时间，自动获取提交时间
	CreatedBy  uint64 `json:"-"`                    // `json:"created_by"`           // 创建人
	ModifiedBy uint64 `json:"-"`                    // `json:"modified_by"`          // 修改人
	IsDel      uint8  `json:"-"`                    // `json:"is_del,omitempty"`     // 是否删除 0 正常,1为删除，默认初始化 0
}

// ArticleTag 公共结构体，查询文章后，查询文章对应的标签(转化为名称)，整合后返回
type ArticleTag struct {
	TagName  []string `json:"tag_name,omitempty"` // 这里可以设置，如果为空就不返回这列，跟前端协商
	UserName []string `json:"user_name"`
	*Article `json:"article"`
}
