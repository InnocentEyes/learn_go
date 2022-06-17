package main

import "fmt"

//切片slice

func main() {
	// 切片的定义
	var s []int     //定义一个存放int类型元素的切片
	var s1 []string //定义一个存放string类型元素的切片
	fmt.Println(s, s1)
	fmt.Println(s == nil)
	fmt.Println(s1 == nil)
	//初始化
	s = []int{1, 2, 3}
	s1 = []string{"gz", "hp", "bl"}
	fmt.Println(s1)
	s1 = []string{"qzl"}
	fmt.Println(s1)

	fmt.Println(s == nil)
	fmt.Println(s1 == nil)

	//长度和容量
	fmt.Printf("len:%d cap:%d ", len(s), cap(s))
	fmt.Printf("len:%d cap:%d ", len(s1), cap(s1))

	//切片(slice)
	//切片指向了一个底层的数组。
	//切片的长度就是它元素的个数。
	//切片的容量是底层数组从切片的第一个元素到最后一个元素的数量

	// 2.由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s2 := a1[0:4] //基于一个数组切割,左包含右不包含,(左闭右开)
	s3 := a1[1:5] // =>
	s4 := a1[:4]
	s5 := a1[3:]
	s6 := a1[:]
	fmt.Println(s2, s3, s4, s5, s6)

	s7 := s4[2:]
	fmt.Printf("len:%d cap:%d ", len(s7), cap(s7))

	a1[0] = 100
	a1[4] = 200

	fmt.Println("结果是: ", a1, s2, s3, s4, s5, s6)
}
