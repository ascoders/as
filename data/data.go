package data

import ()

type Data struct {
	// 主键
	Id int `json:"id" sql:"type:int unsigned AUTO_INCREMENT" gorm:"primary_key"`
}
