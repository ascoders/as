package lib

import (
	"github.com/ascoders/as/lib/buffer"
	"github.com/ascoders/as/lib/captcha"
	"github.com/ascoders/as/lib/csrf"
	"github.com/ascoders/as/lib/http"
	"github.com/ascoders/as/lib/parse"
	"github.com/ascoders/as/lib/response"
	"github.com/ascoders/as/lib/scheduled"
	"github.com/ascoders/as/lib/sort"
	"github.com/ascoders/as/lib/validation"
)

type Lib struct {
	Buffer    buffer.Buffer
	Captcha   captcha.Captcha
	Http      http.Http
	Valid     validation.Valid
	Parse     parse.Parse
	Response  response.Response
	Scheduled scheduled.Scheduled
	Sort      sort.Sort
	Csrf      csrf.Csrf
}

var (
	LibInstance *Lib
)

func init() {
	LibInstance = &Lib{}
}
