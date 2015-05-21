/*==================================================
	http工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	"net/http"
)

// 解析请求参数到map
func ReqFormToMap(req *http.Request, limit ...string) map[string]string {
	params := make(map[string]string)
	// 解析请求
	req.ParseForm()

ParseForm:
	for k, _ := range req.Form {
		// 如果指定了限制参数，只解析指定参数
		if len(limit) > 0 {
			for lk, _ := range limit {
				if k == limit[lk] {
					continue ParseForm
				}
			}
		}

		params[k] = req.Form.Get(k)
	}

	return params
}
