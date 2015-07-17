package parse

import (
	"testing"
)

type TestData struct {
	Data1 string  `json:"data1" bson:"d1" valid:"required"`
	Data2 int     `json:"data2" bson:"d2" valid:"range(1,10)"`
	Data3 string  `json:"data3" bson:"d3" valid:"minlength(4);maxlength(5)"`
	Data4 float32 `json:"data4" bson:"d4" valid:"-"`
	Data5 string  `json:"data5" bson:"d5" valid:"-"`
	Data6 string  `json:"data6" bson:"d6" valid:"enum(a,b,c)"`
}

func TestStruct(t *testing.T) {
	data := &TestData{}

	params1 := map[string]interface{}{}
	if err := ParseInstance.Struct(data, params1); err == nil {
		t.Error("ParseInstance.Struct检测失败")
	}

	params2 := map[string]interface{}{
		"data1": "something",
		"data2": "0",
	}
	if err := ParseInstance.Struct(data, params2); err == nil {
		t.Error("ParseInstance.Struct检测失败")
	}

	params3 := map[string]interface{}{
		"data1": "something",
		"data2": "1",
	}
	if err := ParseInstance.Struct(data, params3); err != nil {
		t.Error("ParseInstance.Struct检测失败")
	}

	params4 := map[string]interface{}{
		"data1": "something",
		"data3": "长度适中",
	}
	if err := ParseInstance.Struct(data, params4); err != nil {
		t.Error("ParseInstance.Struct检测失败")
	}

	params5 := map[string]interface{}{
		"data1": "something",
		"data3": "长度太大了吧",
	}
	if err := ParseInstance.Struct(data, params5); err == nil {
		t.Error("ParseInstance.Struct检测失败")
	}

	params6 := map[string]interface{}{
		"data1": "something",
		"data4": "32.5",
	}
	if err := ParseInstance.Struct(data, params6); err == nil {
		t.Error("禁止修改字段阻止失败")
	}

	params7 := map[string]interface{}{
		"data1": "something",
		"data6": "a",
	}
	if err := ParseInstance.Struct(data, params7); err != nil {
		t.Error("enum失败")
	}

	params8 := map[string]interface{}{
		"data1": "something",
		"data6": "d",
	}
	if err := ParseInstance.Struct(data, params8); err == nil {
		t.Error("enum失败")
	}
}

func TestStructToUpdateMapp(t *testing.T) {
	data := &TestData{}

	params1 := map[string]interface{}{}
	if err, _ := ParseInstance.StructToUpdateMap(data, params1); err != nil {
		t.Error("require但参数未填应该不报错", err)
	}

	params2 := map[string]interface{}{
		"data1": "",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params2); err == nil {
		t.Error("require且参数已填应该报错")
	}

	params3 := map[string]interface{}{
		"data2": "0",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params3); err == nil {
		t.Error("ParseInstance.StructToUpdateMap检测失败")
	}

	params4 := map[string]interface{}{
		"data2": "1",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params4); err != nil {
		t.Error("ParseInstance.StructToUpdateMap检测失败")
	}

	params5 := map[string]interface{}{
		"data3": "长度适中",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params5); err != nil {
		t.Error("ParseInstance.StructToUpdateMap检测失败")
	}

	params6 := map[string]interface{}{
		"data3": "长度太大了吧",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params6); err == nil {
		t.Error("ParseInstance.StructToUpdateMap检测失败")
	}

	params7 := map[string]interface{}{
		"data4": "32.5",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params7); err == nil {
		t.Error("禁止修改字段阻止失败")
	}

	params8 := map[string]interface{}{
		"data4": "32.5",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params8, "data4"); err != nil {
		t.Error("禁止修改字段，加豁免却依然报错")
	}

	params9 := map[string]interface{}{
		"data4": "32.5",
		"data5": "sad",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params9, "data4", "data5"); err != nil {
		t.Error("豁免多个参数失效")
	}

	params10 := map[string]interface{}{
		"data6": "a",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params10); err != nil {
		t.Error("enum失败")
	}

	params11 := map[string]interface{}{
		"data6": "d",
	}
	if err, _ := ParseInstance.StructToUpdateMap(data, params11); err == nil {
		t.Error("enum失败")
	}
}
