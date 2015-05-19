/*==================================================
	基础控制器 - Restful api

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"fmt"
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

func (this *Restful) Gets(req *http.Request) (int, []byte) {
	datas := this.Model.NewDatas()
	fmt.Println(datas)

	req.ParseForm()
	lastId := req.Form.Get("lastId")
	page, _ := strconv.Atoi(req.Form.Get("page"))
	limit, _ := strconv.Atoi(req.Form.Get("limit"))

	// 优先使用lastId查询
	if page > 0 && lastId == "" {
		err := this.Model.GetsByPage(page, limit, datas)
		return response.Must(datas, err)
	} else {
		err := this.Model.GetsById(lastId, limit, datas)
		return response.Must(datas, err)
	}
}

func (this *Restful) Get(param martini.Params) (int, []byte) {
	data := this.Model.NewData()
	err := this.Model.Get(param["id"], data)
	return response.Must(data, err)
}

func (this *Restful) Add(req *http.Request) (int, []byte) {
	data := this.Model.NewDataWithId()
	// 参数解析到结构体
	if err := model.Parse(data, req); err != nil {
		return response.Error(err.Error())
	}

	if err := this.Model.Add(data); err != nil {
		return response.Error(err.Error())
	}

	return response.Success(data)
}

func (this *Restful) Update(param martini.Params, req *http.Request) (int, []byte) {
	data := this.Model.NewData()
	opts := model.ParseTo(data, req)
	err := this.Model.Update(param["id"], opts)
	return response.Must("更新成功", err)
}

func (this *Restful) Delete(params martini.Params) (int, []byte) {
	return response.Must("删除成功", this.Model.Delete(params["id"]))
}
