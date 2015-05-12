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
		&user.User{},
		&article.Article{},
		&app.App{},
	)

	// 加入注解路由
	AutoRoute(r)

	// 匹配未定义的api
	r.Get("/api/**", func() []byte {
		return response.Error("这个api走丢了")
	})

	// 最后匹配的是全局内容
	globalFile, err := os.Open(conf.GLOBAL_PATH)
	if err != nil {
		panic(err)
	}
	globalFileText, err := ioutil.ReadAll(globalFile)
	globalFile.Close()

	r.Get("/**", func() []byte {
		return globalFileText
	})

	return r
}
