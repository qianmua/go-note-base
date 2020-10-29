package test

import (
	"fmt"
	"index/suffixarray"
	"sort"
	"testing"
)

/**
 * @author
 * @date 2020/10/11 15:29
 * create by Goland
 * @version 1.0
 */ 

// index包 ， 提供基于前缀字串检索功能
// 在byte 数组检索指定字串 ，并得到下标

// create
// New
// Bytes 得到原始byte数组


// 检索
// FindAllIndex
// Lookup

// DEMO1:
func TestIM1(t *testing.T) {
	source := []byte("hello world , hello Go!")
	index := suffixarray.New(source)

	offsets := index.Lookup([]byte("hello qianmua" ), -1)

	sort.Ints(offsets)

	fmt.Printf("%v" , offsets)
}