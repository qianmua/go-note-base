package go_regexp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
)

/**
 * @author
 * @date 2020/11/5 14:44
 * create by Goland
 * @version 1.0
 */

// regexp
// 正则表达式


func TestRegexM1(t *testing.T) {
	// 输入是不是IP地址

	f := func(ip string)(b bool) {
		if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
			return false
		}
		return true
	}

	fmt.Println(f("127.0.0.1"))

	f1 := func() {
		if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
			fmt.Println("数字")
		} else {
			fmt.Println("不是数字")
		}
	}
	f1()

}


// 通过正则获取内容

func TestRegexM2(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err!=nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err")
		return
	}
	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))

}
 
