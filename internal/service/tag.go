// Package service 入参校验
// 聚合多个数据库查询数据，返回给上层
package service

import (
	"fmt"

	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/internal/routers/api/v1/params"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/errcode"
	"github.com/aloysZy/gin_web/pkg/setting"

	"go.uber.org/zap"
)

// CreateTag 创建标签的具体逻辑函数
func (svc *Service) CreateTag(param *params.CreateTagRequest) error {
	// 创建tagId
	id, err := setting.GetID()
	if err != nil {
		if err == fmt.Errorf("newSonyFlake not initialized") {
			zap.L().Error("SonyFlake not initialized error: ", zap.Error(err))
			return err
		}
		zap.L().Error("CreateTag app.GetID error: ", zap.Error(err))
		return err
	}
	param.TagId = id
	// 创建的时候标签去重
	// 查询数据库，看看标签存在不存在,存在返回标签存在错误
	count, err := svc.dao.GetTag(param.Name, param.State)
	if count != 0 {
		err = errcode.ErrorTagExists
		// err = errors.New("TagName exists")
		zap.L().Error("svc.dao.GetTag failed:", zap.Error(err))
		return err
	}

	// 业务逻辑操作，处理业务需要的数据和数据库需要的数据，调用 dao操作数据库
	err = svc.dao.CreateTag(param.TagId, param.CreatedBy, param.Name, param.State)
	if err != nil {
		zap.L().Error("svc.dao.CreateTag error: ", zap.Error(err))
		return err
	}
	return nil
}

func (svc *Service) CountTag(param *params.CountTagRequest) (int, error) {
	count, err := svc.dao.CountTag(param.Name, param.State)
	if err != nil {
		zap.L().Error("svc.dao.CountTag error: ", zap.Error(err))
		return 0, err
	}
	return count, nil
}

func (svc *Service) GetTagList(param *params.ListTagRequest, pager *app.Pager) ([]*model.Tag, error) {
	tagList, err := svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
	if err != nil {
		zap.L().Error("svc.dao.GetTagList error: ", zap.Error(err))
		return nil, err
	}
	return tagList, nil
}

func (svc *Service) UpdateTag(param *params.UpdateTagRequest) error {
	// 修改的时候不能修改标签名称，可以重新创建
	err := svc.dao.UpdateTag(param.TagId, param.ModifiedBy, param.State)
	if err != nil {
		zap.L().Error("svc.dao.UpdateTag error: ", zap.Error(err))
		return err
	}
	return nil
}

func (svc *Service) DeleteTag(param *params.DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.TagId, param.ModifiedBy)
}
