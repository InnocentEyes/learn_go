package main

import "fmt"

//defer

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hhhhh") //defer把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("hehehe") //一个函数可以有多个defer语句
	defer fmt.Println("heiheihei") //多个defer语句按照先进后出(后进先出)的顺序延迟执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
