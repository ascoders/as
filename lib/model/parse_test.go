package model

import (
	"testing"
)

type TestData struct {
	Data1 string  `json:"data1" bson:"d1" valid:"required"`
	Data2 int     `json:"data2" bson:"d2" valid:"range(1,10)"`
	Data3 string  `json:"data3" bson:"d3" valid:"minlength(4);maxlength(5)"`
	Data4 float32 `json:"data4" bson:"d4" valid:"-"`
	Data5 string  `json:"data5" bson:"d5" valid:"-"`
}

func TestParse(t *testing.T) {
	data := &TestData{}

	params1 := map[string]string{}
	if err := Parse(data, params1); err == nil {
		t.Error("parse检测失败")
	}

	params2 := map[string]string{
		"data1": "something",
		"data2": "0",
	}
	if err := Parse(data, params2); err == nil {
		t.Error("parse检测失败")
	}

	params3 := map[string]string{
		"data1": "something",
		"data2": "1",
	}
	if err := Parse(data, params3); err != nil {
		t.Error("parse检测失败")
	}

	params4 := map[string]string{
		"data1": "something",
		"data3": "长度适中",
	}
	if err := Parse(data, params4); err != nil {
		t.Error("parse检测失败")
	}

	params5 := map[string]string{
		"data1": "something",
		"data3": "长度太大了吧",
	}
	if err := Parse(data, params5); err == nil {
		t.Error("parse检测失败")
	}

	params6 := map[string]string{
		"data1": "something",
		"data4": "32.5",
	}
	if err := Parse(data, params6); err == nil {
		t.Error("禁止修改字段阻止失败")
	}
}

func TestParseToUpdateMap(t *testing.T) {
	data := &TestData{}

	params1 := map[string]string{}
	if err, _ := ParseToUpdateMap(data, params1); err != nil {
		t.Error("require但参数未填应该不报错", err)
	}

	params2 := map[string]string{
		"data1": "",
	}
	if err, _ := ParseToUpdateMap(data, params2); err == nil {
		t.Error("require且参数已填应该报错")
	}

	params3 := map[string]string{
		"data2": "0",
	}
	if err, _ := ParseToUpdateMap(data, params3); err == nil {
		t.Error("ParseToUpdateMap检测失败")
	}

	params4 := map[string]string{
		"data2": "1",
	}
	if err, _ := ParseToUpdateMap(data, params4); err != nil {
		t.Error("ParseToUpdateMap检测失败")
	}

	params5 := map[string]string{
		"data3": "长度适中",
	}
	if err, _ := ParseToUpdateMap(data, params5); err != nil {
		t.Error("ParseToUpdateMap检测失败")
	}

	params6 := map[string]string{
		"data3": "长度太大了吧",
	}
	if err, _ := ParseToUpdateMap(data, params6); err == nil {
		t.Error("ParseToUpdateMap检测失败")
	}

	params7 := map[string]string{
		"data4": "32.5",
	}
	if err, _ := ParseToUpdateMap(data, params7); err == nil {
		t.Error("禁止修改字段阻止失败")
	}

	params8 := map[string]string{
		"data4": "32.5",
	}
	if err, _ := ParseToUpdateMap(data, params8, "data4"); err != nil {
		t.Error("禁止修改字段，加豁免却依然报错")
	}

	params9 := map[string]string{
		"data4": "32.5",
		"data5": "sad",
	}
	if err, _ := ParseToUpdateMap(data, params9, "data4", "data5"); err != nil {
		t.Error("豁免多个参数失效")
	}
}
