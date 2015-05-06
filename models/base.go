/*==================================================
	基础模型

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package models

type Base struct{}

type BaseModel interface {
	Gets() []*BaseModel
	Get(id int) *BaseModel
	Add(this *BaseModel) error
	Update(this *BaseModel) error
	Delete(id int) error
}

func (this *Base) Gets() []*BaseModel {
	var r []*BaseModel

	return r
}

func (this *Base) Get(id int) *BaseModel {
	var r *BaseModel

	return r
}

func (this *Base) Add() error {
	return nil
}

func (this *Base) Update() error {
	return nil
}

func (this *Base) Delete(id int) error {
	return nil
}
