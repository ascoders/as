/*==================================================
	复写http.Response

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package http

import (
	"net/http"
	"newWoku/lib/redis"
)

func NewResponseWriter(req *http.Request, res http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		Req: req,
		Res: res,
	}
}

type ResponseWriter struct {
	Req *http.Request
	Res http.ResponseWriter
}

func (this *ResponseWriter) Header() http.Header {
	return this.Res.Header()
}

func (this *ResponseWriter) Write(c []byte) (int, error) {
	// GET请求写入缓存
	if this.Req.Method == "GET" && this.Res.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		redis.Set("url-"+this.Req.URL.String(), c)
	}

	return this.Res.Write(c)
}

func (this *ResponseWriter) WriteHeader(h int) {
	this.Res.WriteHeader(h)
}
