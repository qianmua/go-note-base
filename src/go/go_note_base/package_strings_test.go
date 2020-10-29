package test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

/**
 * @author
 * @date 2020/10/12 14:19
 * create by Goland
 * @version 1.0
 */ 
 
// strings 提供对string类型操作

// 类别
// 字符串搜索 匹配
// 字符转拆分
// 字符串修改
// 其他独立函数

// 字符串搜索，匹配，包含/Bool
// contains 检测 是否包含，字符串
// containsRune 字符
// containsAny


// Index 搜索子串，返回起始下标
// IndexByte
// IndexRule
// IndexAny

// Last 获取匹配的最后一个
// LastIndex
// LastIndexByte
// LastIndexAny


// 字符串拆分
// Split 不包含分隔符号
// SplitAfter 包含分隔符号


// 一组分隔符
// FieldFunc
// DEMO1:
func TestStringsM1(t *testing.T) {
	str := "hello& world !!! ASD"
	f := func(c rune) bool{
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println( strings.FieldsFunc(str , f))
}


// 字符串修改
// Trim系列
// 删除字符串首尾连续多余字符
// Trim	首尾
// TrimLeft 首
// TrimRight 尾
// TrimSpace 只删空格

// ToLower
// ToUpper


// Join
// Compare  通过字典比较两个字符串大小
// Count	统计子串数量
// Replace
