package main

import (
	"fmt"
	"time"
)

//函数的定义
//函数的类型
//两者需要区分

//另一个闭包的例子

//闭包是什么?
//闭包是一个函数,这个函数包含了他外部作用域的一个变量

//1.函数可以作为返回值
//2.函数内部查找变量的顺序,先在自己内部找,找不到往外层找
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func pointer(p *func(int) func(int) int) {
	fmt.Printf("the tmp best %p\n",&p)
}

func main() {
	start := time.Now().UnixNano()
	p := new(func(int) func(int)int)
	*p = adder
	fmt.Printf("the best %p\n",&p)
	pointer(p)
	fmt.Println((*p)(1)(100))
	fmt.Println((*p)(12)(1000))
	end := time.Now().UnixNano()
	fmt.Printf("this program cost : %d", end - start)
}
