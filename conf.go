package as

import (
	"github.com/ascoders/as/lib/csrf"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

var (
	Conf *Config // 基础配置
)

type Config struct {
	Debug          bool                  // 是否开启调试模式
	Host           string                // 本机地址
	Port           uint8                 // 运行端口
	StaticDir      string                // 静态文件目录
	StaticOptions  martini.StaticOptions // 静态文件配置参数
	GlobalPath     string                // 全局文件路径
	RedisAddress   string                // redis地址
	MongodbAddress string                // mongodb地址
	CacheExpire    int                   // 缓存时间
	CsrfAuto       bool                  // 是否开启默认csrf
	CsrfOptions    *csrf.Options         // scrf加盐
	SessionName    string                // session名称
	SessionSecret  string                // session密钥
	SessionExpire  int                   // session过期时间
	SessionUserKey string                // session用户id的键名
	SessionOptions sessions.Options      // session参数
	EmailHost      string                // 邮箱地址
	EmailFrom      string                // 邮箱发送方
	EmailPort      int                   // 邮箱端口
	EmailPassword  string                // 邮箱密码
	ErrorType      string                // 类型错误提示
}

func init() {
	Conf = &Config{}

	// 设置默认值
	Conf.Debug = true
	Conf.Host = "127.0.0.1"
	Conf.Port = 80

	Conf.StaticDir = "static"
	Conf.StaticOptions = martini.StaticOptions{
		Prefix:      "/static/",
		SkipLogging: true,
	}

	Conf.GlobalPath = "static/public/global.html"
	Conf.RedisAddress = "127.0.0.1:6379"
	Conf.MongodbAddress = "127.0.0.1:11578"
	Conf.CacheExpire = 60 * 60

	Conf.CsrfAuto = true
	Conf.CsrfOptions = &csrf.Options{
		Secret:     "507dQJpITMRungvQ5kh2fiGVRWLqFg",
		SessionKey: "id", // 根据用户id，为每个用户设置不同的csrf
		SetCookie:  true,
	}

	Conf.SessionName = "asSession"
	Conf.SessionSecret = "vpahHL29ajXuTY0RNhf1VYTHvRIJxX"
	Conf.SessionExpire = 60 * 60 * 24 * 14
	Conf.SessionOptions = sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   Conf.SessionExpire,
		Secure:   false,
		HttpOnly: true,
	}

	Conf.ErrorType = "类型错误"
}
