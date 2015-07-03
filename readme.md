# as
基于 martini 的轻巧框架

# 使用

如果想创建一个监听80端口的服务，只需要一行代码：

~~~go
package main

import (
	"github.com/ascoders/as"
)

func main() {
	as.Run()
}
~~~

# 配置

可以通过`as.Conf.*`配置任意参数，请注意将它放在`as.Run()`之前。同时你也可以单独写在一个配置文件中，查看<a href='/ascoders/as/conf/readme.md'>配置列表</a>。

~~~go
func main() {
	as.Conf.Debug = true
	as.Conf.Host = "59.67.115.1"
	as.Conf.Port = 8080
	
	as.Run()
}
~~~