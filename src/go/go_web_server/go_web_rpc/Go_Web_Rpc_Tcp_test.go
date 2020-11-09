package go_web_rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"testing"
)

/**
 * @author
 * @date 2020/11/9 13:53
 * create by Goland
 * @version 1.0
 */ 
 

// rpc tcp
// service
func TestTcpM1(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8848")
	checkError(err)

	listenTCP, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listenTCP.Accept()
		if err != nil {
			continue
		}
		// 处理阻塞型的单用户的程序
		// 使用goroutine
		go rpc.ServeConn(conn)
	}
}
// client
func TestTcpClientM2(t *testing.T) {

	service := "127.0.0.1:8848"

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)


	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)


}



func checkError(err error){
	if err != nil {
		fmt.Println("err " , err.Error())
		os.Exit(1)
	}
}
