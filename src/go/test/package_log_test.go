package test

import (
	"log"
	"os"
	"testing"
)

/**
 * @author
 * @date 2020/10/11 15:59
 * create by Goland
 * @version 1.0
 */ 
 
// log 模块

func TestLogM1(t *testing.T) {
	log.Println("hello log!")
}


// Logger
// 创建不同 的logger 满足不同的需求
// New


// 一条日志由三个部分组成{日志前缀} {标识1} {标识2} ... {标识n} {日志内容}

// prefix ， flag

// DEMO2:
func TestLogM2(t *testing.T) {
	prefix := "[TEST LOGGER GEN]"
	logger := log.New(os.Stdout , prefix , log.LstdFlags | log.Lshortfile)

	logger.Println("hello world")
}


// 输出方式
// print
// fatal print 之后 os.exit(1) 之前
// panic print之后



// 自定义分级
func TestLogM3(t *testing.T) {
	var(
		logger = log.New(os.Stdout , "INFO:" , log.Lshortfile)
		infof = func(info string){
			logger.Println( info)
		}
	)

	infof("hello world")
}