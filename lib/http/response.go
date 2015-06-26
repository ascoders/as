/*==================================================
	复写http.Response

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	"github.com/ascoders/as/lib/redis"
	_http "net/http"
	"strings"
)

type Http struct{}

var (
	HttpInstance *Http
)

func init() {
	HttpInstance = &Http{}
}

func (this *Http) NewResponseWriter(req *_http.Request, res _http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		Req: req,
		Res: res,
	}
}

type ResponseWriter struct {
	Req *_http.Request
	Res _http.ResponseWriter
}

func (this *ResponseWriter) Header() _http.Header {
	return this.Res.Header()
}

func (this *ResponseWriter) Write(c []byte) (int, error) {
	// GET请求写入缓存
	if strings.HasPrefix(this.Req.URL.String(), "/api") && this.Req.Method == "GET" {
		redis.RedisInstance.Set("url-"+this.Req.URL.String(), c)
	}

	return this.Res.Write(c)
}

func (this *ResponseWriter) WriteHeader(h int) {
	this.Res.WriteHeader(h)
}
