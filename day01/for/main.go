package main

import "fmt"

//for循环

func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//变种2
	var j = 5
	for ; j < 10 ;{
		fmt.Println(j)
		j++
	}

	//变种3
	//for{
	//	fmt.Println("死循环")
	//}

	// for range 循环
	s := "hi,mylove"
	for i, v := range s {
		fmt.Printf("%d %c \n",i,v)
	}



}
