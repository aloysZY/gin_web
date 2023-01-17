// Package service 入参校验
package service

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
	// 这里还是有一个问题，不知道为什么created_by不能找到，添加json tag后能找到了
	// 最后发现，createdby 不添加 json 也可以，应该是"-"导致的问题，以后入参不要写下横线
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"createdBy" binding:"required,min=2,max=100"` // min 和 max 限制的是长度 2-100
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
// 	return svc.dao.CountTag(param.Name, param.State)
// }
//
// func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
// 	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
// }
//
// func (svc *Service) CreateTag(param *CreateTagRequest) error {
// 	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
// }
//
// func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
// 	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
// }
//
// func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
// 	return svc.dao.DeleteTag(param.ID)
// }
