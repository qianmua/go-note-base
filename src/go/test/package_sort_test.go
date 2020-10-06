package test

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * @author
 * @date 2020/10/6 17:04
 * create by Goland
 * @version 1.0
 */ 
 

// 通用排序函数 sot包
// 实现 sort.interface

type Interface interface {
	Len() int
	Less(i  , j int ) bool
	Swap( i , j int)
}

// api use
// 通用 接口 实现


// 排序
func Sort(data Interface){

}
// 排序 保证 排序算法稳定
func Stable( data Interface){

}
// 对 一个 slice 排序
func Slice(slice interface{} , less func( i , j int) bool){

}
// 反转
func Reverse(data Interface)Interface{
	return nil
}
// 判断列表是否有序
func IsSorted(data Interface) bool{
	return false
}
// 在一个有序列表进行二分查找
func Search( n int ,  f func(int) bool)int {

	return 0
}



/// DEMO:

// slice

type Person struct {
	Name string
	Age int
}

func TestEg(t *testing.T) {
	data := []Person{
		{"A" , 20},
		{"B" , 15},
		{"C" , 18},
	}

	sort.Slice(data , func( i , j int) bool{
		return data[i].Age < data[j].Age
	})

	for _ ,e := range data {
		fmt.Println("name :" , e.Name , "Age:" , e.Age)
	}

}



