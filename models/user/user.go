/*==================================================
	用户表

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

// 授权令牌是否有效（用户登陆成功）
func (this *Model) Authentication(account string, password string) (*Model, error) {
	user := &Model{}
	// 账户锁定情况
	if bson.Now().Before(user.StopTime) {
		long := user.StopTime.Sub(bson.Now())
		return nil, errors.New("账号距离解锁还有 " + strconv.FormatFloat(long.Seconds(), 'f', 0, 64) + " 秒")
	}
	// 根据邮箱查找用户
	if err := this.Collection.Find(bson.M{"e": account}).One(user); err != nil {
		return nil, errors.New("账号不存在")
	}
	// 校验密码
	if user.Password != EncodePassword(password) {
		if user.ErrorChance == 0 {
			user.ErrorChance = 6
		}
		if user.ErrorChance <= 1 {
			// 如果尽验证机会，账号锁定10分钟
			minute := time.Duration(10) * time.Minute
			this.Collection.UpdateId(user.Id, bson.M{"$set": bson.M{"er": 6, "st": bson.Now().Add(minute)}})
			return nil, errors.New("为保障安全，您的账号在10分钟后解除锁定状态")
		} else {
			// 验证机会减少1次
			this.Collection.UpdateId(user.Id, bson.M{"$inc": bson.M{"er": -1}})
			return nil, errors.New("密码错误，您还有 " + strconv.Itoa(int(user.ErrorChance)) + " 次机会")
		}
	}
	// 重置验证次数
	this.Collection.UpdateId(user.Id, bson.M{"$set": bson.M{"er": 6}})
	return user, nil
}

// 创建一个授权令牌（用户注册）
func (this *Model) CreateAuthentication(account string, password string) (*Model, error) {
	return nil, nil
}
