// Package dao 组合数据模型，调用数据库方法
package dao

import (
	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/pkg/app"
)

// 聚合数据库操作需要的信息

// CreateTag 创建标签需要的数据整合
// 传入什么参数，是根据数据库列，就是说创建数据库后创建的模型
// 方法用指针类型接收者，因为 dao 结构体太大了
func (d *Dao) CreateTag(name, createdBy string, state uint8) error {
	// 	根据传入的操作初始化结构体，整合操作数据库数据
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	// 调用类型方法
	return tag.Create(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		State: state,
		Name:  name,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	listTag, err := tag.List(d.engine, pageOffset, pageSize)
	if err != nil {
		return nil, err
	}
	return listTag, nil
}
