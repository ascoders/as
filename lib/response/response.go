/*==================================================
	统一输出接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package response

import (
	"github.com/martini-contrib/encoder"
)

func Success(data interface{}) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 200, encoder.Must(enc.Encode(data))
}

func Error(message string) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 400, encoder.Must(enc.Encode(map[string]interface{}{
		"message": message,
	}))
}

func Must(data interface{}, err error) (int, []byte) {
	if err == nil {
		return Success(data)
	} else {
		return Error(err.Error())
	}
}
