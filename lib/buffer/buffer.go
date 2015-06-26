/*==================================================
	字符处理
	[]byte <=> interface{}

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package buffer

import (
	"bytes"
	"encoding/gob"
)

type Buffer struct{}

var (
	BufferInstance *Buffer
)

func init() {
	BufferInstance = &Buffer{}
}

// 编码为字节流
func (this *Buffer) Encode(data interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return []byte("")
	}
	return buf.Bytes()
}

// 字节流解码
func (this *Buffer) Decode(data []byte) interface{} {
	var r interface{}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	dec.Decode(r)
	return r
}
