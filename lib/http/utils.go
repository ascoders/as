/*==================================================
	http工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	"net/http"
)

// 解析请求参数到map
func ReqFormToMap(req *http.Request) map[string]string {
	params := make(map[string]string)
	// 解析请求
	req.ParseForm()

	for k, _ := range req.Form {
		params[k] = req.Form.Get(k)
	}

	return params
}
