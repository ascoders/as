/*==================================================
	工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"math/rand"
	"newWoku/lib"
	"strconv"
	"time"
)

// 密码加密（两次md5）
// @params password string 加密的密码
func EncodePassword(password string) string {
	return lib.Md5(lib.Md5(password))
}

// 生成随机token
func CreateToken() string {
	return strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()))
}
