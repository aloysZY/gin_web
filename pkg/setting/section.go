// Package setting 配置文件对于结构体
package setting

import "time"

// config 的配置结构化到结构体

// ServerSettingS 服务配置
type ServerSettingS struct {
	MachineId    uint16
	Name         string
	RunMode      string // 运行级别，有 dev,test 和release
	HttpPort     string // 这里是字符串，因为后面和字符串组合了
	StartTime    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS 应用配置
type AppSettingS struct {
	Log         *LogSettingS
	Page        *PageSettings
	UploadImage *UploadImageSettings
	JWT         *JWTSettingS
	Email       *EmailSettingS
}

type LogSettingS struct {
	// Model string  #和RunMode设置为一个级别吧
	Level       string // 日志级别 记录和显示的日志信息级别
	LogSavePath string
	LogFileName string
	LogFileExt  string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	LocalTime   bool
	Compress    bool
}

type PageSettings struct {
	DefaultPageSize int
	MaxPageSize     int
}

type UploadImageSettings struct {
	UploadImageMaxSize   int
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageAllowExts []string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

// DatabaseSettingS 数据库配置
type DatabaseSettingS struct {
	Mysql *MysqlSettingS
	Redis *RedisSettingS
}

// MysqlSettingS MySQL 配置
type MysqlSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Host string
	Port string
}
