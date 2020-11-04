package go_xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

/**
 * @author
 * @date 2020/11/4 17:12
 * create by Goland
 * @version 1.0
 */ 

// 解析xml

// func Unmarshal(data []byte, v interface{}) error

// data接收的是XML数据流，
// v是需要输出的结构，定义为interface，
// 也就是可以把XML转换为任意的格式。
// 我们这里主要介绍struct的转换，因为struct和XML都有类似树结构的特征。


// 基础结构体
// 基础标签
type RecServices struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs	[]server `xml:"server"`
	Description string `xml:",innerxml"`
}
// value
type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func TestXml(t *testing.T) {
	file, err := os.Open("service.xml")
	if err != nil {
		fmt.Println("error" , err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("err " , err)
		return
	}
	v := RecServices{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("err" , err)
		return
	}
	fmt.Println(v)

}
