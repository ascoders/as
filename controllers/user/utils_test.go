package user

import (
	"net/http"
	"testing"
)

func TestCheckSign(t *testing.T) {
	// 生成sign
	sign, expire := CreateSign("testtoken", 30, map[string]string{
		"b": "b",
		"a": "a",
		"c": "c",
	})

	// 测试检测sign
	req := &http.Request{}
	req.ParseForm()
	req.Form.Set("c", "c")
	req.Form.Set("b", "b")
	req.Form.Set("a", "a")
	req.Form.Set("sign", sign)
	req.Form.Set("expire", expire)
	if err := CheckSign("testtoken", req); err != nil {
		t.Error("检测签名失败")
	}
}
