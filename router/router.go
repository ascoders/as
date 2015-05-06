package router

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
	"newWoku/conf"
	"newWoku/controllers/article"
	"newWoku/controllers/user"
	"newWoku/lib/autoRouter"
	"os"
)

// 设置路由
func Route() martini.Router {
	r := martini.NewRouter()

	// 用户操作
	u := user.User{}
	r.Get(`/user`, u.Get)
	/*
		r.Get(`/user/:id`, user.User.Get)
		r.Post(`/user`, user.User.Add)
		r.Put(`/user/:id`, user.User.Update)
		r.Delete(`/user/:id`, user.User.Delete)
	*/

	// 生成注解路由
	autoRouter.Parse(
		&user.User{},
		&article.Article{},
	)

	// 加入注解路由
	AutoRoute(r)

	// 最后匹配的是全局内容
	globalFile, err := os.Open(conf.GLOBAL_PATH)
	if err != nil {
		panic(err)
	}
	globalFileText, err := ioutil.ReadAll(globalFile)
	globalFile.Close()

	r.Get("/**", func(w http.ResponseWriter) []byte {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return globalFileText
	})

	return r
}
