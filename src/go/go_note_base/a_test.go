package test

import (
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/10/1 22:04
 * create by Goland
 * @version 1.0
 */ 
 
func TestHelloWorld(t *testing.T){
	t.Log("hello world")

}
// 多个 测试用例

/**
	一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级... 它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

	动态规划
 */
func TestDongtaiGuiHuaA1(t *testing.T) {


}

/**
	true

	"+100"
	"5e2"
	"-123"
	"3.1416"
	"-1E-16"

	false

	"12e"
	"1a3.14"
	"1.2.3"
	"+-5"
	"12e+4.3"
 */
func TestIsNumber(t *testing.T) {

}

// 快排
func TestQuickSort(t *testing.T) {
	ints := make([]int, 4)
	ints[0] =5
	ints[1] =6
	ints[2] =1
	ints[3] =4
	quick(ints , 0 , len(ints) -1)
	fmt.Println(ints)
}

func quick(arr []int , s , e int) {
	if s > e {
		return
	}

	left , right := s , e
	mid := arr[left]

	for left < right {

		for left < right && arr[right] >= mid {
			right --
		}

		arr[left] = arr[right]

		for left < right && arr[left] < mid {
			left ++
		}

		arr[right] = arr[left]

	}
	arr[left] = mid
	quick(arr , s , left -1)
	quick(arr ,left + 1 , e)

}




