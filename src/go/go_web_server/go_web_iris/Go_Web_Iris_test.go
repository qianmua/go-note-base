package go_web_iris

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"log"
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
