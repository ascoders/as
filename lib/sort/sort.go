package sort

import (
	"net/http"
	"sort"
)

type Sort struct{}

// 将map按照key，字母大小顺序排序
// @return []string
func (this *Sort) MapToSlice(params map[string]string) []string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var slice []string
	for k, _ := range keys {
		slice = append(slice, params[keys[k]])
	}

	return slice
}

// 将http.request按照key，字母大小顺序排序
// @return []string
func (this *Sort) FormToSlice(req *http.Request, ignore ...string) []string {
	req.ParseForm()
	params := make(map[string]string)
FormLoop:
	for k, _ := range req.Form {
		for q, _ := range ignore {
			if k == ignore[q] {
				continue FormLoop
			}
		}
		params[k] = req.Form.Get(k)
	}

	return this.MapToSlice(params)
}
