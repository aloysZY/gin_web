// Package service 入参校验
// 聚合多个数据库查询数据，返回给上层
package service

import (
	"context"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/internal/dao"
	"github.com/aloysZy/gin_web/internal/model"
	"github.com/aloysZy/gin_web/pkg/app"
	"github.com/aloysZy/gin_web/pkg/params"
	"go.uber.org/zap"
)

// Service 封装了上下文和 dao
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

// New 初始化svc 上下文和 dao
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.MysqlDBEngine)
	return svc
}

// CreateTag 创建标签的具体逻辑函数
func (svc *Service) CreateTag(param *params.CreateTagRequest) error {
	// 业务逻辑操作，处理业务需要的数据和数据库需要的数据，调用 dao操作数据库
	err := svc.dao.CreateTag(param.Name, param.CreatedBy, param.State)
	if err != nil {
		zap.L().Error("svc.dao.CreateTag error: ", zap.Error(err))
	}
	return err
}

func (svc *Service) CountTag(param *params.CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *params.ListTagRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

// func (svc *Service) UpdateTag(param *param.UpdateTagRequest) error {
// 	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
// }
//
// func (svc *Service) DeleteTag(param *param.DeleteTagRequest) error {
// 	return svc.dao.DeleteTag(param.ID)
// }
