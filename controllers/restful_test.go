package controllers

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"newWoku/models"
	"testing"
	"time"
)

// 测试表控制器
type RestfulTestController struct {
	Base
}

// 测试表数据
type RestfulTestData struct {
	Id       bson.ObjectId `bson:"_id" json:"id" valid:"-"`               // 主键
	Nickname string        `bson:"n" json:"nickname" valid:"required"`    // 昵称
	Password string        `bson:"p" json:"-"`                            // 密码
	Email    string        `bson:"e" json:"email" valid:"required;email"` // 电子邮箱
	Money    float32       `bson:"mo" json:"money" valid:"-"`             // 账户余额
	Free     int           `bson:"f" json:"free"`                         // 每月免费额度 !!!!!!!!mf
	LogCount uint16        `bson:"l" json:"logCount"`                     // 登陆次数
	LastTime time.Time     `bson:"la" json:"lastTime"`                    // 最后操作时间
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

	m := martini.Classic()
	m.Post("/api/restful_test/users", controller.Add)
	m.RunOnAddr(":80")

	r, _ := http.NewRequest("GET", "/api/restful_test/users", nil)
	w := httptest.NewRecorder()
	t.Error(r, w)

	// 删除数据库
	models.Db.C("restful_test_data").DropCollection()
}
