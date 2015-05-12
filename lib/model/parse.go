/*==================================================
	模型 解析

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package model

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 解析url参数
// @param {interface{}} obj 被解析的结构体
// @param {http.Request} req 客户端请求
func Parse(obj interface{}, req *http.Request) error {
	objT := reflect.TypeOf(obj).Elem()
	objV := reflect.ValueOf(obj).Elem()

	// 解析表单
	req.ParseForm()

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

		// 跳过url不存在的参数
		value := req.PostForm.Get(tag)
		if len(value) == 0 {
			continue
		}

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

// 根据struct和req.form解析
// @param {interface{}} obj 被解析的结构体
// @param {http.Request} req 客户端请求
// @return map[string]interface{}
func ParseTo(obj interface{}, req *http.Request) map[string]interface{} {
	opts := make(map[string]interface{})

	objT := reflect.TypeOf(obj).Elem()
	objV := reflect.ValueOf(obj).Elem()

	// 解析表单
	req.ParseForm()

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
		value := req.PostForm.Get(tag)
		if len(value) == 0 {
			continue
		}

		switch fieldT.Type.Kind() {
		case reflect.Bool:
			if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
				opts[bson] = true
				continue
			}
			if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
				opts[bson] = false
				continue
			}
			b, err := strconv.ParseBool(value)
			if err != nil {
				return nil
			}
			opts[bson] = b
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil
			}
			opts[bson] = x
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return nil
			}
			opts[bson] = x
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil
			}
			opts[bson] = x
		case reflect.Interface:
			opts[bson] = reflect.ValueOf(value)
		case reflect.String:
			opts[bson] = value
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
					return nil
				}
				opts[bson] = reflect.ValueOf(t)
			}
		}
	}

	return opts
}
