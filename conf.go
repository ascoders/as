package as

import (
	"github.com/ascoders/as/conf"
)

var (
	Conf *conf.Config // 基础配置
)

func init() {
	// as.Conf == conf.Conf
	// conf.Conf用于内部调用（引用as.conf会出现循环引用）
	Conf = conf.Conf
	Conf.MongodbAddress = "xxx"
}
