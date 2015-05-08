/*==================================================
	基础模型

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"newWoku/conf"
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
	Add() error
	Update(id interface{}, update map[string]interface{}) error
	Delete(id interface{}) error
}

func (this *Base) Gets() []*BaseModel {
	var r []*BaseModel

	fmt.Println("执行了models.Gets")

	return r
}

// 根据id获取某个资源
func (this *Base) Get(id interface{}) {
	this.db.FindId(id).One(this)
}

// 新增资源
func (this *Base) Add() error {
	this.Id = bson.NewObjectId()
	return this.db.Insert(this)
}

// 根据id更新某个资源
func (this *Base) Update(id interface{}, update map[string]interface{}) error {
	return this.db.UpdateId(id, update)
}

// 根据id删除某个资源
func (this *Base) Delete(id interface{}) error {
	return this.db.RemoveId(id)
}
