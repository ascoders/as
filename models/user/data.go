/*==================================================
	用户表

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"newWoku/models"
	"time"
)

type Data struct {
	Id            bson.ObjectId `bson:"_id" json:"id" valid:"-"`                                       // 主键
	Nickname      string        `bson:"n" json:"nickname" valid:"required;minLength(1);maxLength(10)"` // 昵称
	Password      string        `bson:"p" json:"-" valid:"required;minLength(6);maxLength(30)"`        // 密码
	Email         string        `bson:"e" json:"email" valid:"required;email"`                         // 电子邮箱
	Image         string        `bson:"i" json:"image"`                                                // 头像地址
	Money         float32       `bson:"mo" json:"money" valid:"-"`                                     // 账户余额
	Free          int           `bson:"f" json:"free" valid:"-"`                                       // 每月免费额度 !!!!!!!!mf
	LogCount      uint16        `bson:"l" json:"logCount" valid:"-"`                                   // 登陆次数
	LastTime      time.Time     `bson:"la" json:"lastTime" valid:"-"`                                  // 最后操作时间
	ErrorChance   uint8         `bson:"er" json:"errorChance" valid:"-"`                               // 账号输错机会次数
	StopTime      time.Time     `bson:"st" json:"stopTime" valid:"-"`                                  // 账号封停截至时间
	Type          uint8         `bson:"t" json:"type" valid:"-"`                                       // 账号类型 0:超级管理员/董事长 1:会员 2:高级会员 3:白金会员
	Power         []string      `bson:"po" json:"power" valid:"-"`                                     // 模块权限
	UploadSize    int           `bson:"u" json:"uploadSize" valid:"-"`                                 // 今天上传大小
	UploadTime    time.Time     `bson:"ut" json:"uploadTime" valid:"-"`                                // 最后上传文件的时间 !!!!!!!!ud
	LockVersion   int           `bson:"lv" json:"-" valid:"-"`                                         // 乐观锁
	HasOrder      bool          `bson:"h" json:"hasOrder" valid:"-"`                                   // 是否有未处理的账单
	Token         string        `bson:"tk" json:"token" valid:"-"`                                     // 每个账号的密钥
	MessageNumber uint16        `bson:"mn" json:"messageNumber" valid:"-"`                             // 未读消息数量
	MessageAll    uint16        `bson:"ma" json:"messageAll" valid:"-"`                                // 总消息数
	AppCount      uint8         `bson:"a" json:"appCount" valid:"-"`                                   // 建立应用数量 !!!!!!!!!g
}

func New() *Model {
	model := &Model{}
	model.Collection = models.Db.C("users")

	if err := model.Collection.EnsureIndex(mgo.Index{
		Key:    []string{"e"},
		Unique: true,
	}); err != nil {
		panic(err)
	}
	return model
}

func (this *Model) NewData() interface{} {
	var r Data
	return &r
}

func (this *Model) NewDataWithId() interface{} {
	var r Data
	r.Id = bson.NewObjectId()
	return &r
}

func (this *Model) NewDatas() interface{} {
	var r []*Data
	return &r
}
