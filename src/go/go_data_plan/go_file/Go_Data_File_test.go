package go_file

import (
	"fmt"
	"os"
	"testing"
)

/**
 * @author
 * @date 2020/11/5 16:21
 * create by Goland
 * @version 1.0
 */

// file 处理

func TestFileM1(t *testing.T) {
	// 文件名 权限
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("astaxie")

}


// 文件操作

// 新建
// Create
// NewFile

// 打开
// Open
// OpenFile

// 写文件
// Write
// WriteAt
// WriteString

func TestFileM2(t *testing.T) {
	s := "demofile.txt"
	file, err := os.Create(s)
	if err != nil {
		fmt.Println("err" , err)
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		file.WriteString("Just a test!\r\n")
		file.Write([]byte("Just a test!\r\n"))
	}
}

// 读文件
// Read
// ReadAt

func TestFileM3(t *testing.T) {
	s := "demofile.txt"
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("err" , err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])

	}

}

// 删除文件
// Go语言里面删除文件和删除文件夹是同一个函数
// func Remove(name string) Error