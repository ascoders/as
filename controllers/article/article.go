/*==================================================
	文章

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

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
