/*==================================================
	验证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package validation

import (
	"fmt"
)

type Valid struct {
	hasError bool
}

// 字符串不能为空
func (this *Valid) Required(value string) *Valid {
	if this.hasError {
		return this
	}
	this.hasError = value == ""
	return this
}

// 最小值
func (this *Valid) Min(value int, min int) *Valid {
	if this.hasError {
		return this
	}
	fmt.Println("判断了")
	this.hasError = value >= min
	return this
}

// 最大值
func (this *Valid) Max(value int, max int) *Valid {
	if this.hasError {
		return this
	}
	this.hasError = value <= max
	return this
}

// 范围
func (this *Valid) Range(value int, min int, max int) *Valid {
	if this.hasError {
		return this
	}
	this.Min(value, min).Max(value, max)
	return this
}

// 是否出错
func (this *Valid) HasError() bool {
	return this.hasError
}
