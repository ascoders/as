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

type User struct {
	controllers.Base
}

func New() *User {
	controller := &User{}
	controller.NewModel(models.NewUser())
	return controller
}

// @router /users/:id (csrf) [get]
func (this *User) Get(param martini.Params) []byte {
	return this.Restful.Get(param)
}

// @router /example
func (this *User) Other1() []byte {
	return this.Success("bb!")
}

// @router /example/:id
func (this *User) Other2() []byte {
	return this.Success("bb!")
}

// @router /example [POST,pUT]
func (this *User) Other3() []byte {
	return this.Success("bb!")
}

// @router /example (csrf) [DELETE]
func (this *User) Other4() []byte {
	return this.Success("bb!")
}

// @router /example (csrf)
func (this *User) Other5() []byte {
	return this.Success("bb!")
}
