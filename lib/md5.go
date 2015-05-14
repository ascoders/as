/*==================================================
	md5加密

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}
