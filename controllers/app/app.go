/*==================================================
	应用

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package app

import (
	"newWoku/controllers"
	"newWoku/models/app"
)

type Controller struct {
	controllers.Base
}

func New() *Controller {
	controller := &Controller{}
	controller.NewModel(app.New())
	return controller
}

// @router /app/xx [Get]
func (this *Controller) Other() (int, []byte) {
	return this.Success("bb!")
}
