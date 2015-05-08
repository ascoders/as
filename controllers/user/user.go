package user

import (
	"net/http"
	"newWoku/controllers"
	"newWoku/lib/response"
	"newWoku/models"
)

type User struct {
	controllers.Base
}

func (this *User) Before(w http.ResponseWriter) {
	this.Model = &models.User{}
}

// @router /bb [Get]
func (this *User) Other() []byte {
	this.Model.Gets()
	return response.Success("bb!")
}
