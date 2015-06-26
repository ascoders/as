package as

import (
	"github.com/ascoders/as/lib"
)

var (
	Lib *lib.LibInstance
)

func init() {
	Lib = lib.Lib
}
