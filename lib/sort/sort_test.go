package sort

import (
	"net/http"
	"testing"
)

func TestMapToSlice(t *testing.T) {
	params := map[string]string{
		"中文": "4",
		"b":  "2",
		"a":  "1",
		"c":  "3",
	}

	result := SortInstance.MapToSlice(params)
	if len(result) != 4 {
		t.Error("长度和输入不一致")
	}
	if result[0] != "1" || result[1] != "2" || result[2] != "3" || result[3] != "4" {
		t.Error("排序错误")
	}
}

func TestFormToSlice(t *testing.T) {
	req := &http.Request{}
	req.ParseForm()
	req.Form.Set("中文", "4")
	req.Form.Set("b", "2")
	req.Form.Set("a", "1")
	req.Form.Set("c", "3")
	req.Form.Set("sign", "s4se4rfse13f")

	result := SortInstance.FormToSlice(req, "sign")
	if len(result) != 4 {
		t.Error("长度和输入不一致")
	}
	if result[0] != "1" || result[1] != "2" || result[2] != "3" || result[3] != "4" {
		t.Error("排序错误")
	}
}
