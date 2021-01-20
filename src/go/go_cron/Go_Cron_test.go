package go_cron

import (
	"github.com/robfig/cron"
	"log"
	"testing"
	"time"
)

/**
 *
 * @author : jinchao.hu
 * @date 2021/1/20 14:51
 * @version 1.0
 */

/**
 go cron example
 */
func TestCron(t *testing.T) {
	log.Println(" shc Starting...")
	// crate new Cron
	c := cron.New()
	// addFunc or Job
	c.AddFunc("*/15 * * * * *" , printHello)

	c.Start()
	defer c.Stop()

	// 使用Timer 实现
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <- t1.C :
			t1.Reset(time.Second * 10)
			printHello()
		}
	}


}


func printHello(){
	log.Println(" Hello world!")
}