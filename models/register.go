package models

import (
	"fmt"
	"github.com/ascoders/as/db"
)

var (
	baseModels []*Base
)

// 为model注册数据orm
// @param obj 资源
func (this *Base) Register(obj interface{}) {
	this.Data = obj
	baseModels = append(baseModels, this)
}

// 注册所有资源（由run调用，此时可以获取用户自定义配置）
func RegisterAll() {
	for _, v := range baseModels {
		// 自动迁移
		if err := db.Orm.AutoMigrate(v.Data).Error; err != nil {
			fmt.Println("建表失败：", err)
			return
		}

		// 设置db
		v.Db = db.Orm.Model(v.Data)
	}

	// 释放资源
	baseModels = nil
}
