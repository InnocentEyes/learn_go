package main

import "fmt"

// switch 流程控制语句

func main() {
	//switch
	//简化大量判断,一个变量和具体的值作比较
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效的数字")
	}

	//switch 简化上面的代码
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	//另一种写法
	switch i := 1; i {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	switch i := 7; i {
	case 1:
	case 3:
	case 5:
	case 7:
	case 9:
		fmt.Println("奇数")
	case 2:
	case 4:
	case 6:
	case 8:
	case 10:
		fmt.Println("偶数")
	}

	switch i := 7; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	// fallthrough语法可以执行满足条件的case的下一个case,是为了兼容C语言中的case设计的
	switch s := "a"; {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
