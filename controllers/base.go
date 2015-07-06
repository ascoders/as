/*==================================================
	基础控制器

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package controllers

import (
	"github.com/ascoders/as/models"
)

type Base struct {
	Restful    // restful api
	Model      models.BaseModel
	UseRestful bool // 是否自动启用restful路由
}

func (this *Base) NewModel(model models.BaseModel) {
	this.Model = model
	this.Restful.Model = this.Model
}
