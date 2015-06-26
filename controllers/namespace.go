package controllers

import (
	_http "github.com/ascoders/as/lib/http"
	"github.com/ascoders/as/lib/parse"
	"github.com/ascoders/as/lib/response"
	"net/http"
)

// 快捷方法

func (this *Base) Parse(obj interface{}, params map[string]string) error {
	return parse.ParseInstance.Struct(obj, params)
}

func (this *Base) ParseToUpdateMap(obj interface{}, params map[string]string) (error, map[string]interface{}) {
	return parse.ParseInstance.StructToUpdateMap(obj, params)
}

func (this *Base) ReqFormToMap(req *http.Request, limit ...string) map[string]string {
	return _http.HttpInstance.ReqFormToMap(req, limit...)
}

func (this *Base) Success(data interface{}) (int, []byte) {
	return response.ResponseInstance.Success(data)
}

func (this *Base) Error(message string) (int, []byte) {
	return response.ResponseInstance.Error(message)
}

func (this *Base) Must(data interface{}, err error) (int, []byte) {
	return response.ResponseInstance.Must(data, err)
}
