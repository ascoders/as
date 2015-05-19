/*==================================================
	用户表

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"newWoku/models"
	"strconv"
	"time"
)

type Model struct {
	models.Base
}

// 授权令牌是否有效（用户登陆成功）
func (this *Model) Authentication(account string, password string) (*Data, error) {
	user := &Data{}
	// 根据邮箱查找用户
	if err := this.Collection.Find(bson.M{"e": account}).One(user); err != nil {
		return nil, errors.New("账号不存在")
	}
	// 账户锁定情况
	if bson.Now().Before(user.StopTime) {
		long := user.StopTime.Sub(bson.Now())
		return nil, errors.New("账号距离解锁还有 " + strconv.FormatFloat(long.Seconds(), 'f', 0, 64) + " 秒")
	}
	// 校验密码
	if user.Password != EncodePassword(password) {
		if user.ErrorChance == 1 {
			// 如果尽验证机会，账号锁定10分钟
			minute := time.Duration(10) * time.Minute
			this.Collection.UpdateId(user.Id, bson.M{"$set": bson.M{"er": 6, "st": bson.Now().Add(minute)}})
			return nil, errors.New("为保障安全，您的账号在10分钟后解除锁定状态")
		} else {
			if user.ErrorChance == 0 {
				// 默认错误机会为0，重新把错误机会设置为(6-1)
				this.Collection.UpdateId(user.Id, bson.M{"$set": bson.M{"er": 5}})
				user.ErrorChance = 5
			} else {
				// 验证机会减少1次
				this.Collection.UpdateId(user.Id, bson.M{"$inc": bson.M{"er": -1}})
				user.ErrorChance--
			}
			return nil, errors.New("密码错误，您还有 " + strconv.Itoa(int(user.ErrorChance)) + " 次机会")
		}
	}
	// this.Token = strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()))
	// 重置验证次数
	this.Collection.UpdateId(user.Id, bson.M{"$set": bson.M{"er": 6}})
	return user, nil
}

// 创建一个授权令牌（用户注册）
func (this *Model) CreateAuthentication(account string, password string) (*Data, error) {
	user := &Data{}
	// 创建密码
	user.Password = EncodePassword(password)
	return nil, nil
}
