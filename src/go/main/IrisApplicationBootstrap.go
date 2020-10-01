package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iris-contrib/middleware/jwt"
	jsoniter "github.com/json-iterator/go"
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
	app.OnErrorCode(iris.StatusInternalServerError , error)

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

/**
	struct
 */
// controller
type MyController struct {}
type Token struct {
	Token string `json:"token"`
}

/**
	router

 */
// 路由
func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// method,path,funcName,middleware
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler",hello)
	b.Handle("GET", "/hello/{name:string}", "Hello",hello2)
	b.Handle("GET", "/sayHello/{text:string}", "SayHelloText",hello3)
	b.Handle("GET", "/mysqlData/{text:string}", "SayHelloText",hello3)
	// myHandler
}

func (m *MyController) Get() string {
	return "Hello World"
}

/**
	动态传参
	params 解析
 */
func (m *MyController) MyCustomHandler(id int64) string {
	fmt.Println(id)
	return "id:" + string(id)
}
func (m *MyController) Hello(name string) string {
	return "your name is " + name
}
func (m *MyController) SayHelloText(text string) string {
	data := openMysql()
	bytes, err := jsoniter.Marshal(data)
	if err !=nil {
		fmt.Println("err" , err.Error())
		return "null"
	}
	return string(bytes)

}
//	data := openMysql()
/*
func (m *MyController) mysqlQuery() string {
	return " jsonValue"
}
*/

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
	ctx.Next()
}

func notFound(ctx iris.Context) {
	// 数据填充
	//ctx.ViewData("message","Hello world!")
	if true {
		err := ctx.View("hello.html")
		if err != nil {
			fmt.Println("err -> " + err.Error())
		}
	}
}

func error(ctx iris.Context){
	err := ctx.View("500.html")
	if err != nil {
		fmt.Println("err -> " + err.Error())
	}
}

func myHandler( ctx iris.Context){
	user := ctx.Values().Get("jwt").(*jwt.Token)
	//data := openMysql()
	ctx.JSON(user)
}

/**
	api
 */

func openMysql() map[int]map[string]string{
	fmt.Println("open mysql db")

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("link error -> " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("err -> ping :" , err.Error())
		return nil
	}

	// close db
	defer db.Close()

	// query
	//rows , _ := db.Query("select * from users")
	//id := 0
	//name := ""
	/*for rows.Next()  {
	}*/
	//fmt.Println(rows)

	//query all
	rows , _ := db.Query("select * from users;")
	if rows == nil {
		return nil
	}
	// 所有列
	columns, _ := rows.Columns()
	// 一行所有列的值
	vals := make([][]byte , len(columns))
	// 一行的填充数据
	scans := make([]interface{}  , len(columns))

	// 数据 填充
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
		for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := columns[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return result
}

