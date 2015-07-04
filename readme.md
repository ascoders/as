# as
基于 martini 的轻巧框架

# 使用

如何开启一个监听80端口的服务：

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

可以通过`as.Conf.*`配置任意参数，请注意将它放在`as.Run()`之前。同时你也可以单独写在一个配置文件中，查看[配置列表](conf)。

~~~go
func main() {
	as.Conf.Debug = true
	as.Conf.Host = "59.67.115.1"
	as.Conf.Port = 8080
	
	as.Run()
}
~~~

# 控制器

控制器内嵌了`as.Controller`，这样变包含了所有`as.Controller`拥有的方法，并自动实现了`Restful`方法，查看[控制器设计](controllers)。

同时定义`New`方法，由于`as`是一个`restful`框架，一个url、一个控制器、一个模型同时对应一个资源，所以将对应模型对象通过`NewModel`方法注册到其中，供路由使用。

~~~go
package app

import (
	"github.com/ascoders/as"
	"product/models/app"
	
)

type Controller struct {
	as.Controller
}

func New() *Controller {
	controllerInstance := &Controller{}
	controllerInstance.NewModel(app.ModelInstance)
	return controllerInstance
}
~~~

# 模型

# 路由

as框架使用了注释路由，如下写法可以注册一个`/apps`url下的路由，响应`get`方法的请求，并返回ok，响应号为200.

~~~go
// @router /apps [get]
func (this *Controller) Gets() (int, []byte) {
	return this.Success("ok")
}
~~~