// Package param 解析入参参数
package param

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// CreateTagRequest 创建标签入参校验
type CreateTagRequest struct {
	// 这里还是有一个问题，在 postman 请求的不知道为什么created_by不能找到，添加json tag后能找到了，还是字段解析不到
	// 最后发现，createdby 不添加 json 也可以，应该是"_"导致的问题，以后入参不要写下横线
	// 还是需要加json，swagger传入的时候要解析,不然就是结构体名称
	// example:"1" swagger 的时候默认值

	Name      string `json:"name" form:"name" binding:"required,min=2,max=100"`             // 名称；min 和 max 限制的是长度 2-100
	CreatedBy string `json:"created_by" form:"created_by" binding:"required,min=2,max=100"` // 创建人；min 和 max 限制的是长度 2-100s
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1" example:"1"`  // 创建状态；默认是 1
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modifiedBy" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
