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

func (this *App) Before() {
	this.NewModel(models.NewApp())
}

// @router /app/xx [Get]
func (this *App) Other() []byte {
	return this.Success("bb!")
}
