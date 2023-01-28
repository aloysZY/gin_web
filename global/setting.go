// Package global 全局变量，配置文件
package global

import (
	"github.com/aloysZy/gin_web/pkg/email"
	"github.com/aloysZy/gin_web/pkg/setting"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
)

// 解析配置文件需要的结构体
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS

	EmailEngine   *email.Email
	MysqlDBEngine *gorm.DB
	Tracer        opentracing.Tracer
)

// Trans 上下文用到的常量
const (
	Trans   = "trans"
	UserId  = "user_id"
	TraceId = "X-Trace-ID"
	SpanId  = "X-Span-ID"
)
