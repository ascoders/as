/*==================================================
	应用

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package app

import (
	"newWoku/controllers"
	"newWoku/models"
)

type App struct {
	controllers.Base
}

func New() *App {
	controller := &App{}
	controller.NewModel(models.NewUser())
	return controller
}

// @router /app/xx [Get]
func (this *App) Other() []byte {
	return this.Success("bb!")
}
