package test

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/**
 * @author
 * @date 2020/10/29 15:43
 * create by Goland
 * @version 1.0
 */ 


/**

表单处理
 */
//func sayHello(w http.ResponseWriter , r *http.Request){
//	r.ParseForm() // 解析参数
//	fmt.Println(r.Form) // output service
//	fmt.Println("path" , r.URL.Path)
//	fmt.Println("path" , r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("k" , k)
//		fmt.Println("v" , strings.Join(v,""))
//	}
//
//	// response
//	fmt.Fprintf(w , "hello ,this is first go web for http.")
//}

func login(w http.ResponseWriter , r *http.Request){
	fmt.Println("method: " ,r.Method)
	if r.Method == "GET"{
		t,_ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	}else {
		// 解析 form
		r.ParseForm()
		fmt.Println("userName: " , r.Form["username"])
		fmt.Println("password: " , r.Form["password"])
	}
}


//func main() {
//	// route
//	http.HandleFunc("/" , sayHello)
//	http.HandleFunc("/login" , login)
//	//listen
//	err := http.ListenAndServe(":9090", nil)
//	if err !=nil {
//		log.Fatal("listenAndServer: " , err)
//	}
//
//}