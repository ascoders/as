/*==================================================
	http工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	"errors"
	_http "net/http"
)

// 解析请求参数到map
// @params limits ...string 仅解析提供的参数
func (this *Http) ReqFormToMap(req *_http.Request, limits ...string) map[string]interface{} {
	params := make(map[string]interface{})
	// 解析请求
	req.ParseForm()

	if len(limits) == 0 {
		for k, _ := range req.Form {
			params[k] = translateBool(req.Form.Get(k))
		}
	} else {
		for k, _ := range limits {
			_, ok := req.Form[limits[k]]
			// 参数在req里存在才放进来
			if ok {
				params[limits[k]] = translateBool(req.Form.Get(limits[k]))
			}
		}
	}

	return params
}

// true->1 false->1
func translateBool(value string) interface{} {
	switch value {
	case "true":
		return 1
	case "false":
		return 0
	}
	return value
}

// 判断非空参数
func (this *Http) Require(req *_http.Request, keys ...string) error {
	// 解析请求
	req.ParseForm()

	for k, _ := range keys {
		if req.Form.Get(keys[k]) == "" {
			return errors.New(keys[k] + "不能为空")
		}
	}

	return nil
}
