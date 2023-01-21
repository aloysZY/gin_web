// Package model 操作是数据库
package model

import (
	"github.com/jinzhu/gorm"
)

// Tag TAG生成返回的数据
type Tag struct {
	State  uint8          `json:"state"`         // 状态  1 正常 0为禁用
	TagID  uint64         `json:"tag_id,string"` // 设置 tagID  string解决json解析的时候使用这个类型，解决前端传入和传入前端失真
	Name   string         `json:"name"`          // 名称
	*Model `json:"model"` // 公共字段
	// json:"model" 本身是结构体嵌套，不写返回数据和 name 是在一层，写了之后，在一个新层里展示
}

// 函数内没有修改结构体数据，所以这里使用的是值传递，避免修改数据

// TableName 实现了接口，gorm 的时候使用这个表明
func (t Tag) TableName() string { return "web_tag" }

// Create 操作数据库进行创建标签
func (t Tag) Create(db *gorm.DB) error {
	// 这里使用的是方法，&t 就可以获取到调用的tag数据，数据在 dao 已经初始化了
	// &t传入地址是在结构体字段多的时候降低内存开销
	// 这里还有一个问题，要判断标签是否存在，存在，禁用状态不能创建，存在，删除状态可以创建？在上层做判断
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
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) (tags []*Tag, err error) {
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

func (t Tag) Update(db *gorm.DB, values any) error {
	// 当修改的时候，数据库没有对于的数据tag_id，就会返回修改 0 行，但是没有错误，应该要提示数据不存在的，其实无所谓，应为是前端传入的 id，恶意修改数据库本身不存在，也不会修改
	// db = db.Model(&t).Where("tag_id = ? AND is_del = ?", t.TagID, 0).Update(values)
	// if db.RowsAffected == 0 {
	// 	err := errors.New("为进行数据修改")
	// 	return err
	// }
	return db.Model(&t).Where("tag_id = ? AND is_del = ?", t.TagID, 0).Update(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("tag_id = ? AND is_del = ?", t.TagID, 0).Delete(&t).Error
}
