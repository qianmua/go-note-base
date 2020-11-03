package test

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)
/**
 * @author
 * @date 2020/11/3 16:37
 * create by Goland
 * @version 1.0
 */ 
 
// redis

var (
	Pool *redis.Pool
)


func init(){
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool{

	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn , error){
			c,err := redis.Dial("tcp" , server)
			if err != nil {
				return nil , err
			}
			return c , err

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		signal.Notify(c, syscall.SIGKILL)
		go func() {
			<-c
			Pool.Close()
			os.Exit(0)
		}()
}

func Get(key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func TestRedisM1(t *testing.T) {
	test, err := Get("test")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(test)
}
