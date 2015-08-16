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

func (this *Restful) Gets(req *http.Request, where map[string]interface{}, selecter []string) (int, []byte) {
	req.ParseForm()
	page, _ := strconv.Atoi(req.Form.Get("page"))
	limit, _ := strconv.Atoi(req.Form.Get("limit"))

	datas, err := this.Model.Gets(page, limit, where, selecter)

	// 查询总数
	count := this.Model.Count(where)

	return response.ResponseInstance.Must(map[string]interface{}{
		"list":  datas,
		"count": count,
	}, err)
}

func (this *Restful) Get(param martini.Params) (int, []byte) {
	return response.ResponseInstance.Must(this.Model.Get(param["id"]))
}

func (this *Restful) Add(req *http.Request) (int, []byte) {
	params := _http.HttpInstance.ReqFormToMap(req)
	data := this.Model.NewData()

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

	// 参数解析到结构体
	if err := parse.ParseInstance.Struct(data, params); err != nil {
		return response.ResponseInstance.Error(err.Error())
	}

	err := this.Model.Update(param["id"], data)
	return response.ResponseInstance.Must("更新成功", err)
}

func (this *Restful) Delete(params martini.Params) (int, []byte) {
	return response.ResponseInstance.Must("删除成功", this.Model.Delete(params["id"]))
}
