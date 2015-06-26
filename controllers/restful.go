/*==================================================
	基础控制器 - Restful api

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	_http "github.com/ascoders/as/lib/http"
	"github.com/ascoders/as/lib/parse"
	"github.com/ascoders/as/lib/response"
	"github.com/ascoders/as/models"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

type Restful struct {
	Model models.BaseModel
}

func (this *Restful) Gets(req *http.Request) (int, []byte) {
	datas := this.Model.NewDatas()
	req.ParseForm()
	lastId := req.Form.Get("lastId")
	page, _ := strconv.Atoi(req.Form.Get("page"))
	limit, _ := strconv.Atoi(req.Form.Get("limit"))

	// 优先使用lastId查询
	if page > 0 && lastId == "" {
		err := this.Model.GetsByPage(page, limit, datas)
		return response.ResponseInstance.Must(datas, err)
	} else {
		err := this.Model.GetsById(lastId, limit, datas)
		return response.ResponseInstance.Must(datas, err)
	}
}

func (this *Restful) Get(param martini.Params) (int, []byte) {
	data := this.Model.NewData()
	err := this.Model.Get(param["id"], data)

	return response.ResponseInstance.Must(data, err)
}

func (this *Restful) Add(req *http.Request) (int, []byte) {
	data := this.Model.NewDataWithId()
	params := _http.HttpInstance.ReqFormToMap(req)

	// 参数解析到结构体
	if err := parse.ParseInstance.Struct(data, params); err != nil {
		return response.ResponseInstance.Error(err.Error())
	}

	if err := this.Model.Add(data); err != nil {
		return response.ResponseInstance.Error(err.Error())
	}

	return response.ResponseInstance.Success(data)
}

func (this *Restful) Update(param martini.Params, req *http.Request) (int, []byte) {
	data := this.Model.NewData()
	params := _http.HttpInstance.ReqFormToMap(req)

	if err, opts := parse.ParseInstance.StructToUpdateMap(data, params); err == nil {
		err := this.Model.Update(param["id"], opts)
		return response.ResponseInstance.Must("更新成功", err)
	} else {
		return response.ResponseInstance.Error(err.Error())
	}
}

func (this *Restful) Delete(params martini.Params) (int, []byte) {
	return response.ResponseInstance.Must("删除成功", this.Model.Delete(params["id"]))
}
