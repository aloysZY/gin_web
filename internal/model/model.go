package model

import (
	"fmt"

	"github.com/aloysZy/gin_web/global"
	"github.com/aloysZy/gin_web/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model 公共的字段
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"` // 主键
	CreatedBy  string `json:"created_by"`            // 创建人
	ModifiedBy string `json:"modified_by"`           // 修改人
	CreatedOn  uint32 `json:"created_on"`            // 创建时间 ，自动获取提交时间
	ModifiedOn uint32 `json:"modified_on"`           // 修改时间，自动获取提交时间
	DeletedOn  uint32 `json:"deleted_on"`            // 删除时间，自动获取提交时间
	IsDel      uint8  `json:"is_del"`                // 是否删除 0为删除，1 已删除
}

// NewMysqlDBEngine 初始化 MySQL
func NewMysqlDBEngine(mysqlDatabaseSetting *setting.MysqlSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysqlDatabaseSetting.DBType, fmt.Sprintf(
		s,
		mysqlDatabaseSetting.UserName,
		mysqlDatabaseSetting.Password,
		mysqlDatabaseSetting.Host,
		mysqlDatabaseSetting.Port,
		mysqlDatabaseSetting.DBName,
		mysqlDatabaseSetting.Charset,
		mysqlDatabaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	// 根据日志级别来设置日志详情
	if global.AppSetting.Level == "debug" {
		// db.LogMode(false) 关闭 Logger, 不再展示任何日志，即使是错误日志
		db.LogMode(true) // 开启 Logger, 以展示详细的日志
	}
	// 在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。
	db.SingularTable(true) // 让grom转义struct名字的时候不用加上"s"
	// maxIdleCount 最大空闲连接数，默认不配置，是2个最大空闲连接
	// maxOpen 最大连接数，默认不配置，是不限制最大连接数
	// maxLifetime 连接最大存活时间
	// maxIdleTime 空闲连接最大存活时间
	db.DB().SetMaxIdleConns(mysqlDatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysqlDatabaseSetting.MaxOpenConns)
	return db, nil
}
