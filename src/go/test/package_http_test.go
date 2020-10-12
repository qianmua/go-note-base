package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

/**
 * @author
 * @date 2020/10/11 16:10
 * create by Goland
 * @version 1.0
 */

// http
// 组成
// Request
// Response
// Client
// Server

// api
// resp, err := http.Get("http://example.com/") // GET

// resp, err := http.Post("http://example.com/") // POST

// resp, err := http.PostForm("http://example.com/", url.Values{"foo": "bar"}) // 提交表单


// HTTP请求和响应
// Request 主要用于数据的存储
// Response Response对象拥有了当前Request对象的引用


// Client
func TestHttpM1(t *testing.T) {
	//client := http.Client()
	//res, err := client.Get("http://www.google.com")
}
// base api
// Get
// Head
// Post
// PostForm

// Do 传入 Request 提供请求定制化

// DEMO2:
func TestHttpM2(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Cookie" , "name=foo")
	request.Header.Set("Content-Type" , "application/x-www-form-urlencoded")
	client := http.Client{}
	res, err := client.Do(request)

	// defer res.Body().Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}

// Server
// 创建http服务器 对外提供访问

// ListenAndServe

// DEMO3:
func TestHttpM3(t *testing.T) {
	 http.HandleFunc("/echo/", EchoServer)
	log.Fatal(http.ListenAndServe(":8080" , nil))
}
func EchoServer(w http.ResponseWriter , req *http.Request){
	all, _ := ioutil.ReadAll(req.Body)
	n, _ := io.WriteString(w, string(all))
	fmt.Println(n)
}




 
