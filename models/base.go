/*==================================================
	基础模型

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"newWoku/conf"
	"time"
)

var (
	Db *mgo.Database // 数据库连接池
)

func init() {
	//获取数据库连接
	session, err := mgo.Dial(conf.MONGODB_ADDRESS)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	Db = session.DB("woku")
}

type Base struct {
	Id bson.ObjectId   `bson:"_id" json:"id"` // 主键
	db *mgo.Collection // 数据库连接
}

type BaseModel interface {
	Gets() []*BaseModel
	Get(id interface{})
	Add(this *BaseModel) error
	Update(this *BaseModel) error
	Delete(id interface{}) error
}

func (this *Base) Gets() []*BaseModel {
	var r []*BaseModel

	return r
}

func (this *Base) Get(id interface{}) {
	this.db.FindId(id).One(this)
}

func (this *Base) Add() error {
	this.Id = bson.NewObjectId()
	return this.db.Insert(this)
}

func (this *Base) Update() error {
	return nil
}

func (this *Base) Delete(id interface{}) error {
	return this.db.RemoveId(id)
}
