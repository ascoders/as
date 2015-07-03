package router

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	workPath, _          = os.Getwd()
	globalRouterTemplate = `package router

import (
	"github.com/ascoders/as"
	{{.packageInfo}}
)

func init() {
    {{.globalInfo}}
}
`
)

// store the comment for the controller method
type ControllerComments struct {
	Method           string
	Router           string
	AllowHTTPMethods []string
	PrefixMethods    []string
	Params           []map[string]string
}

var (
	commentFilename string
	pkgLastupdate   map[string]int64
	genInfoList     map[string][]ControllerComments
)

func init() {
	pkgLastupdate = make(map[string]int64)
	genInfoList = make(map[string][]ControllerComments)
}

// 自动路由解析入口
func (this *Router) Auto(controls ...interface{}) {
	for _, c := range controls {
		skip := make(map[string]bool, 10)
		reflectVal := reflect.ValueOf(c)
		t := reflect.Indirect(reflectVal).Type()
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			panic("you are in dev mode. So please set gopath")
		}

		pkgpath := ""
		wgopath := filepath.SplitList(gopath)
		for _, wg := range wgopath {
			wg, _ = filepath.EvalSymlinks(filepath.Join(wg, "src", t.PkgPath()))
			if fileExists(wg) {
				pkgpath = wg
				break
			}
		}

		if pkgpath != "" {
			if _, ok := skip[pkgpath]; !ok {
				skip[pkgpath] = true
				this.parserPkg(pkgpath, t.PkgPath())
			}
		}
	}

	genRouterCode()
}

func (this *Router) parserPkg(pkgRealpath string, pkgpath string) error {
	commentFilename = "auto.go"

	fileSet := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fileSet, pkgRealpath, func(info os.FileInfo) bool {
		name := info.Name()
		return !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
	}, parser.ParseComments)

	if err != nil {
		return err
	}

	for _, pkg := range astPkgs {
		for _, fl := range pkg.Files {
			for _, d := range fl.Decls {
				switch specDecl := d.(type) {
				case *ast.FuncDecl:
					if specDecl.Recv != nil {
						parserComments(specDecl.Doc, specDecl.Name.String(),
							fmt.Sprint(specDecl.Recv.List[0].Type.(*ast.StarExpr).X), pkgpath)
					}
				}
			}
		}
	}

	return nil
}

func parserComments(comments *ast.CommentGroup, funcName, controllerName, pkgpath string) error {
	if comments != nil && comments.List != nil {
		for _, c := range comments.List {
			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
			if strings.HasPrefix(t, "@router") {
				elements := strings.TrimLeft(t, "@router ")
				el := strings.SplitN(elements, " ", 3)

				if len(el) < 1 {
					return errors.New("you should has router infomation")
				}

				key := pkgpath + ":" + controllerName
				cc := ControllerComments{}
				cc.Method = funcName

				// 第一个固定为路由地址
				cc.Router = el[0]

				// 如果长度为3,第三个为[method],第二个为(prefix)
				if len(el) == 3 {
					cc.AllowHTTPMethods = strings.Split(strings.Trim(el[2], "[]"), ",")
					cc.PrefixMethods = strings.Split(strings.Trim(el[1], "()"), ",")
				} else if len(el) == 2 {
					// 长度为2,第二个为(perfix)或[method]
					if strings.HasPrefix(el[1], "(") {
						// prefix
						cc.PrefixMethods = strings.Split(strings.Trim(el[1], "()"), ",")
						cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "Get")
					} else {
						// method
						cc.AllowHTTPMethods = strings.Split(strings.Trim(el[1], "[]"), ",")
					}
				} else {
					// 长度为1,默认get方法
					cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "Get")
				}

				genInfoList[key] = append(genInfoList[key], cc)
			}
		}
	}
	return nil

}

func genRouterCode() {
	os.Mkdir(path.Join(workPath, "router"), 0755)

	var globalInfo string
	var packageInfo string

	for k, cList := range genInfoList {
		pathAndControllerName := strings.Split(k, ":")
		packagePathArr := strings.Split(pathAndControllerName[0], "/")
		packageName := packagePathArr[len(packagePathArr)-1]
		packageInfo = packageInfo + `"` + pathAndControllerName[0] + `"
	`

		// init obj
		globalInfo = globalInfo + `
    ` + packageName + ` := ` + packageName + `.New()`

		// 注释路由
		for _, c := range cList {
			if len(c.AllowHTTPMethods) > 0 {
				// 解析前缀路由
				var prefix = ""
				for pk, _ := range c.PrefixMethods {
					switch strings.ToLower(c.PrefixMethods[pk]) {
					case "csrf":
						prefix += "as.Lib.Csrf.Validate, "
					case "captcha":
						prefix += "as.Lib.Captcha.Check, "
					default:
						prefix += packageName + "." + strings.Title(strings.ToLower(c.PrefixMethods[pk])) + ", "
					}
				}

				// add func
				for _, m := range c.AllowHTTPMethods {
					globalInfo = globalInfo + `
    as.Router.Routes.` + strings.Title(strings.ToLower(strings.TrimSpace(m))) + `("/api` + c.Router + `", ` +
						prefix +
						packageName + `.` + strings.TrimSpace(c.Method) + `)`
				}
			}
		}

		globalInfo = globalInfo + `
		`

		// restful api （最后匹配 且设置了自动restful）
		/*
					restful := [][]string{
						[]string{"Get", "", "Gets"},
						[]string{"Get", "/:id", "Get"},
						[]string{"Post", "", "Add"},
						[]string{"Patch", "/:id", "Update"},
						[]string{"Delete", "/:id", "Delete"},
					}

				restfulLoop:
					for _, rest := range restful {
						// 跳过手动复写的方法
						for _, c := range cList {
							if strings.ToLower(rest[2]) == strings.ToLower(c.Method) {
								continue restfulLoop
							}
						}

						// 解析前缀路由
						var prefix = ""

						// 默认开启csrf，只对非get方法有效
						if useCsrf && rest[0] != "Get" {
							prefix += "as.Lib.Csrf.Validate, "
						}

						globalInfo = globalInfo + `
			    as.Router.Routes.` + rest[0] + `("/api/` + packageName + `s` + rest[1] + `", ` +
							prefix +
							packageName + `.` + rest[2] + `)`
					}
		*/
	}

	if globalInfo != "" && packageInfo != "" {
		f, err := os.Create(path.Join(workPath, "router", commentFilename))
		if err != nil {
			panic(err)
		}

		defer f.Close()
		output := globalRouterTemplate
		output = strings.Replace(output, "{{.globalInfo}}", globalInfo, -1)
		output = strings.Replace(output, "{{.packageInfo}}", packageInfo, -1)
		f.WriteString(output)

	}

}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getpathTime(pkgRealpath string) (lastupdate int64, err error) {
	fl, err := ioutil.ReadDir(pkgRealpath)
	if err != nil {
		return lastupdate, err
	}

	for _, f := range fl {
		if lastupdate < f.ModTime().UnixNano() {
			lastupdate = f.ModTime().UnixNano()
		}
	}
	return lastupdate, nil
}
