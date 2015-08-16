/*==================================================
 基础模型

 Copyright (c) 2015 翱翔大空 and other contributors
==================================================*/

package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Base struct {
	Db   *gorm.DB
	Data interface{}
}

type BaseModel interface {
	Add(obj interface{}) error
	Gets(page int, limit int, where map[string]interface{}, selecter []string) (interface{}, error)
	Count(where map[string]interface{}) int
	Get(id interface{}) (interface{}, error)
	Update(id interface{}, update interface{}) error
	Delete(id interface{}) error
	NewData() interface{}
}

// 实现baseModel interface

// 新增资源
func (this *Base) Add(obj interface{}) error {
	return this.Db.Create(obj).Error
}

// 获取资源集
func (this *Base) Gets(page int, limit int, where map[string]interface{}, selecter []string) (interface{}, error) {
	if page == 0 {
		page = 1
	}

	if page < 1 || page > 100 {
		return nil, errors.New("页数在1-100之间")
	}

	if limit == 0 {
		limit = 10
	}

	if limit < 0 || limit > 100 {
		return nil, errors.New("批量查询数量在1-100之间")
	}

	objs := this.NewDatas()
	var err error

	if selecter == nil {
		err = this.Db.Find(objs, where).Error
	} else {
		err = this.Db.Select(selecter).Find(objs, where).Error
	}

	return objs, err
}

// 获取总数
func (this *Base) Count(where map[string]interface{}) int {
	var count int
	this.Db.Where(where).Count(&count)
	return count
}

// 获取某个资源
func (this *Base) Get(id interface{}) (interface{}, error) {
	obj := this.NewData()
	err := this.Db.First(obj, parseInt(id)).Error
	return obj, err
}

// 根据id更新某个资源
func (this *Base) Update(id interface{}, update interface{}) error {
	return this.Db.Where(map[string]interface{}{
		"id": parseInt(id),
	}).Update(update).Error
}

// 根据id删除某个资源
func (this *Base) Delete(id interface{}) error {
	return this.Db.Where(map[string]interface{}{
		"id": parseInt(id),
	}).Delete(this.Data).Error
}
