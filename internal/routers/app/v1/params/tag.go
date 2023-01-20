// Package params 解析入参参数
package params

// CountTagRequest 查询标签总数
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1" example:"1"` // 1 正常 0为禁用
}

// ListTagRequest 查询标签
type ListTagRequest struct {
	Name  string `form:"name" binding:"max=100"`                          // 名称
	State uint8  `form:"state,default=1" binding:"oneof=0 1" example:"1"` // 状态 1 正常 0为禁用
}

// CreateTagRequest 创建标签入参校验
type CreateTagRequest struct {
	// from是将传入的参数和结构体进行绑定，但是名称中有"_"的时候存在问题，可以设置json来解决
	// https://juejin.cn/post/7005465902804123679
	// example:"1"  swagger tag 设置默认值

	TagId     uint64 `json:"tag_id" form:"tag_id" swaggerignore:"true"`                     // 帖子 ID，算法生成
	Name      string `form:"name" binding:"required,min=2,max=100"`                         // 名称；min 和 max 限制的是长度 2-100
	CreatedBy string `json:"created_by" form:"created_by" binding:"required,min=2,max=100"` // 创建人；以后从 token 中获取；min 和 max 限制的是长度 2-100s
	State     uint8  `form:"state,default=1" binding:"oneof=0 1" example:"1"`               // 创建状态；默认是 1 正常 0为禁用、1为启用
}

// swaggerignore:"true"      swagger忽略这个字段
// https://blog.csdn.net/qq_57467091/article/details/123373790
// https://blog.csdn.net/qq_39397165/article/details/108173108

type UpdateTagRequest struct {
	TagId uint64 `json:"tag_id" form:"tag_id" binding:"required" swaggerignore:"true"` // 标签 ID
	Name  string `form:"name" binding:"max=100"`                                       // 名称;要修改的标签名称
	State uint8  `form:"state,default=1" binding:"oneof=0 1" example:"1"`              // 状态；可以更新状态为不可用，需要传入
	// State      uint8  `form:"state" binding:"required,gte=0,lte=1" example:"1"`                // 状态；可以更新状态为不可用，需要传入
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,min=2,max=100"` // 修改人;以后从token中获取
}

type DeleteTagRequest struct {
	TagId uint64 `json:"tag_id" form:"tag_id" binding:"required"`
}
