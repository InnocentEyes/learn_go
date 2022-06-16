package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println("欢迎光临.")
	}else{
		fmt.Println("sorry.您不能进来.")
	}

	//多个判断条件
	if age > 35{
		fmt.Println("人到中年")
	}else if age > 18 {
		fmt.Println("青年")
	}else {
		fmt.Println("好好学习! ")
	}

	if  age += 1; age > 20 {
		fmt.Println("欢迎光临.")
	}else {
		fmt.Println("sorry.您不能进来.")
	}
}
