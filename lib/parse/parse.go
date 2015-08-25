/*==================================================
	模型 解析

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package parse

import (
	"errors"
	"fmt"
	"github.com/ascoders/as/lib/validation"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Parse struct{}

var (
	ParseInstance *Parse
)

func init() {
	ParseInstance = &Parse{}
}

// 解析url参数
// @param {interface{}} obj 被解析的结构体
// @param {http.Request} req 客户端请求
func (this *Parse) Struct(obj interface{}, params map[string]interface{}) error {
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
		valueInterface, ok := params[tag]

		if !ok {
			continue
		}

		value := fmt.Sprint(valueInterface)

		// valid不能存在-属性
		if stringInSlice("-", valids) {
			return errors.New(tag + "不可修改")
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

	switch strings.ToLower(method) {
	case "required":
		return validation.ValidInstance.Required(value)
	case "min":
		return validation.ValidInstance.Min(value, number1)
	case "max":
		return validation.ValidInstance.Max(value, number1)
	case "range":
		return validation.ValidInstance.Range(value, number1, number2)
	case "minlength":
		return validation.ValidInstance.MinLength(value, number1)
	case "maxlength":
		return validation.ValidInstance.MaxLength(value, number1)
	case "length":
		return validation.ValidInstance.Length(value, number1)
	case "alphanumeric":
		return validation.ValidInstance.AlphaNumeric(value)
	case "email":
		return validation.ValidInstance.Email(value)
	case "ip":
		return validation.ValidInstance.IP(value)
	case "base64":
		return validation.ValidInstance.Base64(value)
	case "mobile":
		return validation.ValidInstance.Mobile(value)
	case "tel":
		return validation.ValidInstance.Tel(value)
	case "mobileortel":
		return validation.ValidInstance.MobileOrTel(value)
	case "zipcode":
		return validation.ValidInstance.ZipCode(value)
	case "enum":
		return validation.ValidInstance.Enum(value, keySlice[1])
	}

	return nil
}
