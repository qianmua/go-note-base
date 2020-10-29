package test

import (
	"fmt"
	"io/ioutil"
	"testing"
)

/**
 * @author
 * @date 2020/10/11 15:46
 * create by Goland
 * @version 1.0
 */ 
 
// ioUtil

// api
// ReadAll 从io.Reader 读取所有数据， 返回字节数组
// ReadllDir 从目录读，返回目录所有文件列表
// ReadFile 读取文件内容，返回字节数组

// DEMO1:
func TestIOUtilM1(t *testing.T) {
	content , err := ioutil.ReadFile("demo.text")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(content)
}


// 临时文件/目录
// TempDir/TempFile 创建
// 使用 os.remove 删除文件， 不会自动销毁


// 文件写入
// WriteFile
// params： fileName , data , perm文件信息标识位