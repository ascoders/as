/*==================================================
	验证码

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package captcha

import (
	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	"net/http"
)

// 验证码
type Captcha struct {
	// ID
	CaptchaId string
	// 验证码
	Solution string
}

// 验证图片
func Image(params martini.Params, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/png")
	captcha.WriteImage(w, params["id"], 240, 80)
}

// 获取验证代码
func Code() string {
	return captcha.NewLen(6)
}

// 校验验证码
func Check(capid string, cap string) bool {
	if ok := captcha.VerifyString(capid, cap); !ok {
		return false
	}
	return true
}
