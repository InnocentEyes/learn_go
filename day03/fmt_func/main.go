package main

import "fmt"

//fmt

func main() {
	n,err := fmt.Println("abc")
	fmt.Println(n)
	fmt.Println(err)
	fmt.Print("abc")
	fmt.Println("----------")
	fmt.Println("shahe")
	fmt.Println("nazha")
	//fmt.Printf("格式化字符串",值)
	// %T: 查看类型
	// %d: 十进制数
	// %b: 二进制数
	// %o: 八进制数
	// %x: 十六进制数
	// %c: 字符
	// %s: 字符串
	// %p: 指针
	// %v: 值
	// %f: 浮点数
	// %t: 布尔值

	var m1 = make(map[string]int,1)
	m1["lixiang"] = 100
	fmt.Printf("%v\n",m1)
	fmt.Printf("%#v\n",m1)

	num := 90
	printPrecent(num)

	fmt.Printf("%v\n",100)
	fmt.Printf("%t\n",true) // %t --> 布尔值

	//整数 -> 字符
	fmt.Printf("%q\n",65) //该值对应的单引号括起来的go语法字符字面值,必要时会采用安全的转义字符

	//浮点数和复数
	fmt.Printf("%b\n",3.14159265354697) //科学计数法

	// 字符串
	fmt.Printf("%q\n","qzl有理想")

	/**
		宽度标识符
		%f
		%9f
		%.2f
		%9.2f
		%9.f
	 */
	v := 12.34
	fmt.Printf("%f\n",v)
	fmt.Printf("%9f\n",v)
	fmt.Printf("%.2f\n",v)
	fmt.Printf("%9.2f\n",v)
	fmt.Printf("%9.f\n",v)

	/**
		占位符
		  '+'
	      ''
	      '-'
	      '#'
	      'o'
	 */
	s := "qzl"
	fmt.Printf("%s\n",s)
	fmt.Printf("%5s\n",s)
	fmt.Printf("%-5s\n",s)
	fmt.Printf("%5.7s\n",s)
	fmt.Printf("%-5.7s\n",s)
	fmt.Printf("%5.2s\n",s)
	fmt.Printf("%05s\n",s)

	//获取用户输入
	fmt.Println("----------")
	str := ""
	fmt.Scan(&str)
	fmt.Println("用户输入的内容是:",str)

	var (
		name string
		age int
		class string
	)

	fmt.Scanf("%s %d %s",&name,&age,&class)
	fmt.Println(name,age,class)
	fmt.Scanln(&name,&age,&class)
	fmt.Println(name,age,class)
}

func printPrecent(num int) {
	fmt.Printf("%d%%\n",num)
}
