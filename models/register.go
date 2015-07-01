// 以下操作时数据库连接未建立

package models

import (
	"gopkg.in/mgo.v2"
)

var (
	BaseLists []*Base // 用户创建的数据表，在数据库连接时会被统一初始化
)

// 初始化数据表
func (this *Base) Registe(table string, index ...mgo.Index) {
	this.Table = table
	this.Indexs = index
	BaseLists = append(BaseLists, this)
}
