package models

import (
	"time"
)

type Member struct {
	Base

	Nickname string  `bson:"n" json:"nickname"` // 昵称
	Password string  `bson:"p" json:"password"` // 密码
	Email    string  `bson:"e" json:"email"`    // 电子邮箱
	Money    float32 `bson:"mo" json:"money"`   // 账户余额
	Free     int     `bson:"f" json:"free"`     // 每月免费额度 !!!!!!!!mf

	LogCount uint16    `bson:"l" json:"logCount"`  // 登陆次数
	LastTime time.Time `bson:"la" json:"lastTime"` // 最后操作时间

	ErrorChance uint8     `bson:"er" json:"errorChance"` // 账号输错机会次数
	StopTime    time.Time `bson:"st" json:"stopTime"`    // 账号封停截至时间

	Type  uint8    `bson:"t" json:"type"`   // 账号类型 0:超级管理员/董事长 1:会员 2:高级会员 3:白金会员
	Power []string `bson:"po" json:"power"` // 模块权限

	UploadSize int64     `bson:"u" json:"uploadSize"`  // 今天上传大小
	UploadTime time.Time `bson:"ut" json:"uploadTime"` // 最后上传文件的时间 !!!!!!!!ud

	LockVersion   uint64 `bson:"lv" json:"lockVersion"`   // 乐观锁
	HasOrder      bool   `bson:"h" json:"hasOrder"`       // 是否有未处理的账单
	Token         string `bson:"tk" json:"token"`         // 每个账号的密钥
	Image         string `bson:"i" json:"image"`          // 头像地址
	MessageNumber uint16 `bson:"mn" json:"messageNumber"` // 未读消息数量
	MessageAll    uint16 `bson:"ma" json:"messageAll"`    // 总消息数
	AppCount      uint8  `bson:"a" json:"appCount" `      // 建立应用数量 !!!!!!!!!g
}
