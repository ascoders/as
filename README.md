# as
martini 轻巧脚手架

# 依赖

~~~ go
go get -u github.com/go-martini/martini
go get -u github.com/martini-contrib/encoder
go get -u github.com/hoisie/redis
go get -u github.com/martini-contrib/sessions
go get -u github.com/ascoders/xsrftoken
go get -u gopkg.in/mgo.v2
go get -u github.com/dchest/captcha
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

# 注释路由

在控制器的每个函数上，按照规范添加注释，可以自动生成注释路由文件。

最基础的写法，默认响应`get`请求：
~~~go
// @router /example
~~~

url可以带参数：
~~~go
// @router /example/:id
~~~

设置响应类型，多个用`,`分隔：
~~~go
// @router /example [post]
~~~

设置响应前执行的函数，下例先执行`Csrf`方法，例如做权限控制，多个用`,`分隔：
~~~go
// @router /example (csrf)
~~~

也可以混合使用
~~~go
// @router /example (csrf,before) [put,delete]
~~~

如果开启了自动`restful路由`，可以复写注释路由替换默认的`restful路由`，替换`Gets` `Get` `Add` `Update` `Delete`方法，使用`this.Restful.[RestfulApi]`：
~~~go
// @router /users (before) [get]
func (this *User) Gets(req *http.Request) []byte {
	return this.Restful.Gets(req)
}
~~~

注释路由的参数**对大小写不敏感**

# 控制器规范

~~~
controllers
|-- user
|	|-- user.go
|	|-- extend.go
|-- article
|	|-- article.go
~~~

每个控制器主文件必写`New`方法，供初始化路由时使用，如：
~~~go
func New() *User {
	controller := &User{}
	controller.NewModel(models.NewUser()) // 初始化数据模型
	return controller
}
~~~

每个`package`对应一个资源，并自动开启`restful Api`，需要权限验证或禁用某些api，可以复写`restful`方法。以`user`为例，自动生成的`restful api`如下：

~~~ go
user := &user.User{}
r.Get	("/api/users", 		user.Before, user.Gets)
r.Get	("/api/users/:id", 	user.Before, user.Get)
r.Post	("/api/users", 		user.Before, user.Add)
r.Patch	("/api/users", 		user.Before, user.Update)
r.Delete("/api/users/:id", 	user.Before, user.Delete)
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

# Csrf过滤

如果将 `conf.CSRF_DEFAULT` 设置为 `true`，系统会自动在所有非 `get` 方法添加 `csrf` 验证。也可以在需要的`get`方法中，通过注释路由手动添加，使用 `(crsf)` 关键字，如：
~~~go
@router /example (csrf)
~~~

# 验证码

书写注释路由时，使用 `(captcha)` 关键字，可以开启验证码校验功能：

~~~go
// @router /example (csrf,captcha)
~~~

# 自动验证

使用`this.Parse`解析表单时，如果数据`struct`按照规范书写了验证`tag`，将会自动按照规范进行验证。

`tag`属性：`valid`：

`Required`				字符串不为空
`Min(min int)` 			最小值
`Max(max int)` 			最大值
`Range(min, max int)` 	数值的范围
`MinSize(min int)` 		最小长度
`MaxSize(max int)` 		最大长度
`Length(length int)`	长度
`AlphaNumeric` 			alpha字符或数字
`Email` 				邮箱格式
`IP` 					IP格式
`Base64` 				base64编码
`Mobile` 				手机号
`Tel` 					固定电话号
`Phone` 				手机号或固定电话号
`ZipCode` 				邮政编码
`Disabled`				不允许用户对此字段的修改

> 特别说明：(POST)新增数据时会检查结构体中所有被标识为`required`的字段，而局部更新时(PATCH)只会对提交的属性进行验证。

# 其它`tag`功能说明

`tag`属性：`bson`：

`someValue` 写入数据库时的`key`
`-`			这个字段不会写入数据库

`tag`属性：`json`：

`someValue`	与客户端交互时`json`的`key`
`-`			这个字段不会发送到客户端