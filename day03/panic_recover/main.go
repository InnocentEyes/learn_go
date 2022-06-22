package main

import "fmt"

/// panic 和 recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	//刚刚打开了一个数据库链接
	defer func() {
		if err := recover(); err != nil  {
			fmt.Println(err)
			fmt.Println("释放数据库资源...")
		}
	}()
	panic("error happen") //程序崩溃退出
	fmt.Println("b")
}

func funcC()  {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
