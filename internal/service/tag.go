// Package service 入参校验
package service

// 这层基本做的是业务逻辑封装，以后可以抽离一层接口校验层

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
