/*==================================================
	程序入口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package as

import (
	"github.com/ascoders/as/db"
	"github.com/ascoders/as/redis"
	"github.com/ascoders/as/router"
	"strconv"
)

func Run() {
	// 实例化martini对象
	m := newClassic()

	// 初始化并加载路由规则
	router.RouterListen()

	m.Action(router.RouterInstance.Routes.Handle)

	// 初始化数据库链接
	db.Connect()

	// 初始化redis链接
	redis.Connect()

	// 监听端口
	m.RunOnAddr(":" + strconv.Itoa(int(Conf.Port)))
}
