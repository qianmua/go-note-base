package test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

/**
 * @author
 * @date 2020/11/10 17:06
 * create by Goland
 * @version 1.0
 */ 
 
//chan

func TestChanM1(t *testing.T) {
	ch := make(chan interface{})

	go func() {
		fmt.Println("B")

		ch <- 10

		fmt.Println("C ")
	}()

	fmt.Println("A")

	data := <- ch
	// 忽略接收的值
	// <- ch

	fmt.Println(data)

	fmt.Println("D")

}

func TestM22(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")


}