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

- `Required`				字符串不为空
- `Min(min int)` 			最小值
- `Max(max int)` 			最大值
- `Range(min, max int)` 	数值的范围
- `MinSize(min int)` 		最小长度
- `MaxSize(max int)` 		最大长度
- `Length(length int)`		长度
- `AlphaNumeric` 			alpha字符或数字
- `Email` 					邮箱格式
- `IP` 						IP格式
- `Base64` 					base64编码
- `Mobile` 					手机号
- `Tel` 					固定电话号
- `Phone` 					手机号或固定电话号
- `ZipCode` 				邮政编码
- `-`						解析请求时会忽略此字段（用户无法通过提交参数修改此字段）

> 特别说明：(POST)新增数据时会检查结构体中所有被标识为`required`的字段，而局部更新时(PATCH)只会对提交的属性进行验证。

# 其它tag功能说明

`tag`属性：`bson`：

- `someValue` 写入数据库时的`key`
- `-`			这个字段不会写入数据库

`tag`属性：`json`：

- `someValue`	与客户端交互时`json`的`key`
- `-`			这个字段不会发送到客户端

# 增改最佳实践

**新增一条记录**

用户需要提交所有记录应该包含的字段，**服务器*需要验证*全部`required`的字段**以避免非空字段出现空值。同时通过设置`valid:"-"`对不应由用户设置的字段进行屏蔽（例如账户余额），而在业务逻辑中添加：

~~~go
func (this *Controller) Add(req *http.Request) (int, []byte) {
	// 实例化用户表结构体
	// @return type struct
	user := &user.Data{}
	
	// 将req.form转换为map
	// @return map[string]string
	params := this.ReqFormToMap(req)
	
	// 根据用户表结构体的设置，将params参数解析到结构体中，参数类型错误返回error
	// 同时进行自动验证，验证失败返回error
	// 如果提交参数包含valid:"-"的，会返回禁止修改的error
	if err := this.Parse(user, params); err != nil {
		return this.Error(err.Error())
	}
	
	// 手动设置敏感字段，例如给新注册用户送10元现金
	// user.Money = 10
	
	// 添加用户
	if err := this.Model.Add(data); err != nil {
		return response.Error(err.Error())
	}
}
~~~

**更新一条记录**

用户需要提交条目的`id`以及需要修改的所有字段，**此时*不会验证*全部`required`的字段**，而是仅仅对提交的字段做验证（因为未改动的字段符合过滤规范），同时也适用`valid:"-"`规则：

~~~go
func (this *Controller) Add(req *http.Request) (int, []byte) {
	// 实例化用户表结构体
	// @return type struct
	user := &user.Data{}
	
	// 将req.form转换为map
	// @return map[string]string
	params := this.ReqFormToMap(req)
	
	// 将用户提交参数转换为可以执行update命令的参数
	// 例如nickname字段设置了bson:"n"，会将提交的nickname参数修改为n
	// 同时进行自动验证，验证失败返回error
	// 如果提交参数包含valid:"-"的，会返回禁止修改的error
	// @return map[string]string
	if err, opts := this.ParseToUpdateMap(data, params); err == nil {
		// 手动设置敏感字段，例如刷新用户操作时间
		// user.LastOperate = time.Now()
		
		// 更新用户
		err := this.Model.Update(param["id"], opts)
		return response.Must("更新成功", err)
	} else {
		return response.Error(err.Error())
	}
}
~~~

有时会进行特殊更新，即只更新指定的字段（用户无法通过指定任意字段来更新任意字段，或某个字段的更新要经过特殊的过滤逻辑）例如修改密码。为了避免任意更新，这些字段一般会被设置`valid:"-"`规则，无法通过`this.ParseToUpdateMap`的验证，此时可以通过传入额外参数豁免`valid:"-"`的规则：

~~~go
// 用户通过邮箱验证，才有修改密码的权限
// if bla..bla..bla
// return "token验证失败"

// 为password和nickname字段豁免valid:"-"的验证
// 用户可能通过了更高级别的验证，而获取了对这些“敏感”字段的修改权
// 但同样会根据valid的设置验证字段格式
if err, opts := this.ParseToUpdateMap(data, params, "password", "nickname"); err == nil {
	// success!
}
~~~

自动生成的`restful api`符合此规则。
