# 模型

数据库采用mongodb，使用`Registe`方法在模型中注册一个数据表，第一个参数为表名，第二个参数为配置，详见mgo配置。

~~~go
var (
	ModelInstance *Model
)

func init() {
	ModelInstance = &Model{}
	ModelInstance.Registe("apps", mgo.Index{
		Key:    []string{"e"},
		Unique: true,
	})
}
~~~

如果实现了`NewData` `NewDataWithId` `NewDatas`这三个方法，模型就自动实现了restful增删改查的全部功能。

~~~go
func (this *Model) NewData() interface{} {
	var r Data
	return &r
}

func (this *Model) NewDataWithId() interface{} {
	var r Data
	r.Id = bson.NewObjectId()
	return &r
}

func (this *Model) NewDatas() interface{} {
	var r []*Data
	return &r
}
~~~