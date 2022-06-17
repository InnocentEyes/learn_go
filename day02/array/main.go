package main

import (
	"fmt"
)

//数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// 数组的长度是数组类型的一部分
// 数组是值对象 值类型

func main() {

	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("first: %T , second: %T", a1, a2)

	// 数组的初始化
	//如果不初始化: 默认元素都是零值(布尔值: false,整型和浮点型都是0,字符串: "")
	fmt.Println(a1, a2)
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2
	// a10 := [9]int{0,1,2,3,4,5,6,7,8}
	//根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a10)
	// 3.初始化方式3,根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"} //索引: 0~2 citys[0] citys[1] citys[2]
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//多维数组
	// [[1 2] [3 4] [5 6]]
	var app = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(app)

	//多维数组的遍历
	for _, v1 := range app {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3} //[1 2 3]
	b2 := b1              //[1 2 3] Ctrl+C Ctrl+V
	b2[0] = 100           // b2:[100 2 3]
	fmt.Println(b1, b2)   // b1:?

	//数组支持 == !=
}
