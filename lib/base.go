package lib

import (
	"github.com/ascoders/as/lib/buffer"
	"github.com/ascoders/as/lib/captcha"
	"github.com/ascoders/as/lib/email"
	"github.com/ascoders/as/lib/http"
	"github.com/ascoders/as/lib/md5"
	"github.com/ascoders/as/lib/parse"
	"github.com/ascoders/as/lib/redis"
	"github.com/ascoders/as/lib/response"
	"github.com/ascoders/as/lib/router"
	"github.com/ascoders/as/lib/scheduled"
	"github.com/ascoders/as/lib/sort"
	"github.com/ascoders/as/lib/validation"
)

type LibStruct struct {
	Buffer    buffer.Buffer
	Captcha   captcha.Captcha
	Email     email.Email
	Http      http.Http
	Md5       md5.Md5
	Valid     validation.Valid
	Parse     parse.Parse
	Redis     redis.Redis
	Response  response.Response
	Router    router.Router
	Scheduled scheduled.Scheduled
	Sort      sort.Sort
}

var (
	Lib *LibStruct
)

func init() {
	Lib = &LibStruct{}
}
