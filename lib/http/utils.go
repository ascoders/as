/*==================================================
	http工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	_http "net/http"
)

// 解析请求参数到map
// @params limit ...string 仅解析提供的参数
func ReqFormToMap(req *_http.Request, limit ...string) map[string]string {
	params := make(map[string]string)
	// 解析请求
	req.ParseForm()

	if len(limit) == 0 {
		for k, _ := range req.Form {
			params[k] = req.Form.Get(k)
		}
	} else {
		for k, _ := range limit {
			params[limit[k]] = req.Form.Get(limit[k])
		}
	}

	return params
}
