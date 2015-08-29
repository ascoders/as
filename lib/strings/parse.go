package strings

import (
	"strconv"
)

type Strings struct{}

var (
	StringsInstance *Strings
)

func init() {
	StringsInstance = &Strings{}
}

// 转化为int
func (this *Strings) ParseInt(number interface{}) int {
	result := 0
	switch number.(type) {
	case int: // good!
		result = number.(int)
	case string:
		result, _ = strconv.Atoi(number.(string))
	}

	return result
}
