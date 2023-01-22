package service

import (
	"errors"
	"fmt"

	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/pkg/setting"
	"go.uber.org/zap"
)

// CreateAuth 创建用户
func (svc *Service) CreateAuth(param *params.AuthRequest) error {
	// 雪花算法生成 ID
	id, err := setting.GetID()
	if err != nil {
		if err == fmt.Errorf("newSonyFlake not initialized") {
			zap.L().Error("SonyFlake not initialized error: ", zap.Error(err))
			return err
		}
		zap.L().Error("CreateAuth app.GetID error: ", zap.Error(err))
		return err
	}
	param.UserId = id
	if err = svc.dao.CreateAuth(param.UserId, param.AppKey, param.AppSecret); err != nil {
		zap.L().Error("svc.dao.CreateAuth error: ", zap.Error(err))
		return err
	}
	return nil
}

// CheckAuth 检查用户
func (svc *Service) CheckAuth(param *params.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		zap.L().Error("svc.dao.GetAuth error: ", zap.Error(err))
		return err
	}
	// 这里判断是怕初始化的值查询失败，传回来了
	if auth.UserID > 0 {
		zap.L().Error("svc.dao.GetAuth userIDs error: ", zap.Error(err))
		return nil
	}
	return errors.New("auth info does not exist.")
}
