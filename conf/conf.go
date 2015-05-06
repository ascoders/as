/*==================================================
	配置文件

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package conf

const (
	// 本机地址
	HOST string = "127.0.0.1"

	// 运行端口
	PORT uint8 = 80

	// 静态文件目录
	STATIC_DIR string = "static"

	// 全局文件路径
	GLOBAL_PATH string = "static/static/html/public/global.html"

	// redis地址
	REDIS_ADDRESS string = "127.0.0.1:6379"

	// 缓存时间
	CACHE_EXPIRE int64 = 60 * 60
)
