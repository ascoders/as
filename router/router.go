package router

import (
	"github.com/ascoders/as/conf"
	"github.com/ascoders/as/lib/captcha"
	"github.com/ascoders/as/lib/response"
	"github.com/go-martini/martini"
	"io/ioutil"
	"os"
)

type Router struct {
	Routes martini.Router // 所有路由规则表
}

var (
	RouterInstance *Router
)

func init() {
	RouterInstance = &Router{}
	RouterInstance.Routes = martini.NewRouter()
}

func RouterListen() {
	// 添加注解路由
	// AutoRoute(RouterInstance.Routes)

	// 是否开启验证码
	if conf.ConfInstance.CaptchaUrl != "" {
		// 获取验证码
		RouterInstance.Routes.Get(conf.ConfInstance.CaptchaUrl+"/:id", captcha.CaptchaInstance.Image)
		// 创建验证码
		RouterInstance.Routes.Post(conf.ConfInstance.CaptchaUrl, func() (int, []byte) {
			return response.ResponseInstance.Success(map[string]interface{}{
				"captchaCode": captcha.CaptchaInstance.Code(),
			})
		})
		// 验证验证码
		RouterInstance.Routes.Get(conf.ConfInstance.CaptchaUrl, captcha.CaptchaInstance.Check)
	}

	// 匹配未定义的api
	RouterInstance.Routes.Any("/api/**", func() (int, []byte) {
		return response.ResponseInstance.Error("Api Not Found")
	})

	// 是否开启全局模版文件
	if conf.ConfInstance.GlobalPath != "" {
		globalFile, err := os.Open(conf.ConfInstance.GlobalPath)
		if err != nil {
			panic(err)
		}
		globalFileText, err := ioutil.ReadAll(globalFile)
		globalFile.Close()
		RouterInstance.Routes.Get("/**", func() (int, []byte) {
			return 200, globalFileText
		})
	}
}
