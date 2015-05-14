/*==================================================
	用户

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"github.com/go-martini/martini"
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

// @router /users/:id (csrf) [get]
func (this *Controller) Get(param martini.Params) []byte {
	return this.Restful.Get(param)
}

// @router /example
func (this *Controller) Other1() []byte {
	return this.Success("bb!")
}

// @router /example/:id
func (this *Controller) Other2() []byte {
	return this.Success("bb!")
}

// @router /example [POST,pUT]
func (this *Controller) Other3() []byte {
	return this.Success("bb!")
}

// @router /example (csrf) [DELETE]
func (this *Controller) Other4() []byte {
	return this.Success("bb!")
}

// @router /example (csrf)
func (this *Controller) Other5() []byte {
	return this.Success("bb!")
}
