package test

import "errors"

/**
 * @author
 * @date 2020/10/11 15:27
 * create by Goland
 * @version 1.0
 */

// error 包

// 接口定义
type error interface {
	Error() string
}


/// errors
// 更简单的对错误自定义
// New

//DEMO1:
var(
	AError = errors.New("a")
	BError = errors.New("b")
	CError = errors.New("c")
)
 
