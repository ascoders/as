/*==================================================
	用户表

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package app

import (
	"gopkg.in/mgo.v2/bson"
	"newWoku/models"
)

type Model struct {
	models.Base `bson:"-" json:"-"`
	Id          bson.ObjectId   `bson:"_id" json:"id"`        // 主键 也是讨论组英文路径
	Manager     bson.ObjectId   `bson:"m" json:"manager"`     // 管理员
	Helper      []bson.ObjectId `bson:"hl" json:"helper"`     // 助手（协助版主只能操作帖子，不能进管理台） !!!!!!! ms
	Name        string          `bson:"n" json:"name"`        // 名称
	Type        uint8           `bson:"t" json:"type"`        // 分类 1~4 智益休闲 养成RPG 竞技类 棋牌类 0为其他，不作为游戏
	ScreenShot  []string        `bson:"st" json:"screenShot"` // 截图 最多6个 !!!!!im
	Size        float32         `bson:"s" json:"size"`        // 应用大小
	Version     float32         `bson:"v" json:"version"`     // 版本
	Need        string          `bson:"nd" json:"need"`       // 系统要求
	Description string          `bson:"d" json:"description"` // 简介
	Download    string          `bson:"dl" json:"download"`   // 下载地址
	Image       string          `bson:"im" json:"image"`      // 游戏图标 !!!!!! gi
	Icon        string          `bson:"i" json:"icon"`        // 游戏icon图标地址
	Hot         int             `bson:"h" json:"hot"`         // 活跃度
	Categorys   int             `bson:"c" json:"categorys"`   // 分类数
	//UpdateTime  time.Time     `bson:"ut" json:""`           // 游戏介绍最后更新时间
	//Time time.Time `bson:"tm" json:""` // 成立时间
}

func New() *Model {
	model := &Model{Id: bson.NewObjectId()}
	model.Collection = models.Db.C("apps")
	return model
}

func (this *Model) NewObj() interface{} {
	var r *Model
	return &r
}

func (this *Model) NewSlice() interface{} {
	var r []*Model
	return &r
}
