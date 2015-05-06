package article

import (
	"newWoku/controllers"
	"newWoku/lib/response"
)

type Article struct {
	controllers.Base
}

// @router /article [Get]
func (this *Article) Other() (int, []byte) {
	return response.Success("Delete success!")
}
