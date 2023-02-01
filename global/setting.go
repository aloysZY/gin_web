// Package global 全局变量，配置文件
package global

import (
	"github.com/aloysZy/gin_web/pkg/setting"
)

// 配置文件解析到结构体
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	// JWTSetting      *setting.JWTSettingS
	// EmailSetting    *setting.EmailSettingS
)

// 上下文用到的常量
const (
	// Trans   = "trans"

	UserId  = "user_id"
	TraceId = "X-Trace-ID"
	SpanId  = "X-Span-ID"
)
