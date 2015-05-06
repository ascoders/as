package user

import (
	"newWoku/controllers"
	"newWoku/lib/response"
)

type User struct {
	controllers.Base
}

// @router /other [Get]
func (this *User) Other() []byte {
	return response.Success("Delete success!")
}

// @router /xxx [Post]
func (this *User) Xxxx() []byte {
	return response.Success("Delete success!")
}
