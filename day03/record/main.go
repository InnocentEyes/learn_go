package main

import "fmt"

//day03复习


func main() {
	var name string
	name = "理想"
	fmt.Println(name)
	var ages [30]int //声明了一个变量ages,它是[30]int类型
	ages = [30]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(ages)
	var age2 = [...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(age2)
	age3 := [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(age3)
	age4 := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(age4)
	age5 := [...]int{1:100,99:200}
	fmt.Println(age5)
	//二维数组
	var a1 [3][2]int = [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	} // [[1,2],[3,4]]
	fmt.Println(a1)
	//多维数组只有第二层可以使用多维数组
	a2 := [...][2]int{
		{1,2},
		{3,4},
		{5,6},
		{7,8},
		{9,10},
	}
	fmt.Println(a2)

	//数组是值类型
	x := [3]int{1,2,3}
	y := x //? 把x的值拷贝了一份给了y
	y[1] = 200 //修改的是副本y 并不影响x
 	f1(x)
	fmt.Println(x)
}
 func f1(a [3]int){
	 //Go 语言中的函数传递的都是值(Ctrl+C Crtl+V)
	 a[1] = 100
}
