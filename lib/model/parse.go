/*==================================================
	模型 解析

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package model

import (
	"errors"
	"newWoku/lib/validation"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 解析url参数
// @param {interface{}} obj 被解析的结构体
// @param {http.Request} req 客户端请求
// 用于添加
func Parse(obj interface{}, params map[string]string) error {
	objT := reflect.TypeOf(obj).Elem()
	objV := reflect.ValueOf(obj).Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		// 从标签获json参数
		fieldT := objT.Field(i)
		tags := strings.Split(fieldT.Tag.Get("json"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			// 忽略：-
			continue
		} else {
			tag = tags[0]
		}

		// 从标签获取valid参数
		valids := strings.Split(fieldT.Tag.Get("valid"), ";")

		// 结构体的参数在提交参数不存在，则跳过
		value, ok := params[tag]
		if !ok {
			// 如果跳过的参数是required的，则返回一个错误
			if stringInSlice("required", valids) {
				return errors.New("缺少" + tag + "参数")
			}
			continue
		} else {
			// 参数存在，则valid不能存在-属性
			if stringInSlice("-", valids) {
				return errors.New(tag + "不可修改")
			}
		}

		// 自定义验证
		for k, _ := range valids {
			if err := validKey(valids[k], value); err != nil {
				return errors.New(tag + err.Error())
			}
		}

		// 解析到结构体
		switch fieldT.Type.Kind() {
		case reflect.Bool:
			if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
				fieldV.SetBool(true)
				continue
			}
			if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
				fieldV.SetBool(false)
				continue
			}
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		case reflect.Slice:
			// 数组型 :TODO
		case reflect.Struct:
			switch fieldT.Type.String() {
			case "time.Time":
				format := time.RFC3339
				if len(tags) > 1 {
					format = tags[1]
				}
				t, err := time.Parse(format, value)
				if err != nil {
					return err
				}
				fieldV.Set(reflect.ValueOf(t))
			}
		}
	}

	return nil
}

// 在Parse基础上返回用于更新的map
// @param {interface{}} obj 被解析的结构体
// @param {http.Request} req 客户端请求
// @return map[string]interface{}
// 用于更新
func ParseToUpdateMap(obj interface{}, params map[string]string, exempts ...string) (error, map[string]interface{}) {
	opts := make(map[string]interface{})
	objT := reflect.TypeOf(obj).Elem()
	objV := reflect.ValueOf(obj).Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		// 从标签获取到json字段名
		fieldT := objT.Field(i)
		tags := strings.Split(fieldT.Tag.Get("json"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			// 忽略：-
			continue
		} else {
			tag = tags[0]
		}

		// 获取bson字段
		bson := fieldT.Tag.Get("bson")

		// 跳过url不存在的参数
		value, ok := params[tag]

		// 跳过参数中结构体不存在的参数
		if !ok {
			continue
		}

		// 从标签获取valid参数
		valids := strings.Split(fieldT.Tag.Get("valid"), ";")

		// valid不能存在-属性（除非被豁免）
		if stringInSlice("-", valids) && !stringInSlice(tag, exempts) {
			return errors.New(tag + "不可修改"), nil
		}

		// 自定义验证
		for k, _ := range valids {
			if err := validKey(valids[k], value); err != nil {
				return errors.New(tag + err.Error()), nil
			}
		}

		// 解析到结构体
		switch fieldT.Type.Kind() {
		case reflect.Bool:
			if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
				opts[bson] = true
				fieldV.SetBool(true)
				continue
			}
			if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
				opts[bson] = false
				fieldV.SetBool(false)
				continue
			}
			b, err := strconv.ParseBool(value)
			if err != nil {
				return errors.New(tag + "类型错误"), nil
			}
			opts[bson] = b
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return errors.New(tag + "类型错误"), nil
			}
			opts[bson] = x
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return errors.New(tag + "类型错误"), nil
			}
			opts[bson] = x
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return errors.New(tag + "类型错误"), nil
			}
			opts[bson] = x
			fieldV.SetFloat(x)
		case reflect.Interface:
			opts[bson] = reflect.ValueOf(value)
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			opts[bson] = value
			fieldV.SetString(value)
		case reflect.Slice:
			// 数组型 :TODO
		case reflect.Struct:
			switch fieldT.Type.String() {
			case "time.Time":
				format := time.RFC3339
				if len(tags) > 1 {
					format = tags[1]
				}
				t, err := time.Parse(format, value)
				if err != nil {
					return errors.New(tag + "类型错误"), nil
				}
				opts[bson] = reflect.ValueOf(t)
				fieldV.Set(reflect.ValueOf(t))
			}
		}
	}

	return nil, opts
}

// 判断字符串是否在数组中
func stringInSlice(str string, array []string) bool {
	for _, v := range array {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// 验证
func validKey(key string, value string) error {
	// 排除空规则
	if key == "" {
		return nil
	}

	// 解析key
	var method string
	var number1, number2 int
	keySlice := strings.Split(key, "(")
	method = keySlice[0]

	if len(keySlice) >= 2 {
		paramSlice := strings.Split(strings.Trim(keySlice[1], ")"), ",")
		if len(paramSlice) >= 1 {
			number1, _ = strconv.Atoi(paramSlice[0])
		}
		if len(paramSlice) >= 2 {
			number2, _ = strconv.Atoi(paramSlice[1])
		}
	}

	valid := validation.Valid{}

	switch strings.ToLower(method) {
	case "required":
		return valid.Required(value)
	case "min":
		return valid.Min(value, number1)
	case "max":
		return valid.Max(value, number1)
	case "range":
		return valid.Range(value, number1, number2)
	case "minlength":
		return valid.MinLength(value, number1)
	case "maxlength":
		return valid.MaxLength(value, number1)
	case "length":
		return valid.Length(value, number1)
	case "alphanumeric":
		return valid.AlphaNumeric(value)
	case "email":
		return valid.Email(value)
	case "ip":
		return valid.IP(value)
	case "base64":
		return valid.Base64(value)
	case "mobile":
		return valid.Mobile(value)
	case "tel":
		return valid.Tel(value)
	case "mobileortel":
		return valid.MobileOrTel(value)
	case "zipcode":
		return valid.ZipCode(value)
	}

	return nil
}
