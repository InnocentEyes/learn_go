package main

import "fmt"

//pointer

func main() {
	// 1.&取地符
	// 2.*: 根据地址取值
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n",p)
	//2. *: 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n",m)

	var a1 *int // nil pointer
	fmt.Println(a1)
	var a2 = new(int)
	*a2 = 100
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 200
	fmt.Println(*a2)

	//new 函数申请一个内存地址

	//make和new区别
	//1.make和new都是用来申请内存的
	//2.new很少用,一般用来给基本数据类型申请内存,string\int。。。返回的是对应类型的指针(*string、*int)
	//3.make是用来给slice、go_map、chan申请内存的,make函数返回的是对应的这三个类型本身



}
