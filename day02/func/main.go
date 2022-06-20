package main

import "fmt"

//函数

//函数存在的意义?
//函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，给它起个名字,每次用到它的时候直接用函数名调用就行了
// 使用函数能够让代码结构更清晰、更简洁

//函数定义
func sum(x int,y int)  (int) {
	return x + y
}

//有参数 没有返回值
func print(x int,y int){
	fmt.Println(x + y)
}

//没有参数没有返回值
func noParamNoReturn() {
	fmt.Println("f2")
}

//没有参数但是有返回值的
func noParamHasReturn() (res int) {
	return 1
}

//参数可以命名也可以不命名
//命名的返回值就相当于在函数中声明一个变量
func sumReturnHasName(x int,y int) (res int) {
	res = x + y
	return  //使用命名返回值可以return后省略
}

//多个返回值
func manyReturn(int,string) (int,string){
	return 1,"沙河"
}

//参数的类型简写
// 当参数中连续多个参数的类型一致时,我们可以将非最后一个参数的类型省略
func simpleParams(x,y,z int,m,n string,i,j bool) int {
	return x + y
}

//可变长参数
//可变长参数必须放在函数参数的最后
func canBeLong(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y类型是切片 []int
}

//Go语言中 函数没有默认参数这个概念
//js、python中有默认参数这个概念

func main() {
	res := simpleParams(1,2,3,"1","1",true,true)
	fmt.Println(res)
}
