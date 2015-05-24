/*==================================================
	登陆认证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"net/http"
	//	"newWoku/lib/redis"
	"newWoku/models/user"
)

// 登陆（获取授权令牌）
// @router /users/authentication [get]
func (this *Controller) Authentication(req *http.Request) (int, []byte) {
	req.ParseForm()
	return this.Must(Model.Authentication(req.Form.Get("account"), req.Form.Get("password")))
}

// 注册（创建授权令牌）
// 并不会注册用户，只会发邮件
// @router /users/authentication (captcha) [post]
func (this *Controller) CreateAuthentication(req *http.Request) (int, []byte) {
	user := &user.Data{}
	params := this.ReqFormToMap(req)

	if err := this.Parse(user, params); err != nil {
		return this.Error(err.Error())
	}

	//	expire := 30
	// 获得安全令牌

	// 保存有效令牌到缓存
	// redis.SetWithExpire(token, []byte("213"), expire)

	// 发送邮件

	return this.Success("123")
}
