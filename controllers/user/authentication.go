/*==================================================
	登陆认证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"net/http"
)

// 登陆（获取授权令牌）
// @router /users/authentication [get]
func (this *Controller) Authentication(res *http.Request) (int, []byte) {
	res.ParseForm()
	return this.Must(Model.Authentication(res.Form.Get("account"), res.Form.Get("password")))
}

// 注册（创建授权令牌）
// @router /users/authentication (captcha) [post]
func (this *Controller) CreateAuthentication(res *http.Request) (int, []byte) {
	//return this.Must(Model.Authentication(res.Form.Get("account"), res.Form.Get("password")))
	return this.Success("123")
}
