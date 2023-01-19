// Package model 操作是数据库
package model

import "github.com/jinzhu/gorm"

// Tag TAG生成返回的数据
type Tag struct {
	State  uint8          `json:"state"` // 状态
	Name   string         `json:"name"`  // 名称
	*Model `json:"model"` // 公共字段
	// json:"model" 本身是结构体嵌套，不写返回数据和 name 是在一层，写了之后，在一个新层里展示
}

// 函数内没有修改结构体数据，所以这里使用的是值传递，避免修改数据

// TableName 实现了接口，gorm 的时候使用这个表明
func (t Tag) TableName() string { return "web_tag" }

// Create 操作数据库进行创建标签
// 这里方法使用的是指针，因为要修改tag信息
func (t Tag) Create(db *gorm.DB) error {
	// 这里使用的是方法，&t 就可以获取到调用的tag数据，数据在 dao 已经初始化了
	// &t传入地址是在结构体字段多的时候降低内存开销
	return db.Create(&t).Error
}

// Count 查找符合条件标签数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// List 查找符合条件标签列表
// 稍后设置为默认按照时间反序排列
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var err error
	var tags []*Tag
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Order("modified_on DESC").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
