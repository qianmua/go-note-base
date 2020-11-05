package go_web_socket

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

/**
 * @author
 * @date 2020/11/5 16:53
 * create by Goland
 * @version 1.0
 */

// socket 编程

// 网络的Socket数据传输是一种特殊的I/O

// Socket类型

// 流式Socket（SOCK_STREAM）和数据报式Socket（SOCK_DGRAM）

// 流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；
//数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用


// 网络中的进程之间如何通过Socket通信
// 三元组（ip地址，协议，端口）就可以标识网络的进程


// TCP Socket和UDP Socket


// TCP client
func checkError(err error){
	if err !=nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
// client
func TestTCPClientM1(t *testing.T) {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	//  host:port
	service := "127.0.0.1:7777"
	// ipv4
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// header
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

// service
func handleClient(conn net.Conn){

	// time sync
	defer conn.Close()
	s := time.Now().String()
	conn.Write([]byte(s))
}

func TestServerM2(t *testing.T) {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

