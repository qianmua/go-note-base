package go_crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/11/10 13:41
 * create by Goland
 * @version 1.0
 */ 

// base 加密

func base64Encode(src []byte) []byte{
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte , error){
	return base64.StdEncoding.DecodeString(string(src))
}

func TestBase64M1(t *testing.T) {
	hello := "hello world 你好，世界！"
	debyte := base64Encode([] byte(hello))
	fmt.Println(debyte)

	enbyte, err := base64Decode(debyte)
	if err != nil{
		fmt.Println(err)
	}
	if hello != string(enbyte){
		fmt.Println("no")
	}
	fmt.Println(string(enbyte))

}