package go_web_iris

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/macro"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"log"
	"net"
	"strconv"
	"testing"
)

/**
 * @author
 * @date 2020/11/11 16:38
 * create by Goland
 * @version 1.0
 */

// Demo1
func TestIrisM1(t *testing.T) {
	app := iris.New()

	app.Get("/" , func(ctx iris.Context){})

	err := app.Run(iris.Addr(":8080"))
	if err!= nil {
		log.Fatal("err: " , err.Error())
	}

}


//Ddmo2
func TestIrisM2(t *testing.T) {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET" , "/" , func(ctx iris.Context){
		ctx.HTML("<h1>hello world! Iris.</h1>")
	})

	app.Get("/ping" , func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello" , func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Message" : "hello Iris."})
		
	})

	app.Run(iris.Addr(":9090"))
}


// Demo3:
type ExampleController struct {}

func TestIrisM3(t *testing.T) {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app).Handle(new(ExampleController))

	app.Run(iris.Addr(":9090"))
}

// get
func (c *ExampleController) Get() mvc.Result{
	return mvc.Response{
		ContentType:"text:html",
		Text:"<h1> hello world Iris!</h1>",
	}
}

func (c *ExampleController) GetPing() string{
	return "pong"
}

func (c *ExampleController) GetHello() interface{}{
	return map[string]string{"Message" : "hello Iris."}
}


/// Demo4 listener
func TestIrisM4(t *testing.T) {
	app := iris.New()

	// 自定义监听
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	// 前置处理
	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("server terminated")
		})
	})

	app.Run(iris.Listener(l))

	// 后置处理
	//app.Run(iris.Addr(":8080", func(h *iris.Supervisor) {
	//	h.RegisterOnShutdown(func() {
	//		println("server terminated")
	//	})
	//}))

}


// Demo5 路由

func TestIrisM5(t *testing.T) {

	// 分组
	app := iris.New()

	users := app.Party("/users", myAuthMiddlewareHandler)

	// http://localhost:8080/users/42/profile
	users.Get("/{id:int}/profile", myAuthMiddlewareHandler)
	// http://localhost:8080/users/messages/1
	users.Get("/messages/{id:int}", myAuthMiddlewareHandler)

}

func myAuthMiddlewareHandler(ctx iris.Context){
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}

// Demo6 dynamic
func TestIrisM6(t *testing.T) {
	app := iris.New()

	//
	app.Get("/username/{name}" , func(ctx iris.Context){
		ctx.Writef("hello  %s" , ctx.Params().Get("name"))
	})

	// 宏函数
	// "min" = 函数名称
	// "minValue" = 函数参数
	// func(string) bool = 宏的路径参数的求值函数，当用户请求包含 min(...)
	// 宏的 :int 路径参数的 URL 时，这个函数将会执行
	macro.Int.RegisterFunc("min" , func(minValue int) func (string) bool{
		return func(paramValue string ) bool{
			n, err := strconv.Atoi(paramValue)
			if err != nil {
				return false
			}
			return n >= minValue
		}
	})

	// http://localhost:8080/profile/id>=1
	// 如果路由是这样：/profile/0, /profile/blabla, /profile/-1，将抛出404错误
	// 宏参数函数是可以省略的。
	app.Get("/profile/{id:int min(1)}" , func(ctx iris.Context){
		// // 第二个参数是错误的，但是它将是 nil，因为我们用了宏
		// // 验证已经执行了。
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("hello id : %s" , id)
	})

	// 改变路径参数错误引起的错误码：
	app.Get("/profile/{id:int min(1)}/friends/{f:int min(1) else 504}" , func(ctx iris.Context){
		id, _ := ctx.Params().GetInt("id")
		f, _ := ctx.Params().GetInt("f")

		ctx.Writef("hello id : %s , fId is : %s " , id , f)

	})

	// // 这将抛出 504 错误码，而不是404，如果所有的路由宏没通过。


	// http://localhost:8080/game/a-zA-Z/level/0-9
	// 记着，alphabetical 仅仅允许大小写字母。
	app.Get("/game/{name:alphabetical}/level/{level:int}" , func(ctx iris.Context){
		ctx.Writef("name : %s | level: %s" , ctx.Params().Get("name") , ctx.Params().Get("level"))
	})


	// 自定义正则表达式来验证一个路径参数仅仅是小写字母
	//  http://localhost:8080/lowercase/anylowercase
	app.Get("/lwc/{name:string regexp(^[a-z]+)}" , func(ctx iris.Context){
		ctx.Writef("name should lowercase : %s" ,ctx.Params().Get("name"))
	})


	// http://localhost:8080/single_file/app.js
	app.Get("/single_file/{my_file:file}" , func(ctx iris.Context) {
		ctx.Writef("file type file: %s" , ctx.Params().Get("my_file"))
	})

	// http://localhost:8080/myfiles/any/directory/here/
	// 这是唯一一个接受任何数量路径片段的宏类型。




}