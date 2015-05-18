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
	"newWoku/lib/captcha"
	"newWoku/lib/response"
	"newWoku/lib/router"
	"os"
)

// 设置路由
func Route() martini.Router {
	r := martini.NewRouter()

	// 生成注解路由
	if conf.DEBUG {
		router.Auto(
			router.Options{
				AutoCsrf: conf.AUTO_SCRF,
			},
			&user.Controller{},
			&article.Controller{},
			&app.Controller{},
		)
	}

	// 添加注解路由
	AutoRoute(r)

	// 验证码
	// 获取验证码
	r.Get("/api/captcha/:id", captcha.Image)
	// 创建验证码
	r.Post("/api/captcha", func() (int, []byte) {
		return response.Success(map[string]interface{}{
			"captchaCode": captcha.Code(),
		})
	})
	// 验证验证码
	r.Get("/api/captcha", captcha.Check)

	// 匹配未定义的api
	r.Any("/api/**", func() (int, []byte) {
		return response.Error("Api Not Found")
	})

	// 全局模版文件
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
