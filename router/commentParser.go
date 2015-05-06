/**
 * 注解路由
 */

package router

import (
	"github.com/go-martini/martini"
	"newWoku/controllers"
	"reflect"
)

// 解析controllers的注释，自动注册到路由
func CommentParse(r martini.Router, controls ...controllers.BaseController) {
	for _, control := range controls {
		reflectVal := reflect.ValueOf(control)

		in := make([]reflect.Value, 0)
		method := reflectVal.MethodByName("Get")
		method.Call(in)
		//r.Get("/user/:id", reflectVal.MethodByName("Get").Call([]reflect.Value{}))

		/*
			if comm, ok := GlobalControllerRouter[key]; ok {
				for _, a := range comm {
					p.Add(a.Router, control, strings.Join(a.AllowHTTPMethods, ",")+":"+a.Method)
				}
			}
		*/
	}
}
