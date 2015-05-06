/*==================================================
	复写http.ResponseWriter

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

// 写入res.body之前，先写入缓存
func (this *ResponseWriter) Write(c []byte) (int, error) {
	redis.Set("url-"+this.Req.URL.String(), c)

	return this.Res.Write(c)
}

func (this *ResponseWriter) WriteHeader(h int) {
	this.Res.WriteHeader(h)
}
