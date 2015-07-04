# 配置

 通过设置`as.Conf.*`可以进行灵活配置，`*`可选列表如下：

- Debug `bool`

 是否开启调试模式：默认为`false`，在调试模式下会显示运行时间的log，并且自动路由会生成路由文件在目录`router/auto.go`下。

- Host `string`

 本机地址：默认为`127.0.0.1`。

- Port `string`

 运行端口：默认为`80`

- StaticDir `string`

 静态文件目录：默认为`static`

- StaticOptions `martini.StaticOptions`

 静态文件配置参数，详情参考martini文档

- GlobalPath `string`

 全局文件路径：默认为`空`，当没有路由规则匹配到时返回的内容，内容为该路径的文件内容。

[返回简介](/#配置)