package main

import "fmt"

//函数: 一段代码的封装

func printSomething() {
	fmt.Println("Hello gz")
}

func printName(name string) {
	fmt.Println("hello",name)
}

//带参数和返回值的函数
func sum(x, y int) int {
	sum := x + y
	return sum
}

//可变参数
func f5(title string, y ...int)  int {
	fmt.Println(y) //y是一个int类型的切片
	return 1
}

//命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值,那么在函数中可以直接使用返回值变量
	return //return 后面可以省略返回值变量
}

//Go语言中支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main()  {
	printSomething()
	printName("qzl")
	sum(100,200)
	fmt.Println(sum(100,200)) //调用函数

	ret := sum(100,200)
	fmt.Println(ret)

	f5("qzl",1,2,3,4,5,6)

	//在一个命名的函数中不能够再声明命名函数
}
