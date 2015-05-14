/*==================================================
	用户

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"github.com/go-martini/martini"
	"newWoku/controllers"
	"newWoku/models/user"
)

type Controller struct {
	controllers.Base
}

var (
	Model *user.Model
)

func New() *Controller {
	controller := &Controller{}
	Model = user.New()
	controller.NewModel(Model)
	return controller
}

// @router /users/:id (csrf) [get]
func (this *Controller) Get(param martini.Params) (int, []byte) {
	return this.Restful.Get(param)
}
