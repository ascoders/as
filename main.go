package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"newWoku/conf"
	"newWoku/lib/scheduled"
	"newWoku/router"
	"strconv"
)

func main() {
	scheduled.Run(22, 32, 30, func() {
		fmt.Println("oh!")
	})

	m := martini.Classic()
	m.Use(martini.Recovery())
	m.Use(martini.Static(conf.STATIC_DIR))

	// 默认响应类型：Json
	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	// 路由
	m.Action(router.Route().Handle)

	m.RunOnAddr(":" + strconv.Itoa(int(conf.PORT)))
}
