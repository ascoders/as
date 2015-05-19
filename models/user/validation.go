/*==================================================
	验证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package user

import (
	"newWoku/lib/validation"
)

func (this *Model) Validation(data *Data) bool {
	valid := &validation.Valid{}

	return valid.HasError()
}
