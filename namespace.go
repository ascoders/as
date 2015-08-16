package as

import (
	"github.com/ascoders/as/conf"
	"github.com/ascoders/as/controllers"
	"github.com/ascoders/as/data"
	"github.com/ascoders/as/email"
	"github.com/ascoders/as/lib"
	"github.com/ascoders/as/models"
	"github.com/ascoders/as/redis"
	"github.com/ascoders/as/router"
)

var (
	Lib    *lib.Lib       // 基础库
	Conf   *conf.Conf     // 基础配置
	Router *router.Router // 路由
	Redis  *redis.Redis   // 缓存
	Email  *email.Email   // 邮件
)

type Controller struct {
	controllers.Base
}

type Model struct {
	models.Base
}

type Data struct {
	data.Data
}

func init() {
	Lib = lib.LibInstance
	Conf = conf.ConfInstance
	Router = router.RouterInstance
	Redis = redis.RedisInstance
	Email = email.EmailInstance
}
