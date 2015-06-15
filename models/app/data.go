/*==================================================
	api表

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
	Id         bson.ObjectId `bson:"_id" json:"id" valid:"-"` // 主键
	Name       string        `bson:"n" json:"name"`           // 讨论组名称
	Path       string        `bson:"p" json:"path"`           // 英文路径
	Manager    bson.ObjectId `bson:"m" json:"manager"`        // 管理员
	Managers   []string      `bson:"ms" json:"managers"`      // 版主列表（协助版主只能操作帖子，不能进管理台）
	Type       uint8         `bson:"t" json:"type"`           // 分类 1~4 智益休闲 养成RPG 竞技类 棋牌类 0为其他，不作为游戏
	Logo       string        `bson:"l" json:"logo"`           // logo地址 ! gi
	Icon       string        `bson:"i" json:"icon"`           // icon地址
	Hot        int           `bson:"h" json:"hot"`            // 活跃度
	Categorys  int           `bson:"c" json:"categorys"`      // 分类数
	UpdateTime time.Time     `bson:"ut" json:"updateTime"`    // 游戏介绍最后更新时间
	Created    time.Time     `bson:"ct" json:"created"`       // 成立时间 !tm
}

func New() *Model {
	model := &Model{}
	model.Collection = models.Db.C("apps")

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
