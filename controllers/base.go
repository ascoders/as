/*==================================================
	基础控制器

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"net/http"
	"newWoku/lib/response"
)

type Base struct{}

type BaseController interface {
	Before(w http.ResponseWriter) []byte
	Gets() []byte
	Get() []byte
	Add() []byte
	Update() []byte
	Delete() []byte
	Count() []byte
}

// 逻辑之前执行
// 子类复写后可以做公共初始化或验证
// w.write(),之后逻辑路由不会执行
func (this *Base) Before(w http.ResponseWriter) {}

// @router [Get]
func (this *Base) Gets() []byte {
	return response.Success("Gets success!")
}

// @router [Get]
func (this *Base) Get() []byte {
	return response.Success("Get success!")
}

// @router [Post]
func (this *Base) Add() []byte {
	return response.Success("Add success!")
}

// @router [Put]
func (this *Base) Update() []byte {
	return response.Success("Update success!")
}

// @router [Delete]
func (this *Base) Delete() []byte {
	return response.Success("Delete success!")
}

func (this *Base) Count() []byte {
	return response.Success("Count")
}
