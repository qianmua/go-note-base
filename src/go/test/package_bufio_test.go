package test

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

/**
 * @author
 * @date 2020/10/6 17:25
 * create by Goland
 * @version 1.0
 */


/// bufio模块
// 缓冲区
//


/// Reader
// 创建 NewReader

// 函数声明

func NewReader(rd io.Reader) *io.Reader{
	return nil
}
func NewReaderSize(rd io.Reader , size  int) *io.Reader{
	return nil
}

// io.Reader bufio.Reader
// base api
// Read 读取n个byte数据
// Discard 丢弃接下来n个数据
// Peek 获取当前缓冲区 接下来的n个byte ， 不移动指针
// Rest 清空整个缓冲区

// 高级api
// ReadByte	读取一个byte
// ReadRune 读取一个utf-8字符
// ReadLine 读取一行数据 \n 分割
// ReadBytes 读取一个byte列表
// ReadString 读取一个字符串


// DEMO1:
func TestDemo1(t *testing.T){
	r := strings.NewReader("hello world!")

	reader := bufio.NewReader(r)

	for{
		readString, err := reader.ReadString(byte(' '))
		fmt.Println(readString)
		if err != nil {
			return
		}
	}
}



// Scanner
// Scanner 读取数据

// splitFunc api
// ScanBytes	按照byte拆分
// ScanLines 	按照 ‘\n’拆分
// ScanRunes	按照utf8字符拆分
// ScanWords	按照单词 “” 拆分

//Split 接收一个参数
// 参数类型为 splitFunc

// 读
// Scan
// 取
// Text
// Bytes

// DEMO2:

func TestDemo2(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("hello world!"))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



// Writer
// 与 Reader类型

// base api
// Write 写入 n个数据
// Reset 重置当前缓冲区
// Flush 清空当前缓冲区


// 数据写入
// WriteByte	写入一个 字节
// WriteRune 	写入一个字符
// WriteString	写入一个 字符串

 
