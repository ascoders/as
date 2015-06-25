package controllers

import (
	"github.com/ascoders/as/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"testing"
)

// 测试表控制器
type RestfulTestController struct {
	Base
}

// 测试表数据
type RestfulTestData struct {
	Id       bson.ObjectId `bson:"_id" json:"id" valid:"-"`               // 主键
	Nickname string        `bson:"n" json:"nickname" valid:"required"`    // 昵称
	Email    string        `bson:"e" json:"email" valid:"required;email"` // 电子邮箱
	Money    float32       `bson:"mo" json:"money" valid:"-"`             // 账户余额
}

// 测试表操作模型
type RestfulTestModel struct {
	models.Base
}

func New() *RestfulTestModel {
	model := &RestfulTestModel{}
	model.Collection = models.Db.C("restful_test_data")

	if err := model.Collection.EnsureIndex(mgo.Index{
		Key:    []string{"e"},
		Unique: true,
	}); err != nil {
		panic(err)
	}
	return model
}
func (this *RestfulTestModel) NewData() interface{} {
	var r RestfulTestData
	return &r
}
func (this *RestfulTestModel) NewDataWithId() interface{} {
	var r RestfulTestData
	r.Id = bson.NewObjectId()
	return &r
}
func (this *RestfulTestModel) NewDatas() interface{} {
	var r []*RestfulTestData
	return &r
}

/*
r.Get("/api/users", user.Gets)
r.Get("/api/users/:id", user.Get)
r.Post("/api/users", user.Add)
r.Patch("/api/users/:id", user.Update)
r.Delete("/api/users/:id", user.Delete)
*/

func TestAdd(t *testing.T) {
	controller := &RestfulTestController{}
	Model := New()
	controller.NewModel(Model)

	req1 := &http.Request{}
	req1.ParseForm()
	req1.Form.Set("nickname", "testName")
	if status, message := controller.Add(req1); status != 400 {
		t.Error("add方法缺少required参数未提示", string(message))
	}

	req2 := &http.Request{}
	req2.ParseForm()
	req2.Form.Set("nickname", "testName")
	req2.Form.Set("email", "test")
	if status, message := controller.Add(req2); status != 400 {
		t.Error("email格式错误未报错", string(message))
	}

	req3 := &http.Request{}
	req3.ParseForm()
	req3.Form.Set("nickname", "testName")
	req3.Form.Set("email", "123@qq.com")
	if status, message := controller.Add(req3); status != 200 {
		t.Error("add方法即使required参数全也报错", string(message))
	}

	// 删除数据库
	models.Db.C("restful_test_data").DropCollection()
}

func TestGets(t *testing.T) {
	controller := &RestfulTestController{}
	Model := New()
	controller.NewModel(Model)

	for v := 0; v <= 10; v++ {
		req := &http.Request{}
		req.ParseForm()
		req.Form.Set("nickname", "testName")
		req.Form.Set("email", "123@qq.com"+strconv.Itoa(v))
		if status, message := controller.Add(req); status != 200 {
			t.Error("add方法错误", string(message))
		}
	}

	req2 := &http.Request{}
	req2.ParseForm()
	if status, message := controller.Gets(req2); status != 200 {
		t.Error("gets方法错误", string(message))
	}

	// 删除数据库
	models.Db.C("restful_test_data").DropCollection()
}
