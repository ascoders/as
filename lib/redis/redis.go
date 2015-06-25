/*==================================================
	Redis缓存接口

	Copyright (c) 2015 翱翔大空 and other contributors
 ==================================================*/

package redis

import (
	"github.com/ascoders/as/conf"
	"github.com/hoisie/redis"
)

type Redis struct{}

var (
	client redis.Client
)

func init() {
	//初始化redis数据库
	client.Addr = conf.Conf.RedisAddress
}

// 设置缓存
func (this *Redis) Set(key string, val []byte) {
	client.Setex(key, conf.Conf.CacheExpire, val)
}

// 设置缓存（包含时间）
func (this *Redis) SetWithExpire(key string, val []byte, expire int64) {
	client.Setex(key, expire, val)
}

// 读取缓存
func (this *Redis) Get(key string) ([]byte, error) {
	return client.Get(key)
}

// 删除缓存
func (this *Redis) Delete(key string) {
	client.Del(key)
}

// 根据前缀删除缓存
func (this *Redis) DeletePrefix(prefix string) {
	keys, _ := client.Keys(prefix + "*")

	for k, _ := range keys {
		client.Del(keys[k])
	}
}
