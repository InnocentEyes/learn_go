package main

import "fmt"

func main() {
	//跳出多层for循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出for循环
		}
	}

	//goto + label 实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto code //跳到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	return

code: //label标签
	fmt.Println("over")
}
