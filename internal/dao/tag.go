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

// CountTag 根据传入参数统计返回值数量
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

// GetTagList 根据传入参数，返回查询结果
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

// UpdateTag 更新标签
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	// 这里面传给 model 执行应该是数据处理好的
	// tag := model.Tag{
	// 	State: state,
	// 	Name:  name,
	// 	Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	// }

	tag := model.Tag{Model: &model.Model{ID: id}}
	value := map[string]any{
		// 这样写字段名称要和数据库对应了，应为没有解析了
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		value["name"] = name
	}
	return tag.Update(d.engine, value)
}
