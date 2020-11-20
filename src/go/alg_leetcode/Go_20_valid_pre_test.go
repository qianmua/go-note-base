package alg_leetcode

import "errors"

/**
 * @author
 * @date 2020/11/20 17:35
 * create by Goland
 * @version 1.0
 */ 

// 括号匹配问题。


// 栈！
type Stack struct{
	top int
	size int
	element []interface {}
}
/**
  判断stack是否为空
*/
func (stack *Stack)IsEmpty() bool{
	return stack.top==stack.size-1
}
/**
  判断stack是否已经满了
*/
func (stack *Stack)IsFull() bool{
	return stack.top == 0
}
/**
  弹出一个元素
*/
func (stack *Stack)Pop() (interface {},error){
	if stack.IsEmpty()==true{
		return nil, errors.New("The Stack is empty")
	}
	stack.top=stack.top+1
	return stack.element[stack.top-1], nil
}

/**
	查看下一个元素
 */

func (stack *Stack)Peek() interface {} {
	if stack.IsEmpty()==true{
	}
	//stack.top=stack.top+1
	return stack.element[stack.top-1]
}
/**
  压入一个元素
*/
func (stack *Stack)Push(e interface {}) error{
	if stack.IsFull()==true{
		return errors.New("The Stack is full")
	}
	stack.top=stack.top-1
	stack.element[stack.top]=e
	return nil
}




func isValid(s string) bool{
	var brackets Stack
	for i := 0; i < len(s) ;i++  {
		u := s[i]
		switch u {
		case '(':
		case '[':
		case '{':
			brackets.Push(u)

		case ')':
			if !brackets.IsEmpty() {

			}
			

		}
	}
}
