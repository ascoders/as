/*==================================================
	应用

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package app

import (
	"newWoku/controllers"
	"newWoku/models"
)

type Controller struct {
	controllers.Base
}

func New() *Controller {
	controller := &Controller{}
	controller.NewModel(models.NewUser())
	return controller
}

// @router /app/xx [Get]
func (this *Controller) Other() []byte {
	return this.Success("bb!")
}
