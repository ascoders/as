package db

import (
	"github.com/ascoders/as/conf"
	"gopkg.in/mgo.v2"
)

var (
	Db *mgo.Database // 数据库连接池
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(conf.Conf.MongodbAddress)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	Db = session.DB("woku")
}
