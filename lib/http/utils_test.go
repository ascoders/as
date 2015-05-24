package http

import (
	"net/http"
	"testing"
)

func TestReqFormToMap(t *testing.T) {
	req1 := &http.Request{}
	req1.ParseForm()
	req1.Form.Set("nickname", "testName")
	params1 := ReqFormToMap(req1)
	if len(params1) != 1 || params1["nickname"] != "testName" {
		t.Error("ReqFormToMap错误")
	}

	req2 := &http.Request{}
	req2.ParseForm()
	req2.Form.Set("nickname", "testName")
	req2.Form.Set("other", "other")
	params2 := ReqFormToMap(req2, "other")
	if len(params2) != 1 || params2["nickname"] != "" || params2["other"] != "other" {
		t.Error("ReqFormToMap错误")
	}

	req3 := &http.Request{}
	req3.ParseForm()
	req3.Form.Set("nickname", "testName")
	req3.Form.Set("other", "other")
	req3.Form.Set("ttt", "sss")
	params3 := ReqFormToMap(req3, "other", "ttt")
	if len(params3) != 2 || params3["nickname"] != "" || params3["other"] != "other" || params3["ttt"] != "sss" {
		t.Error("ReqFormToMap错误")
	}
}
