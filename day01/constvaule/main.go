package main

import "fmt"

//常量
//定义了常量之后不能修改
//在程序运行期间不会改变
const pi = 3.1415926
//e := 123
//简短变量声明,只能在函数里面用

const(
	statusOk = 200
	notFound = 404
	not
	notin = 100
)


// 批量声明常量时,如果某一行声明后没有赋值,默认就和上一行一致 ---> 这个特性只有在常量时才有
const (
	n1 = 200
	n2
	n3
)


//iota:类似枚举
const (
	a1 = iota //0
	a2 //1
	a3 //2
)

const (
	b1 = iota // 0
	b2 //1
	_ //2
	b3 //3
)
//插队
const(
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

const (
	d1,d2 = iota + 1,iota + 2 //d1:1 d2:2

	d3,d4 = iota + 1,iota + 2 //d3:2 d4:4
)

//定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("statusOk: ",statusOk," notFount: ",notFound," not: ",not," not in: ",notin)
	fmt.Println("n1: ",n1," n2: ",n2," n3: ",n3)
	fmt.Println("a1: ",a1," a2: ",a2," a3: ",a3)
	fmt.Println("b1: ",b1," b2: ",b2," b1: ",b2)
	fmt.Println("c1: ",c1," c2: ",c2," c3: ",c3," c4: ",c4)
	fmt.Println("d1: ",d1," d2: ",d2," d3: ",d3," d4: ",d4)
}
