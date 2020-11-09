package go_web_rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

/**
 * @author
 * @date 2020/11/9 16:05
 * create by Goland
 * @version 1.0
 */ 
 
/// web json
// rpc

// server
// 基于 tcp
func TestJsonM1(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8848")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err !=nil {
			continue
		}

		jsonrpc.ServeConn(conn)
	}
}

// client
func TestJsonClientM2(t *testing.T) {
	service := "127.0.0.1:8848"

	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing :" , err)
		
	}

	// sync call
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
