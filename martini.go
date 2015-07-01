package as

import (
	"github.com/ascoders/as/lib/csrf"
	"github.com/ascoders/as/lib/redis"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strings"
)

func newClassic() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()

	// 日志
	if Conf.Debug {
		m.Use(martini.Logger())
	}

	// 捕获所有错误
	m.Use(martini.Recovery())

	// 静态处理
	m.Use(martini.Static(Conf.StaticDir, Conf.StaticOptions))

	// Session
	store, _ := sessions.NewRediStore(10, "tcp", Conf.RedisAddress, "", []byte(Conf.SessionSecret))
	store.Options(Conf.SessionOptions)
	m.Use(sessions.Sessions(Conf.SessionName, store))

	// csrf
	m.Use(csrf.CsrfInstance.Generate(Conf.CsrfOptions))

	// 缓存中间件
	m.Use(func(c martini.Context, req *http.Request, w http.ResponseWriter) {
		// 覆写ResponseWriter接口
		res := Lib.Http.NewResponseWriter(req, w)
		c.MapTo(res, (*http.ResponseWriter)(nil))

		// Api请求
		if strings.HasPrefix(req.URL.String(), "/api") {
			// 响应类型：Json
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			// GET请求读取缓存
			if !Conf.Debug && req.Method == "GET" {
				// 缓存没过期
				if cache, err := redis.RedisInstance.Get("url-" + req.URL.String()); err == nil {
					w.Write(cache)
					return
				}
			}
		} else {
			// 其他类型请求响应类型：Html
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}
	})

	// 返回可以直接调用的路由
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	return &martini.ClassicMartini{m, r}
}
