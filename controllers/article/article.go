/*==================================================
	文章

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package article

import (
	"newWoku/controllers"
	"newWoku/lib/response"
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

// @router /aa [Get]
func (this *Controller) Other() []byte {
	return response.Success("Delete success!")
}
