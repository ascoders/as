/*==================================================
	统一输出接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package response

import (
	"fmt"
	"github.com/martini-contrib/encoder"
)

type Response struct{}

var (
	ResponseInstance *Response
)

func init() {
	ResponseInstance = &Response{}
}

func (this *Response) Success(data interface{}) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 200, encoder.Must(enc.Encode(data))
}

func (this *Response) Error(message interface{}) (int, []byte) {
	fmt.Println(message)
	enc := encoder.JsonEncoder{}
	return 400, encoder.Must(enc.Encode(message))
}

func (this *Response) Must(data interface{}, err error) (int, []byte) {
	if err == nil {
		return this.Success(data)
	} else {
		return this.Error(err.Error())
	}
}

// 输出空
func (this *Response) Empty() (int, []byte) {
	return 200, nil
}
