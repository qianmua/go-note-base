package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
 * @author
 * @date 2020/10/29 15:27
 * create by Goland
 * @version 1.0
 */



func sayHello(w http.ResponseWriter , r *http.Request){
	r.ParseForm() // 解析参数
	fmt.Println(r.Form) // output service
	fmt.Println("path" , r.URL.Path)
	fmt.Println("path" , r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("k" , k)
		fmt.Println("v" , strings.Join(v,""))
	}

	// response
	fmt.Fprintf(w , "hello ,this is first go web for http.")
}

func main() {
	// route
	http.HandleFunc("/" , sayHello)
	//listen
	err := http.ListenAndServe(":9090", nil)
	if err !=nil {
		log.Fatal("listenAndServer: " , err)
	}

}

