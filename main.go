/*==================================================
	程序入口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
	"newWoku/conf"
	"newWoku/lib/csrf"
	_http "newWoku/lib/http"
	"newWoku/lib/redis"
	"newWoku/router"
	"strconv"
	"strings"
)

func NewClassic() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()

	// 日志
	if conf.DEBUG {
		m.Use(martini.Logger())
	}

	// 捕获所有错误
	m.Use(martini.Recovery())

	// 静态处理
	m.Use(martini.Static(conf.STATIC_DIR, martini.StaticOptions{
		Prefix:      "/static/",
		SkipLogging: true,
	}))

	// Session
	store, _ := sessions.NewRediStore(10, "tcp", "127.0.0.1:6379", "")
	m.Use(sessions.Sessions("sessionStore", store))

	// csrf防范
	m.Use(csrf.Generate(&csrf.Options{
		Secret:     "V&2Xa6IAKZg5QjX8",
		SessionKey: "id", // 根据用户id，为每个用户设置不同的csrf
		Header:     "X-CSRFToken",
		SetCookie:  true,
		Cookie:     "_csrf", // cookie里设置csrf以便客户端取值
		ErrorFunc: func(w http.ResponseWriter) {
			w.Header().Del("X-Csrftoken")
			http.Error(w, "Bad Request", 400)
		},
	}))

	// 缓存中间件
	m.Use(func(c martini.Context, req *http.Request, w http.ResponseWriter) {
		// 覆写ResponseWriter接口
		res := _http.NewResponseWriter(req, w)
		c.MapTo(res, (*http.ResponseWriter)(nil))

		// Api请求
		if strings.HasPrefix(req.URL.String(), "/api") {
			// 响应类型：Json
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			// GET请求读取缓存
			if !conf.DEBUG && req.Method == "GET" {
				// 缓存没过期
				if cache, err := redis.Get("url-" + req.URL.String()); err == nil {
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

func main() {
	m := NewClassic()

	// 启用路由
	m.Action(router.Route().Handle)

	// 监听端口
	m.RunOnAddr(":" + strconv.Itoa(int(conf.PORT)))
}
