/*==================================================
	基础控制器 - 工具方法

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"net/http"
	_http "newWoku/lib/http"
	"newWoku/lib/model"
	"newWoku/lib/response"
)

func (this *Base) Parse(obj interface{}, params map[string]string) error {
	return model.Parse(obj, params)
}

func (this *Base) ParseToUpdateMap(obj interface{}, params map[string]string) (error, map[string]interface{}) {
	return model.ParseToUpdateMap(obj, params)
}

func (this *Base) ReqFormToMap(req *http.Request, limit ...string) map[string]string {
	return _http.ReqFormToMap(req, limit)
}

func (this *Base) Success(data interface{}) (int, []byte) {
	return response.Success(data)
}

func (this *Base) Error(message string) (int, []byte) {
	return response.Error(message)
}

func (this *Base) Must(data interface{}, err error) (int, []byte) {
	return response.Must(data, err)
}
