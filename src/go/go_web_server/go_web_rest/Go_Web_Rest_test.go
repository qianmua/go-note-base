package go_web_rest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"testing"
)

/**
 * @author
 * @date 2020/11/6 15:37
 * create by Goland
 * @version 1.0
 */ 
 
// rest

// RESTful架构：
/**

（1）每一个URI代表一种资源；
（2）客户端和服务器之间，传递这种资源的某种表现层；
（3）客户端通过四个HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"

 */

// 资源我们通过URI来定位，也就是一个URI表示一个资源。
// 实体展现出来就是表现层，例如一个txt文本信息，他可以输出成html、json、xml等格式，一个图片他可以jpg、png等方式展现，这个就是表现层的意思。
// 状态转化（State Transfer）
// 客户端能通知服务器端的手段，只能是HTTP协议。
//具体来说，就是HTTP协议里面，四个表示操作方式的动词：GET、POST、PUT、DELETE。
//它们分别对应四种基本操作：GET用来获取资源，POST用来新建资源（也可以用于更新资源），PUT用来更新资源，DELETE用来删除资源。

func Index(w http.ResponseWriter , r *http.Request , _ httprouter.Params){
	fmt.Fprint(w, "Welcome!\n")
	fmt.Fprint(w , "<h1>hello world!</h1>\n")
}
func Hello(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func getuser(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}
func adduser(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}
func deleteuser(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}
func modifyuser(w http.ResponseWriter , r *http.Request , ps httprouter.Params){
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func TestRestMain(t *testing.T) {
	// base router
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/user/:uid", getuser)

	router.POST("/adduser/:uid", adduser)

	router.DELETE("/deluser/:uid", deleteuser)

	router.PUT("/moduser/:uid", modifyuser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
