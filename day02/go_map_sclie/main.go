package main

import "fmt"

//	map和slice组合
//	切片和map在使用的时候一定要初始化
func main() {

	//元素类型为map的切片
	var s1 = make([]map[int]string,1,10)
	fmt.Println(len(s1))
	//没有对内部的map做初始化
	s1[0] = make(map[int]string,1)
	s1[0][10] = "沙河"
	fmt.Println(s1)
	//值为切片类型的map
	var m1 = make(map[string][]int,10)
	m1["bz"] = []int{10,20,30}
	fmt.Println(m1)
}
