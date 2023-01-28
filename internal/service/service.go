package service

import (
	"context"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

// Service 封装了上下文和 dao
type Service struct {
	ctx context.Context
	dao *dao.Dao
}

// New 初始化svc 上下文和 dao
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// svc.dao = dao.New(global.MysqlDBEngine)
	// 上下文链路传给 dao，这样 dao 执行的时候记录一条链路
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.MysqlDBEngine))
	return svc
}
