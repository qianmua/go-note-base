package go_password

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"testing"
)

/**
 * @author
 * @date 2020/11/9 17:42
 * create by Goland
 * @version 1.0
 */ 
 
// password

// 单向hash
func TestPswM1(t *testing.T) {
	h := sha256.New()
	io.WriteString(h, "1")
	fmt.Printf("%x \n", h.Sum(nil))

	//import "crypto/sha1"
	h = sha1.New()
	io.WriteString(h, "1")
	fmt.Printf("%x \n", h.Sum(nil))

	//import "crypto/md5"
	h = md5.New()
	io.WriteString(h, "1")
	fmt.Printf("%x \n", h.Sum(nil))
}


// 加盐hash
func TestPswM2(t *testing.T) {
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "123456")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

func TestPswM3(t *testing.T) {
	// encrypt=bdcbaf30203dd142cd1d40f01978b96750ca166d09289ccdd64aff94dac0cda3
	// encrypt=bdcbaf30203dd142cd1d40f01978b96750ca166d09289ccdd64aff94dac0cda3
	// encrypt=bdcbaf30203dd142cd1d40f01978b96750ca166d09289ccdd64aff94dac0cda3
	// encrypt=0ae24b4b6295bce10ab7e4aefbd46ba1c1b4303fee687151ea38a7551d2d2350
	salt := make([]byte, 32)
	salt = []byte("加油！")
	fmt.Printf("salt=%v\n", salt)
	hash , _ := scrypt.Key([]byte("some password"),salt , 16384, 8, 1, 32)
	fmt.Printf("encrypt=%x\n", hash)
}