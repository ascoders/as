/*==================================================
	用户

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"newWoku/controllers"
	"newWoku/models"
)

type User struct {
	controllers.Base
}

func (this *User) Before() {
	this.NewModel(models.NewUser())
}

// @router /example
func (this *User) Other1() []byte {
	return this.Success("bb!")
}

// @router /example/:id
func (this *User) Other2() []byte {
	return this.Success("bb!")
}

// @router /example [post,put]
func (this *User) Other3() []byte {
	return this.Success("bb!")
}

// @router /example (csrf,before) [delete]
func (this *User) Other4() []byte {
	return this.Success("bb!")
}

// @router /example (csrf)
func (this *User) Other5() []byte {
	return this.Success("bb!")
}
