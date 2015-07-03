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

// 逻辑之前执行
// 子类复写后可以做公共初始化或验证
// w.write(),之后逻辑路由不会执行
func (this *Base) Before() {}
