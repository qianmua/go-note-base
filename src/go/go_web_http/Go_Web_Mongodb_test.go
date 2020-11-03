package test

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

/**
 * @author
 * @date 2020/11/3 17:01
 * create by Goland
 * @version 1.0
 */ 

// mongodb

type Person struct {
	Name string
	Phone string
}

func TestMongoM1(t *testing.T) {
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	//session, err := mgo.Dial("server1.example.com,server2.example.com")
	session, err := mgo.Dial("mongodb://47.96.230.44:27017/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

/**
	阶乘
 */
func TestJieM2(t *testing.T) {
	i := 10
	res := Factorial(int64(i))
	fmt.Println(res)
}

func Factorial( n int64)(res int64){
	if n > 0 {
		res = n * Factorial(n - 1)
		return
	}
	return 1
}
