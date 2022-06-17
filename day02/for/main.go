package main

import "fmt"

//流程控制之跳出for循环
func main() {
	// 当i =5 时就跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	fmt.Println("--------------------")
	//当i ==5 时跳过此次for循环(不执行for循环内部的打印语句).继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
