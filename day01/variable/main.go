package main

import "fmt"

//Go语言推荐使用驼峰式命名法

//var student_name string

var studentName string

//var StudentName string

//单独声明
var name string
var age int
var isOk bool

//批量声明
//var (
//	code  int
//	msg   string
//	value string
//)

var a,_ = 1,2

func main() {
	name = "qzl"
	age = 16
	isOk = true
	studentName = "邱泽林"
	//Go语言中非全局变量声明必须使用,不使用就编译不过去
	fmt.Print(isOk)//在终端中输出要打印的内容
	fmt.Printf("name: %s,studentName: %s \n", name,studentName)//%s 占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)//打印完指定的内容之后会在后面加一个换行符

	//声明变量同时赋值
	var s1 string = "whb"

	fmt.Println(s1)

	//类型推导(根据值判断变量是什么类型)
	var s2 = "20"

	fmt.Println(s2)

	//简短变量声明,只能在函数里面用
	s3 := "困死了"

	//s1 := "10" 同一个作用域({})不能重复声明同名的变量

	// _ 匿名变量 lua中叫做哑元变量

	fmt.Println(s3)

	fmt.Printf("_的值是 %c",'c')
}
