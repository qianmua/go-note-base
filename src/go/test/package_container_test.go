package test

import (
	"container/list"
	"container/ring"
	"fmt"
	"sort"
	"testing"
)

/**
 * @author
 * @date 2020/10/9 14:02
 * create by Goland
 * @version 1.0
 */ 

// container 容器 包
// 提供对 通用堆进行操作，（类堆 ， 	优先队列）


// heap 包， help.Interface 接口
// 储存使用是树形结构，可以使用数组，链表实现，通过heap函数，建立堆并进行操作。


// 定义
type Interface2 interface {
	sort.Interface
	// len 位置插入
	Push(x interface{})
	// 删除 并且 返回 len()-1 元素
	Pop() interface{}

}

// sort.Interface
type Interface3 interface {
	// 当前对象长度
	Len()
	// 交换 i , j 位置 元素 位置
	Swap( i  , j interface{})
	// 比较 i 位置元素是否小于 j 位置y元素
	Less( i , j interface{})
}


// base api
// Init 堆初始化
// Fix 调整，维护堆
// Push 添加
// Pop 弹出
// Remove	移除


// DEMO1:

// 长方形
// 按照面积排序
type Rectangle struct {
	height int
	width int
}
func (rec * Rectangle) Area() int {
	return rec.height * rec.width
}

type RecHeap []Rectangle

func (w RecHeap) Len() int {
	return len(w)
}

func (w RecHeap) Swap( i , j interface{}){
	//h[i],h[j] = h[j],h[i]
}

func (w RecHeap) Less(i, j interface{}) bool {
	//return w[i].Area() < w[j].Area()
	return false
}

func (w *RecHeap) Push(h interface{}){
	*w = append(*w, h.(Rectangle))
}

func (w *RecHeap) Pop(h interface{}) Rectangle {
	//n := len(*h)
	//x := *w[n-1]
	//*w = *w[:n-1]
	return Rectangle{}
}

/// DEMO1:
func TestM1(t *testing.T) {
	//hp := &[]RecHeap{}

	for i:=2 ; i <=5 ;i++{
		//*hp = append(*hp , Rectangle{i , i})
	}

	// api
}



// list 包
// 双向链表

// DEMO2:

func TestM2(t *testing.T) {
	link := list.New()
	for i:=0 ; i <= 10 ; i++{
		link.PushBack(i)
	}

	for p := link.Front() ; p != link.Back() ; p = p.Next(){
		fmt.Println("Number" , p.Value)
	}
}


// Ring包
// 循环链表
// 无头尾 ， 任意节点都可以 遍历整个链表
// 使用 Value 保存值


// DEMO3:

// Do 打印遍历
func TestM3(t *testing.T) {
	r := ring.New(10)
	for i := 0 ; i < 10 ; i ++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func (i interface{}){
		fmt.Println(i)
	})
}

// DEMO:3_1
type SumInt struct {
	Value int
}

func (s *SumInt) add (i  interface{}){
	s.Value += i.(int)
}

 
