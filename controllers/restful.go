/*==================================================
	基础控制器 - Restful api

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"github.com/go-martini/martini"
	"net/http"
	"newWoku/lib/model"
	"newWoku/lib/response"
	"newWoku/models"
	"strconv"
)

type Restful struct {
	Model models.BaseModel
}

func (this *Restful) Gets(req *http.Request) []byte {
	objs := this.Model.NewSlice()

	req.ParseForm()
	lastId := req.Form.Get("lastId")
	page, _ := strconv.Atoi(req.Form.Get("page"))
	// 查询数量
	number, _ := strconv.Atoi(req.Form.Get("number"))

	// 优先使用lastId查询
	if page > 0 && lastId == "" {
		err := this.Model.GetsByPage(page, number, objs)
		return response.Must(objs, err)
	} else {
		err := this.Model.GetsById(lastId, number, objs)
		return response.Must(objs, err)
	}
}

func (this *Restful) Get(param martini.Params) []byte {
	obj := this.Model.New()
	err := this.Model.Get(param["id"], obj)
	return response.Must(obj, err)
}

func (this *Restful) Add(req *http.Request) []byte {
	// 参数解析到结构体
	if err := model.Parse(this.Model, req); err != nil {
		return response.Error(err.Error())
	}

	if err := this.Model.Add(this.Model); err != nil {
		return response.Error(err.Error())
	}

	return response.Success(this.Model)
}

func (this *Restful) Update(param martini.Params, req *http.Request) []byte {
	opts := model.ParseTo(this.Model, req)
	err := this.Model.Update(param["id"], opts)
	return response.Must("", err)
}

func (this *Restful) Delete(params martini.Params) []byte {
	return response.Must(nil, this.Model.Delete(params["id"]))
}
