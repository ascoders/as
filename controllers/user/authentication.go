/*==================================================
	登陆认证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"net/http"
	"newWoku/models/user"
)

// 登陆（获取授权令牌）
// @router /users/authentication [get]
func (this *Controller) Authentication(req *http.Request) (int, []byte) {
	req.ParseForm()
	return this.Must(Model.Authentication(req.Form.Get("account"), req.Form.Get("password")))
}

// 注册（创建授权令牌）
// @router /users/authentication (captcha) [post]
func (this *Controller) CreateAuthentication(req *http.Request) (int, []byte) {
	// 解析表单
	user := &user.Data{}
	if err := this.Parse(user, req); err != nil {
		return this.Error(err.Error())
	}
	// 验证数据
	Model.Validation(user)
	//return this.Must(Model.Authentication(req.Form.Get("account"), req.Form.Get("password")))
	return this.Success("123")
}
