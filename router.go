package as

import (
	"fmt"
	//"github.com/ascoders/as/lib/captcha"
	"github.com/ascoders/as/lib/response"
	"github.com/go-martini/martini"
	"io/ioutil"
	"os"
)

var (
	Routers martini.Router
)

func init() {
	Routers = martini.NewRouter()

	// 生成注解路由
	/*
		if Conf.Debug {
			router.Auto(
				router.Options{
					AutoCsrf: Conf.CsrfAuto,
				},
				&user.Controller{},
				&article.Controller{},
			)
		}
	*/

	// 添加注解路由
	// AutoRoute(Routers)

	// 获取验证码
	/*
		Routers.Get("/api/captcha/:id", captcha.Image)
		// 创建验证码
		Routers.Post("/api/captcha", func() (int, []byte) {
			return response.Success(map[string]interface{}{
				"captchaCode": captcha.Code(),
			})
		})
		// 验证验证码
		Routers.Get("/api/captcha", captcha.Check)
	*/

	// 匹配未定义的api
	Routers.Any("/api/**", func() (int, []byte) {
		return response.Error("Api Not Found")
	})

	// 全局模版文件
	fmt.Println("xxxxxxxx", Conf.GlobalPath)
	if Conf.GlobalPath != "" {
		globalFile, err := os.Open(Conf.GlobalPath)
		if err != nil {
			panic(err)
		}
		globalFileText, err := ioutil.ReadAll(globalFile)
		globalFile.Close()
		Routers.Get("/**", func() (int, []byte) {
			return 200, globalFileText
		})
	}
}
