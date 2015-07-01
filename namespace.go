package as

import (
	"github.com/ascoders/as/conf"
	"github.com/ascoders/as/controllers"
	"github.com/ascoders/as/db"
	"github.com/ascoders/as/lib"
	"github.com/ascoders/as/models"
	"github.com/ascoders/as/router"
)

var (
	Db     *db.Db         // 数据库连接池
	Lib    *lib.Lib       // 基础库
	Conf   *conf.Conf     // 基础配置
	Router *router.Router // 路由
)

type Controller struct {
	controllers.Base
}

type Model struct {
	models.Base
}

func init() {
	Db = db.DbInstance
	Lib = lib.LibInstance
	Conf = conf.ConfInstance
	Router = router.RouterInstance
}
