/*==================================================
	基础控制器

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"github.com/go-martini/martini"
	"net/http"
	"newWoku/lib/response"
	"newWoku/models"
)

type BaseController interface {
	Before(w http.ResponseWriter) []byte
	Gets() []byte
	Get() []byte
	Add() []byte
	Update() []byte
	Delete() []byte
}

type Base struct {
	Model models.BaseModel
}

// 逻辑之前执行
// 子类复写后可以做公共初始化或验证
// w.write(),之后逻辑路由不会执行
func (this *Base) Before(w http.ResponseWriter) {

}

func (this *Base) Gets() []byte {
	this.Model.Gets()
	return response.Success("Gets success!")
}

func (this *Base) Get() []byte {
	return response.Success("Get success!")
}

func (this *Base) Add() []byte {
	return response.Success("Add success!")
}

func (this *Base) Update(params martini.Params) []byte {
	return response.Success("Update success!")
}

func (this *Base) Delete() []byte {
	return response.Success("Delete success!")
}
