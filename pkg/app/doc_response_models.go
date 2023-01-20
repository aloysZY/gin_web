package app

// swagger返回数据用的结构体，返回的还是可能要多个结构体嵌套，在这里组合

// https://blog.csdn.net/qq_38371367/article/details/123005909  swagger 文档

import "github.com/aloysZy/gin_web/internal/model"

// SwaggerTage swagger接口返回数据用的结构体,将需要Tag以外的字段的时候，使用这个，这样就将字段数据包含进去了，这个就是额外添加了页码信息
type SwaggerTage struct {
	List []*model.Tag
	Page *Pager
}

type Swagger struct{}
