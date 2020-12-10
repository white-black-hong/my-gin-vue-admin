package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"my-gin-vue-admin/global"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("Mysql")
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if global.GVA_CONFIG.Mysql.LogZap {
		return &gorm.Config{
			Logger:                                   Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
		if mod {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Info),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		} else {
			return &gorm.Config{
				Logger:                                   logger.Default.LogMode(logger.Silent),
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}
	}

}
