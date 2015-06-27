package db

import (
	"fmt"
	"github.com/ascoders/as/conf"
	"gopkg.in/mgo.v2"
)

var (
	Db *mgo.Database // 数据库连接池
)

func Connect() {
	//获取数据库连接
	session, err := mgo.Dial(conf.ConfInstance.MongodbAddress)

	if err != nil {
		fmt.Println(conf.ConfInstance.MongodbAddress)
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	Db = session.DB("woku")
}
