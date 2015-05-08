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
	//"newWoku/lib/redis"
	"newWoku/router"
	"strconv"
	"strings"
)

func main() {
	m := martini.Classic()
	m.Use(martini.Recovery())
	m.Use(martini.Static(conf.STATIC_DIR))

	// Session
	// store, _ := sessions.NewRediStore(1024, "", "127.0.0.1", "")
	// m.Use(sessions.Sessions("sessionStore", store))

	m.Use(func(c martini.Context, req *http.Request, w http.ResponseWriter) {
		// 覆写ResponseWriter接口
		res := _http.NewResponseWriter(req, w)
		c.MapTo(res, (*http.ResponseWriter)(nil))

		// Api请求
		if strings.HasPrefix(req.URL.String(), "/api") {
			// 响应类型：Json
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			/*
				// GET请求读取缓存
				if req.Method == "GET" {
					// 缓存没过期
					if cache, err := redis.Get("url-" + req.URL.String()); err == nil {
						w.Write(cache)
						return
					}
				}
			*/
		} else {
			// 其他类型请求响应类型：Html
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}
	})

	// 启用路由
	m.Action(router.Route().Handle)

	// 监听端口
	m.RunOnAddr(":" + strconv.Itoa(int(conf.PORT)))
}
