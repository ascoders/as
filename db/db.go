package db

import (
	"github.com/ascoders/as/conf"
	//"github.com/ascoders/as/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db gorm.DB // 数据库连接
)

func init() {
	Db, _ = gorm.Open("mysql", "root:aaaa@/woku?charset=utf8&parseTime=True&loc=Local")

	// 设置最大空闲连接数、最大连接数
	Db.DB().SetMaxIdleConns(conf.ConfInstance.DBMaxIdleConns)
	Db.DB().SetMaxOpenConns(conf.ConfInstance.DBMaxOpenConns)

	// 禁止表名多元化
	Db.SingularTable(true)
}

func Connect() {
	/*
		//获取数据库连接
		session, err := mgo.Dial(conf.ConfInstance.MongodbAddress)

		if err != nil {
			panic(err)
		}

		session.SetMode(mgo.Monotonic, true)
		DbInstance.DataBase = session.DB(conf.ConfInstance.MongodbDbName)

		// 实例化各个表
		for _, v := range models.BaseLists {
			v.Collection = DbInstance.DataBase.C(v.Table)

			for k, _ := range v.Indexs {
				if err := v.Collection.EnsureIndex(v.Indexs[k]); err != nil {
					panic(err)
				}
			}
		}

		// 释放内存
		models.BaseLists = nil
	*/
}
