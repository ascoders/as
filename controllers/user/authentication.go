/*==================================================
	登陆认证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"math/rand"
	"net/http"
	"newWoku/lib/email"
	"newWoku/lib/redis"
	"newWoku/models/user"
	"strconv"
	"time"
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

	// 获得安全令牌
	// 先生成随机数作为token
	token := strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()))
	expire := 30
	sign, expireUnix := CreateSign(token, expire, map[string]string{
		"email":    user.Email,
		"nickname": user.Nickname,
		"password": user.Password,
	})

	// 保存有效令牌到缓存
	redis.SetWithExpire(token, []byte(sign), int64(expire))

	// 发送邮件
	go email.Send([]string{user.Email}, "我酷游戏：请激活账号", "点击下面链接 "+
		"http://wokugame.com/register?"+
		"expire="+expireUnix+"&"+
		"token="+token+"&"+
		"email="+user.Email+"&"+
		"nickname="+user.Nickname+"&"+
		"password="+user.Password)

	return this.Success("")
}

// 注册
// 验证邮箱令牌，并注册用户
// @router /users/authentication/email [post]
func (this *Controller) CreateEmailAuthentication(req *http.Request) (int, []byte) {
	return this.Success("ok")
}
