package setting

import "time"

// config 的配置结构化到结构体

// ServerSettingS 服务配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSettingS 应用配置
type AppSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
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
	ParseTime    string
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Host string
	Port string
}

// DatabaseSettingS 数据库配置
type DatabaseSettingS struct {
	Mysql *MysqlSettingS
	Redis *RedisSettingS
}

// ReadSection 序列化到结构体
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
