package as

import (
	"github.com/ascoders/as/db"
	"gopkg.in/mgo.v2"
)

var (
	Db *mgo.Database // 数据库连接池
)

func init() {
	Db = db.Db
}
