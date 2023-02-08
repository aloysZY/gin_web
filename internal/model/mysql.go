package model

import (
	"fmt"
	"time"

	"gin_web/global"
	"gin_web/pkg/setting"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// https://www.jianshu.com/p/d9304373acc0 gorm操作数据库特殊配置

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
	if global.AppSetting.Log.Level == "debug" {
		// Gorm 建立了对 Logger 的支持，默认模式只会在错误发生的时候打印日志
		// db.LogMode(false) 关闭 Logger, 不再展示任何日志，即使是错误日志
		db.LogMode(true) // 开启 Logger, 以展示详细的日志
	}
	// 在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。
	db.SingularTable(true) // 让gorm转义struct名字的时候不用加上"s"

	// 指定表前缀，修改默认表名,就是配置的表名前缀和操作数据库的结构体名称
	// 这里后面使用func (t Tag) TableName() string {return "blog_tag"}来替换了
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return global.DatabaseSetting.Mysql.TablePrefix + defaultTableName
	// }

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// maxIdleCount 最大空闲连接数，默认不配置，是2个最大空闲连接
	// maxOpen 最大连接数，默认不配置，是不限制最大连接数
	// maxLifetime 连接最大存活时间
	// maxIdleTime 空闲连接最大存活时间
	db.DB().SetMaxIdleConns(mysqlDatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysqlDatabaseSetting.MaxOpenConns)
	// sql追踪回调
	otgorm.AddGormCallbacks(db)
	return db, nil
}

// 创建时间的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		// 添加到数据库是一个时间戳
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新时间的回调函数
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除时间的回调函数
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
