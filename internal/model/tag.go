package model

import "github.com/aloysZy/gin_web/pkg/app"

// Tag TAG生成返回的数据
type Tag struct {
	*Model        // 公共字段
	Name   string `json:"name"`   // 名称
	Status string `json:"status"` // 状态
}

// TageSwagger swagger接口返回数据用的结构体,将需要Tag以外的字段的时候，使用这个，这样就将字段数据包含进去了，这个就是额外添加了页码信息
type TageSwagger struct {
	List []*Tag
	Page *app.Pager
}
