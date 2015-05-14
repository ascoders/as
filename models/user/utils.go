/*==================================================
	工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"newWoku/lib"
)

// 密码加密（两次md5）
func EncodePassword(password string) string {
	return lib.Md5(lib.Md5(password))
}
