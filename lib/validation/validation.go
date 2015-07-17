/*==================================================
	验证

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package validation

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	emailPattern   = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	ipPattern      = regexp.MustCompile("^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$")
	base64Pattern  = regexp.MustCompile("^(?:[A-Za-z0-99+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")
	mobilePattern  = regexp.MustCompile("^((\\+86)|(86))?(1(([35][0-9])|(47)|[8][012356789]))\\d{8}$")
	telPattern     = regexp.MustCompile("^(0\\d{2,3}(\\-)?)?\\d{7,8}$")
	zipCodePattern = regexp.MustCompile("^[1-9]\\d{5}$")
)

type Valid struct{}

var (
	ValidInstance *Valid
)

func init() {
	ValidInstance = &Valid{}
}

// 字符串不能为空
func (this *Valid) Required(value string) error {
	if value != "" {
		return nil
	}
	return errors.New("不能为空")
}

// 最小值
func (this *Valid) Min(value string, min int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num >= min {
			return nil
		}
	}
	return errors.New("最小为" + strconv.Itoa(min))
}

// 最大值
func (this *Valid) Max(value string, max int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num <= max {
			return nil
		}
	}
	return errors.New("最大为" + strconv.Itoa(max))
}

// 数值范围
func (this *Valid) Range(value string, min int, max int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num >= min && num <= max {
			return nil
		}
	}
	return errors.New("范围为" + strconv.Itoa(min) + "-" + strconv.Itoa(max))
}

// 最小长度
func (this *Valid) MinLength(value string, min int) error {
	if len([]rune(value)) >= min {
		return nil
	}
	return errors.New("最小长度为" + strconv.Itoa(min))
}

// 最大长度
func (this *Valid) MaxLength(value string, max int) error {
	if len([]rune(value)) <= max {
		return nil
	}
	return errors.New("最大长度为" + strconv.Itoa(max))
}

// 长度范围
func (this *Valid) Length(value string, length int) error {
	if len([]rune(value)) == length {
		return nil
	}
	return errors.New("长度应为" + strconv.Itoa(length))
}

// 字母或数字的组合
func (this *Valid) AlphaNumeric(value string) error {
	for _, v := range value {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			return errors.New("只能为字母或数字")
		}
	}
	return nil
}

// 电子邮箱
func (this *Valid) Email(value string) error {
	if emailPattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为电子邮箱")
}

// IP地址
func (this *Valid) IP(value string) error {
	if ipPattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为IP地址")
}

// Base64编码
func (this *Valid) Base64(value string) error {
	if base64Pattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为Base64编码")
}

// 手机号
func (this *Valid) Mobile(value string) error {
	if mobilePattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为手机号")
}

// 固话
func (this *Valid) Tel(value string) error {
	if telPattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为固话")
}

// 手机号或固话
func (this *Valid) MobileOrTel(value string) error {
	if mobilePattern.MatchString(value) || telPattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为手机号或固话")
}

// 邮政编码
func (this *Valid) ZipCode(value string) error {
	if zipCodePattern.MatchString(value) {
		return nil
	}
	return errors.New("格式为邮政编码")
}

// 指定字符串
func (this *Valid) Enum(value string, enums string) error {
	// 拆分字符串
	each := strings.Split(enums, ",")
	for k, _ := range each {
		if value == strings.TrimSpace(each[k]) {
			return nil
		}
	}
	return errors.New("不在指定范围内")
}
