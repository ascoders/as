package response

import (
	"github.com/martini-contrib/encoder"
)

type Format struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 200, encoder.Must(enc.Encode(Format{true, "", data}))
}

func Error(message string) (int, []byte) {
	enc := encoder.JsonEncoder{}
	return 200, encoder.Must(enc.Encode(Format{false, message, nil}))
}
