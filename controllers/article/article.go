package article

import (
	"newWoku/controllers"
	"newWoku/lib/response"
)

type Article struct {
	controllers.Base
}

// @router /aa [Get]
func (this *Article) Other() []byte {
	return response.Success("Delete success!")
}
