package article

import (
	"newWoku/controllers"
	"newWoku/lib/response"
	"newWoku/models/user"
)

type Controller struct {
	controllers.Base
}

func New() *Controller {
	controller := &Controller{}
	controller.NewModel(user.New())
	return controller
}
