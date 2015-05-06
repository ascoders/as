/*==================================================
	统一输出接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package response

import (
	"github.com/martini-contrib/encoder"
)

type Format struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) []byte {
	enc := encoder.JsonEncoder{}
	return encoder.Must(enc.Encode(Format{true, "", data}))
}

func Error(message string) []byte {
	enc := encoder.JsonEncoder{}
	return encoder.Must(enc.Encode(Format{false, message, nil}))
}
