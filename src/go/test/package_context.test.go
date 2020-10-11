package test

import (
	"context"
	"flag"
	"time"
)

/**
 * @author
 * @date 2020/10/11 15:07
 * create by Goland
 * @version 1.0
 */ 

// context 包
// 并发环境常用标准库
// 协程之间传递上下文信息


// 声明方式
type ContextTest interface {
	Deadline() (deadline time.Time , ok bool)
	// 获取 deadline
	Done() <-chan struct{}

	Err() error
	Value(key interface{}) interface{}
}


// create
// 1,Background // BASE Context
// 2,TODO

// 衍生context

// withCancel
// cancelFunc & Done  发送cancel信号 Done 返回一个通道 ， 协作起来 对子协程传递cancel信号

// DEMO1:
func Stream(ctx context.Context , out chan <- flag.Value){
	for{
		//v, err=
		select {
		case <- ctx.Done():
			//case out <- v:

		}
	}

}
// 子协程不停地运行并检查当前任务是否被取消，若被取消则结束当前任务并返回。


//WithDeadLine  WithTimeout

// 于withCancel 类似 ， 额外接收一个消亡时间， 超时时间

// WithValue 在上下文传递

