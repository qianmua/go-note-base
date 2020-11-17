package test

import (
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/11/18 0:49
 * create by Goland
 * @version 1.0
 */ 

// 切片


// 模拟栈和队列

func TestStackM1(t *testing.T) {
	// create
	stack :=make([]int , 0)

	// push
	stack = append(stack, 10)
	// pop
	v := stack[len(stack) -1]
	stack = stack[:len(stack) -1]

	fmt.Println(v)
}

// queue
func TestQueueM1(t *testing.T) {
	queue := make([]int , 0)

	// 入队
	queue = append(queue , 10)

	// 出队
	v := queue[0]
	queue = queue[1:]

	fmt.Println(v)
}


// 字典
// map[string]int


// 库
// sort
// math
// copy
