/*==================================================
	Redis缓存接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package redis

import (
	"github.com/hoisie/redis"
	"newWoku/conf"
)

var (
	Client redis.Client
)

func init() {
	//初始化redis数据库
	Client.Addr = conf.REDIS_ADDRESS
}

// 设置缓存
func Set(key string, val []byte) {
	Client.Setex(key, conf.CACHE_EXPIRE, val)
}

// 设置缓存（包含时间）
func SetWithExpire(key string, val []byte, expire int64) {
	Client.Setex(key, expire, val)
}

// 读取缓存
func Get(key string) ([]byte, error) {
	return Client.Get(key)
}

// 删除缓存
func Delete(key string) {
	Client.Del(key)
}

// 根据前缀删除缓存
func DeletePrefix(prefix string) {
	keys, _ := Client.Keys(prefix + "*")

	for k, _ := range keys {
		Client.Del(keys[k])
	}
}
