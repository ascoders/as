/*==================================================
	基础控制器 - 工具方法

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"net/http"
	"newWoku/lib/model"
	"newWoku/lib/response"
)

func (this *Base) Parse(obj interface{}, req *http.Request) error {
	return model.Parse(obj, req)
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
