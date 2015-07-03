/*==================================================
	基础控制器

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"github.com/ascoders/as/models"
	"github.com/go-martini/martini"
	"net/http"
)

type Base struct {
	Restful    // restful api
	Model      models.BaseModel
	UseRestful bool // 是否自动启用restful路由
}

func (this *Base) NewModel(model models.BaseModel) {
	this.Model = model
	this.Restful.Model = this.Model
}

// 逻辑之前执行
// 子类复写后可以做公共初始化或验证
// w.write(),之后逻辑路由不会执行
func (this *Base) Before() {

}

func (this *Base) Gets(req *http.Request) (int, []byte) {
	return this.Restful.Gets(req)
}

func (this *Base) Get(param martini.Params) (int, []byte) {
	return this.Restful.Get(param)
}

func (this *Base) Add(req *http.Request) (int, []byte) {
	return this.Restful.Add(req)
}

func (this *Base) Update(param martini.Params, req *http.Request) (int, []byte) {
	return this.Restful.Update(param, req)
}

func (this *Base) Delete(params martini.Params) (int, []byte) {
	return this.Restful.Delete(params)
}
