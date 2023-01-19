package app

// swagger返回数据用的结构体，返回的还是可能要多个结构体嵌套，在这里组合

import "github.com/aloysZy/gin_web/internal/model"

// SwaggersTage swagger接口返回数据用的结构体,将需要Tag以外的字段的时候，使用这个，这样就将字段数据包含进去了，这个就是额外添加了页码信息
type SwaggersTage struct {
	List []*model.Tag
	Page *Pager
}
