package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/**
 * @author
 * @date 2020/10/12 16:23
 * create by Goland
 * @version 1.0
 */

// sync

// atomic
// int32 int64 uint32 uint64 uintptr unsafe unsafe.Pointer
// 原子操作
// Add 添加或者减少
// CAS cas
// Swap 交换
// Load 读取
// Store 储存

// AddInt32 AddInt64.....

// DEMO1:
func TestAtomicM1(t *testing.T) {
	var a int32
	a += 10
	atomic.AddInt32(&a , 10)
	fmt.Println(a == 20)

	var b uint32
	b += 20

	// 无符号数
	// 补码机制
	// b -= 10
	atomic.AddUint32(&b , ^uint32(10-1))
	// b -= N
	//atomic.AddUint32(&b , ^uint32(N-1))
	fmt.Println(b == 10)

}

// CAS


// sync
// 同步


// DEMO1: ?
func TestNotSyncM1(t *testing.T) {
	data := 1
	go add1(&data)
	go add1(&data)
	time.Sleep(10)

	// 3? cloud 2? 3?
	fmt.Println(data)
}
func add1(data *int){
	temp := *data
	*data = temp + 1
}

// 互斥锁
// sync.Mutex
// api lock unlock

// DEMO2:
func TestSyncM1(t *testing.T) {
	mu := sync.Mutex{}
	data := 1
	go add2(&data , &mu)
	go add2(&data , &mu)
	time.Sleep(10)

	// 3? cloud 2? 3?
	fmt.Println(data)
}
func add2(data *int , mu *sync.Mutex){
	mu.Lock()
	defer mu.Unlock()
	temp := *data
	*data = temp + 1
}

// 读写锁 RWMutex
// 信号量 Cond
// 		Signal	Broadcast 唤醒，等待中的线程
//		Wait

// WaitGroup
//	add
//	Done
// 	Wait

// DEMO3:
func TestSyncM2(t *testing.T) {

}

 
