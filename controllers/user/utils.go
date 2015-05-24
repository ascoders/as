/*==================================================
	工具

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"errors"
	"net/http"
	"newWoku/lib"
	"newWoku/lib/sort"
	"strconv"
	"time"
)

// 生成签名
// @return sign
// @return expire
func CreateSign(token string, expire int, params map[string]string) (string, string) {
	params["expire"] = strconv.Itoa(int(time.Now().Unix()) + expire)
	params["token"] = token

	// 将参数排序
	opts := sort.MapToSlice(params)

	var text string
	for k, _ := range opts {
		text += opts[k]
	}

	return lib.Md5(text), params["expire"]
}

// 校验签名
func CheckSign(token string, req *http.Request) error {
	req.ParseForm()
	req.Form.Set("token", token)

	// 是否过期
	expire, _ := strconv.Atoi(req.Form.Get("expire"))
	if time.Unix(int64(expire), 0).Before(time.Now()) {
		return errors.New("请求已过期")
	}

	// 剔除sign参数后字母从小到大排序
	opts := sort.FormToSlice(req, "sign")

	// 对比sign
	var text string
	for k, _ := range opts {
		text += opts[k]
	}
	if req.Form.Get("sign") != lib.Md5(text) {
		return errors.New("签名校验失败")
	}

	return nil
}
