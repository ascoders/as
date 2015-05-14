/*==================================================
	总路由
	内部引用了auto.go（注解路由/自动路由）

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package router

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"newWoku/conf"
	"newWoku/controllers/app"
	"newWoku/controllers/article"
	"newWoku/controllers/user"
	"newWoku/lib/response"
	"newWoku/lib/router"
	"os"
)

// 设置路由
func Route() martini.Router {
	r := martini.NewRouter()

	// 生成注解路由
	router.Auto(
		router.Options{
			AutoCsrf: conf.AUTO_SCRF,
		},
		&user.Controller{},
		&article.Controller{},
		&app.Controller{},
	)

	// 加入注解路由
	AutoRoute(r)

	// 匹配未定义的api
	r.Any("/api/**", func() (int, []byte) {
		return response.Error("Api Not Found")
	})

	// 最后匹配的是全局内容
	globalFile, err := os.Open(conf.GLOBAL_PATH)
	if err != nil {
		panic(err)
	}
	globalFileText, err := ioutil.ReadAll(globalFile)
	globalFile.Close()

	r.Get("/**", func() (int, []byte) {
		return 200, globalFileText
	})

	return r
}
