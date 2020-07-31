package main

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











