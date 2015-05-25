/*==================================================
	配置文件

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package conf

const (
	DEBUG bool = true

	// 本机地址
	HOST string = "127.0.0.1"

	// 运行端口
	PORT uint8 = 80

	// 静态文件目录
	STATIC_DIR string = "static"

	// 全局文件路径
	GLOBAL_PATH string = "static/html/public/global.html"

	// redis地址
	REDIS_ADDRESS string = "127.0.0.1:6379"

	// mongodb地址
	MONGODB_ADDRESS string = "127.0.0.1:11578"

	// 缓存时间
	CACHE_EXPIRE int64 = 60 * 60

	// 是否开启默认csrf
	AUTO_SCRF = true

	// 邮箱
	EMAIL_HOST     = SECERT_EMAIL_HOST
	EMAIL_FROM     = SECERT_EMAIL_FROM
	EMAIL_PORT     = SECERT_EMAIL_PORT
	EMAIL_PASSWORD = SECERT_EMAIL_PASSWORD

	// 错误类型
	ERROR_TYPE string = "类型错误"
)
