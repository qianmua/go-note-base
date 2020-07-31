package main

import "fmt"

/**
 * @author
 * @date 2020/7/31 13:03
 * create by Goland
 * @version 1.0
 */


/*
	类型递归
自己调用自己

针对 成员 而不是针对自己
指针
有点像 链表解构
循环链表


*/

// 函数 递归
func m1()  {
	m1()
}

// -----------------------------

/*
map
字典树
对 value 递归
*/
// 类型定义
type (
	Char = rune
	Trie map[Char]Trie
)

/*
有限状态自动机
使用map构造 转移状态
*/


//-----------------------------

/*
fun 应用
栈
stack
类似 链表栈
top
pop
*/
type(
	T = int
	Stack func()(top T , pop Stack)
)
// 压栈
func Push(stk Stack , x T)Stack{
	return func()(T , Stack){
		return x , stk
	}
}
// 遍历
func IterStack(stk Stack , vis func(T)){
	res := stk
	for res != nil {
		var x T
		x , res = res()
		vis(x)
	}
}

//---------------------------------------

/*
指针
表现出 指针特点 ， 又表现出 非指针特点
解 指针 指向 自己
*/
type Pointer *Pointer

func m2(){
	var p *Pointer
	// 强制转换
	// 强类型
	p = (*Pointer) (&p)
	fmt.Println(p)
}
//---------------------------------------
/*
slice
切片
好括号列？
*/
type Balance []Balance

func m3(){
	arr := Balance{{} , {{},{},{}} ,{}}
	fmt.Println(arr)
}

//---------------------------------------
/*
Y 组合子
// 实现 无显示递归
将自己 作为自己的 一个 参数
*/
type DFunc func(DFunc , int)int

func fact(impl DFunc , n int)int{
	if n == 0 {
		return 1
	}
	return n * impl(impl , n - 1)
}

//---------------------------------------
/*
church 数
//
套娃 数
*/








