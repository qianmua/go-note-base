package test

import (
	"fmt"
	"testing"
)

/**
 * @author
 * @date 2020/11/6 14:01
 * create by Goland
 * @version 1.0
 */ 
 
// 语法

func TestBaseM1(t *testing.T) {
	// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := 5; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
}

func TestBaseM2(t *testing.T) {
	// 标签名是大小写敏感的。
	i := 0
	Here:   //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	// goto 符号
	goto Here   //跳转到Here去
}

// for
// 有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1
// break和continue还可以跟着标号，用来跳到多重循环中的外层循环
func TestBaseM3(t *testing.T) {

	var map1 map[string]string
	map1 = make(map[string]string)
	map1["A"] = "A"
	map1["B"] = "B"
	for k,v:=range map1 {
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}

	for _,v:=range map1 {
		fmt.Println("map's val:",v)
	}
}


// Go里面switch默认相当于每个case最后带有break
// 匹配成功后不会自动向下执行其他case，而是跳出整个switch
// 但是可以使用fallthrough强制执行后面的case代码



// 变参
func myfunc(arg ...int) {
	fmt.Println("$hhh")
	//panic("hhhh$user")
	fmt.Println(`hhhhhh${dsa}`)

}
// 变量arg是一个int的slice


// OOP
type Cicle struct {
	width , height float64
}

func (c Cicle) area() float64{
	return c.width * c.height
}

// 全部实现 - -！
func (c *Cicle)SayHi(){
	fmt.Println("hhhhh")
}

func (c *Cicle)Sing(lyrics string){

}
func (c *Cicle)Guzzle(beerStein string){

}



type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段Human
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段Human
	company string
	money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}
// 某个对象实现了某个接口的所有方法，则此对象就实现了此接口
// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

/// 断言 data.(type)
