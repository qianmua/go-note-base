package go_web_web_socket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"testing"
)

/**
 * @author
 * @date 2020/11/6 14:52
 * create by Goland
 * @version 1.0
 */ 
 
// websocket
// ws

// WebSocket的协议颇为简单，在第一次handshake通过以后，连接便建立成功，其后的通讯数据都是以”\x00″开头，以”\xFF”结尾。在客户端，这个是透明的，WebSocket组件会自动将原始数据“掐头去尾”。

// WebSocket连接请求，然后服务器发出回应，然后连接建立成功，这个过程通常称为“握手” (handshaking)

// 在请求中的"Sec-WebSocket-Key"是随机的 是一个经过base64编码后的数据
// f7cb4ezEAl6C3wRaU6JORA==

//  服务器端接收到这个请求之后需要把这个字符串连接上一个固定的字符串：
// 258EAFA5-E914-47DA-95CA-C5AB0DC85B11

// 即：f7cb4ezEAl6C3wRaU6JORA==连接上那一串固定字符串，生成一个这样的字符串：
// f7cb4ezEAl6C3wRaU6JORA==258EAFA5-E914-47DA-95CA-C5AB0DC85B11

// 对该字符串先用 sha1安全散列算法计算出二进制的值，然后用base64对其进行编码，即可以得到握手后的字符串：
// rE91AJhfC+6JdVcVXOGJEADEJdQ=


func Echo(ws *websocket.Conn){
	var err error
	for{
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
		// msg
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
// link
func TestWebSocketMain(t *testing.T) {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		// err
		log.Fatal("ListenAndServe:", err)
	}
}