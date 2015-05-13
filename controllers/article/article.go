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

type Article struct {
	controllers.Base
}

func New() *Article {
	controller := &Article{}
	controller.NewModel(models.NewUser())
	return controller
}

// @router /aa [Get]
func (this *Article) Other() []byte {
	return response.Success("Delete success!")
}
