package db

import (
	"github.com/ascoders/as/conf"
	"github.com/ascoders/as/models"
	"gopkg.in/mgo.v2"
)

type Db struct {
	dataBase *mgo.Database // 数据库连接池
}

var (
	DbInstance *Db
)

func init() {
	DbInstance = &Db{}
}

func Connect() {
	//获取数据库连接
	session, err := mgo.Dial(conf.ConfInstance.MongodbAddress)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	DbInstance.dataBase = session.DB("woku")

	// 实例化各个表
	for _, v := range models.BaseLists {
		v.Collection = DbInstance.dataBase.C(v.Table)

		for k, _ := range v.Indexs {
			if err := v.Collection.EnsureIndex(v.Indexs[k]); err != nil {
				panic(err)
			}
		}
	}

	// 释放内存
	models.BaseLists = nil
}
