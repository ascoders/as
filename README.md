# as
martini 轻巧脚手架

# 依赖

~~~ go
go get github.com/go-martini/martini
go get github.com/martini-contrib/encoder
go get github.com/hoisie/redis
go get github.com/martini-contrib/sessions
go get gopkg.in/mgo.v2
~~~

# 自动路由

自动路由最强悍的地方就是每次运行程序，都会根据`controllers`中的注释生成路由文件`router/auto.go`。为了享受自动生成路由带来的便利，你可以如下编写注释：

~~~ go
type Article struct {
	controllers.Base
}

// @router /article [Get]
func (this *Article) Other() []byte {
	return response.Success("Delete success!")
}

~~~

> 第一次执行程序会生成自动路由文件，但无法调用它，第二次启动就能自动加载第一次生成的路由文件了。

**控制器规范**

~~~
controllers
	|-- user
	|	|-- user.go
	|	|-- extend.go
	|-- article
	|	|-- article.go
~~~

每个`package`对应一个资源，并自动开启`restful Api`，需要权限验证或禁用某些api，可以复写`restful`方法。以`user`为例，自动生成的`restful api`如下：

~~~ go
user := &user.User{}
r.Get("/users", user.Before ,user.Gets)
r.Get("/users/:id", user.Before ,user.Get)
r.Post("/users", user.Before ,user.Add)
r.Put("/users", user.Before ,user.Update)
r.Delete("/users/:id", user.Before ,user.Delete)
~~~

覆写`Before`方法，可以做权限验证等处理：

~~~ go
func (this *User) Before(w http.ResponseWriter) {
	// 只要有输出语句，后面的路由逻辑不会执行
	w.Write([]byte("没有权限"))
}
~~~

# 自动缓存

脚手架使用`martini`映射接口的特性，覆盖了`http.ResponseWriter`并重写`write()`方法，在其调用前生成以当前`url`作为`key`，当前输出内容为`value`的缓存，并在http请求发生前优先使用缓存。如果路由遵循`restful`规范，只有`get`请求会使用缓存（因为其他操作数据可能发生了变化），这一切都是自动的。

需要注意如下两点：

1. 手动清除缓存：引入`lib/redis`，调用`redis.Delete`方法。
2. 在配置文件`conf/conf.go`中修改默认缓存周期`CACHE_EXPIRE`。