package user

import (
	"newWoku/controllers"
	"newWoku/lib/response"
)

type User struct {
	controllers.Base
}

func (this *User) Other() (int, []byte) {
	return response.Success("Delete success!")
}
