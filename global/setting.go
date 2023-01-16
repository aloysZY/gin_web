package global

import (
	"github.com/aloysZy/gin_web/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	MysqlDBEngine   *gorm.DB
)
