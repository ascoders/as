/*==================================================
	程序入口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package main

import (
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/sessions"
	"net/http"
	"newWoku/conf"
	_http "newWoku/lib/http"
	"newWoku/lib/redis"
	"newWoku/router"
	"strconv"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Recovery())
	m.Use(martini.Static(conf.STATIC_DIR))

	// 默认响应类型：Json
	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	// Session
	// store, _ := sessions.NewRediStore(1024, "", "127.0.0.1", "")
	// m.Use(sessions.Sessions("sessionStore", store))

	// 缓存
	m.Use(func(c martini.Context, req *http.Request, w http.ResponseWriter) {
		// 覆盖ResponseWriter接口
		res := _http.NewResponseWriter(req, w)
		c.MapTo(res, (*http.ResponseWriter)(nil))

		// GET请求读取缓存
		if req.Method == "GET" {
			// 缓存没过期
			if cache, err := redis.Get("url-" + req.URL.String()); err == nil {
				w.Write(cache)
				return
			}
		}
	})

	// 路由
	m.Action(router.Route().Handle)

	m.RunOnAddr(":" + strconv.Itoa(int(conf.PORT)))
}
