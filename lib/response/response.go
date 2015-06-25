/*==================================================
	统一输出接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package response

import (
	"github.com/martini-contrib/encoder"
)

type Response struct{}

func (this *Response) Success(data interface{}) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 200, encoder.Must(enc.Encode(data))
}

func (this *Response) Error(message string) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 400, encoder.Must(enc.Encode(map[string]interface{}{
		"message": message,
	}))
}

func (this *Response) Must(data interface{}, err error) (int, []byte) {
	if err == nil {
		return this.Success(data)
	} else {
		return this.Error(err.Error())
	}
}
