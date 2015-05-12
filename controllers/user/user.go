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

// @router /bb [Get]
func (this *User) Other() []byte {
	return this.Success("bb!")
}
