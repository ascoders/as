package md5

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5 struct{}

var (
	Md5Instance *Md5
)

func init() {
	Md5Instance = &Md5{}
}

func (this *Md5) Md5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}
