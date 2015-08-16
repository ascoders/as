package conf

import (
	"github.com/ascoders/as/lib/csrf"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

var (
	ConfInstance *Conf // 基础配置
)

type Conf struct {
	Debug          bool                  // 是否开启调试模式
	Host           string                // 本机地址
	Port           uint8                 // 运行端口
	StaticDir      string                // 静态文件目录
	StaticOptions  martini.StaticOptions // 静态文件配置参数
	GlobalPath     string                // 全局文件路径
	RedisAddress   string                // redis地址
	DBMaxIdleConns int                   // 数据库最大空闲连接数
	DBMaxOpenConns int                   // 数据库最大在线连接数
	DbAddress      string                // 数据库地址
	DbUserName     string                // 数据库用户名
	DbPassword     string                // 数据库密码
	DbName         string                // 数据库名称
	CacheExpire    int64                 // 缓存时间
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
	CaptchaUrl     string                // 验证码地址
	CaptchaName    string                // 验证码内容校验key
	CaptchaIdName  string                // 验证码id校验key
	ErrorType      string                // 类型错误提示
}

func init() {
	ConfInstance = &Conf{}

	// 设置默认值
	ConfInstance.Debug = true
	ConfInstance.Host = "127.0.0.1"
	ConfInstance.Port = 80

	ConfInstance.StaticDir = "static"
	ConfInstance.StaticOptions = martini.StaticOptions{
		Prefix:      "/static/",
		SkipLogging: true,
	}

	ConfInstance.GlobalPath = ""
	ConfInstance.RedisAddress = "127.0.0.1:6379"
	ConfInstance.DbAddress = "127.0.0.1:27017"
	ConfInstance.DbUserName = "root"
	ConfInstance.CacheExpire = 60 * 60

	ConfInstance.CsrfAuto = true
	ConfInstance.CsrfOptions = &csrf.Options{
		Secret:     "507dQJpITMRungvQ5kh2fiGVRWLqFg",
		SessionKey: "id", // 根据用户id，为每个用户设置不同的csrf
		SetCookie:  true,
	}

	ConfInstance.DBMaxIdleConns = 10
	ConfInstance.DBMaxOpenConns = 100

	ConfInstance.SessionName = "asSession"
	ConfInstance.SessionSecret = "vpahHL29ajXuTY0RNhf1VYTHvRIJxX"
	ConfInstance.SessionExpire = 60 * 60 * 24 * 14
	ConfInstance.SessionOptions = sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   ConfInstance.SessionExpire,
		Secure:   false,
		HttpOnly: true,
	}

	ConfInstance.CaptchaName = "captcha"
	ConfInstance.CaptchaIdName = "capid"

	ConfInstance.ErrorType = "类型错误"
}
