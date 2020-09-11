package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

/**
 * @author
 * @date 2020/7/19 23:21
 * create by Goland
 * @version 1.0
 */


func stringTest() {
	s1 := "hello哈哈"
	s2 := []rune(s1)
	fmt.Println(string(s2[0]))
}

func main() {
	app := iris.New()

	// 配置
	mvc.Configure(app.Party("/root"),myMvc)

	html := iris.HTML("./template", ".html")
	app.RegisterView(html)
	// 错误码
	app.OnErrorCode(iris.StatusNotFound , notFound)

	// start and set charset
	app.Run(iris.Addr(":9090"),iris.WithCharset("UTF-8"))
}

/**
	mvc start
 */

func myMvc(app *mvc.Application) {
	// handle
	app.Handle(new(MyController))
}

// controller
type MyController struct {}

// 再添加路由
func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// method,path,funcName,middleware
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler",hello)
	b.Handle("GET", "/hello/{name:string}", "Hello",hello2)
	b.Handle("GET", "/sayHello/{text:string}", "SayHelloText",hello3)
}

func (m *MyController) Get() string {
	return "Hello World"
}

/**
	动态传参
 */
func (m *MyController) MyCustomHandler(id int64) string {
	fmt.Println(id)
	return "id:" + string(id)
}
func (m *MyController) Hello(name string) string {
	return "your name is " + name
}
func (m *MyController) SayHelloText(text string) string {
	return " say :" + text
}
/**
	view
 */
func (m *MyController) helloView(text string) string {
	return " say :" + text
}
/**
	handle 处理
	每次 调用接口 都会触发 handle

	通过 Next 控制向下传递
	相当于 tomcat 阀门
 */
func hello(ctx iris.Context) {
	fmt.Println("use base")
	ctx.Next()
}

func hello2(ctx iris.Context) {
	fmt.Println("print name")
	ctx.Next()
}

func hello3(ctx iris.Context) {
	fmt.Println("say :")
	ctx.Next()
}

func notFound(ctx iris.Context) {
	if true {
		err := ctx.View("hello.html")
		if err != nil {
			fmt.Println("err -> " + err.Error())
		}
	}
}
 
