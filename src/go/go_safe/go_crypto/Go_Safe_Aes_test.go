package go_crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"testing"
)

/**
 * @author
 * @date 2020/11/10 14:10
 * create by Goland
 * @version 1.0
 */ 
 
// aes

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func TestAESM1(t *testing.T) {

	name := []byte("name is qianmuna")
	//if len(os.Args) > 1 {
	//	name = []byte(os.Args[1])
	//}

	//aes的加密字符串
	// 16 24 32
	// aes128 aes192 aes256
	key_text := "qianmua12798akljzmknm.ahkjkljl;k"
	/*if len(os.Args) > 2 {
		key_text = os.Args[2]
	}*/

	fmt.Println(len(key_text))

	// create aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}


	// encode
	cfbe := cipher.NewCFBEncrypter(c, commonIV)
	var1 := make([]byte, len(name))
	cfbe.XORKeyStream(var1 , name)
	// var1 加密后
	fmt.Printf("%s=>%x\n", name, var1)

	//decode
	cfbd := cipher.NewCFBDecrypter(c, commonIV)
	var2 := make([]byte, len(name))
	//var var3 []byte
	cfbd.XORKeyStream(var2 , var1)

	fmt.Printf("%x=>%s\n", var1, var2)

}