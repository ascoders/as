package models

import (
	"reflect"
	"strconv"
)

// 复制一个新实例
func (this *Base) NewData() interface{} {
	return reflect.New(reflect.TypeOf(this.Data).Elem()).Interface()
}

// 复制实例数组
func (this *Base) NewDatas() interface{} {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(this.Data)), 10, 10)
	datas := reflect.New(slice.Type())
	datas.Elem().Set(slice)
	return datas.Interface()
}

// 转化为int
func parseInt(number interface{}) int {
	result := 0
	switch number.(type) {
	case int: // good!
		result = number.(int)
	case string:
		result, _ = strconv.Atoi(number.(string))
	}

	return result
}
