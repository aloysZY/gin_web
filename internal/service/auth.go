package service

import (
	"errors"
	"fmt"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/internal/routers/api/params"
	"github.com/aloysZy/gin_web/pkg/setting"
	"github.com/gin-gonic/gin"
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
	// 如果大于 0，就是正常的，赋值后返回
	if auth.UserId > 0 {
		param.UserId = auth.UserId // 有多重方法可以实现这个值传出去，这里比较方便
		return nil
	}
	return errors.New("auth info does not exist.")
}

func GetUserID(c *gin.Context) (uint64, error) {
	_userID, ok := c.Get(global.UserId)
	if !ok {
		zap.L().Error("GetUserID not found user")
		return 0, fmt.Errorf("GetUserID not found user")
	}
	userID, ok := _userID.(uint64)
	if !ok {
		zap.L().Error("GetUserID not found user")
		return 0, fmt.Errorf("GetUserID not found user")
	}
	return userID, nil
}
