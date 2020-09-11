package main

import "fmt"
import "github.com/kataras/iris"
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
	fmt.Println("aaa")
	app := iris.New()

	app.Get("/" , func(ctx iris.Context){
		ctx.WriteString("hello iris")
	})

	app.Run(iris.Addr(":9090"))

}
 
