package db

import (
	"github.com/ascoders/as/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Orm gorm.DB // 数据库连接
)

func InitDatabase() {
	Orm, _ = gorm.Open("mysql", conf.ConfInstance.DbUserName+
		":"+conf.ConfInstance.DbPassword+
		"@/"+conf.ConfInstance.DbName+"?charset=utf8&parseTime=True&loc=Local")

	// 设置最大空闲连接数、最大连接数
	Orm.DB().SetMaxIdleConns(conf.ConfInstance.DBMaxIdleConns)
	Orm.DB().SetMaxOpenConns(conf.ConfInstance.DBMaxOpenConns)

	// 禁止表名多元化
	Orm.SingularTable(true)

	// 不抛出log
	Orm.LogMode(false)
}
