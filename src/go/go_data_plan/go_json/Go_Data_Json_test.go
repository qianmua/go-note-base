package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/11/5 13:56
 * create by Goland
 * @version 1.0
 */

// go 操作json

// {"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}

// 解析 func Unmarshal(data []byte, v interface{}) error

type Server struct {
	ServerName string `json:"server_name"`
	ServerIP string `json:"server_ip"`
}

type ServerSlice struct {
	Servers []Server `json:"servers"`
}

/**

首先查找tag含有Foo的可导出的struct字段(首字母大写)
其次查找字段名是Foo的导出字段
最后查找类似FOO或者FoO这样的除了首字母之外其他大小写不敏感的导出字段

能够被赋值的字段必须是可导出字段(即首字母大写）

JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略

 */
func TestJsonM1(t *testing.T) {
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(str) , &s)
	fmt.Println(s)
}

// 解析到interface

// JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。
//Go类型和JSON类型的对应
/**
bool 代表 JSON booleans,
float64 代表 JSON numbers,
string 代表 JSON strings,
nil 代表 JSON null.
 */


// b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

func TestJsonM2(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	/**
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
	 */
	// 断言：
	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}


}


// 生成JSON

func TestJsonM3(t *testing.T) {
	var s ServerSlice
	s.Servers = append(s.Servers , Server{ServerName:"ASD" , ServerIP:"127.0.0.4"})
	s.Servers = append(s.Servers , Server{ServerName:"QWE" , ServerIP:"47.0.230.4"})
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(data))
}

// 定义struct tag
/*

字段的tag是"-"，那么这个字段不会输出到JSON

tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName

tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中

如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，
那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串

 */




 
