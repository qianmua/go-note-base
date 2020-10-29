package test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

/**
 * @author
 * @date 2020/10/11 15:35
 * create by Goland
 * @version 1.0
 */


// io
// Reader base

// LimitReader 读指定长度
// MultiReader 聚合多个Reader
// TeeReader	将一个输入绑定到一个输出

//DEMO1:
func TestIOM1(t *testing.T) {
	r := strings.NewReader("io Stream read \n")
	lr := io.LimitReader(r , 4)
	written, err := io.Copy(os.Stdout, lr)
	if err != nil {

	}
	fmt.Println(written)
}

//ReadAtLeast  ReadFull
// 从 Reader 读数据到缓冲区， 读N ， 读满



// Writer
// 对象输出设备

// MultiWriter

// WriteString
// 向 writer写入字符串



// ReadWriter 可以进行读和写


// Copy Writer to Reader 直到EOF

// CopyBuffer
// CopyN


// PipeReader，PipeWriter，ByteReader，ByteWriter等针对具体类型的实例和一些辅助函数

 
