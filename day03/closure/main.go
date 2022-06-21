package main

import "fmt"

//重要的事情写三次
//	函数类型
//	函数类型
//	函数类型

//闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
func f3(f func(x, y int), x, y int) func() {
	return func() {
		f(x,y)
	}
}

func main() {
	a := f3(f2,100,200)//把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(a)
}
