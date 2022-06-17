package main

import (
	"fmt"
)

//运算符

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句,不能放在=的右边赋值
	b-- //单独的语句,不能放在=的右边复制

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型,相同的变量才能比较
	fmt.Println(a != b) //不等于
	fmt.Println(a >= b) //大于等于
	fmt.Println(a > b)  //大于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a < b)  //小于

	//逻辑运算符
	//如果年龄大于18岁并且年龄小于60岁
	if age := 20; age > 20 && age < 60 {
		fmt.Println("hello,world!")
	} else {
		fmt.Println("hello,world!(the age less than 20 :) )")
	}
	//如果年龄大于18岁 或者 年龄小于60岁
	if age := 20; age < 18 || age > 60 {
		fmt.Println("hello,world!(the age less than 20 :) )")
	} else {
		fmt.Println("hello,world!")
	}

	// not 取反,原来为真就为假 原来为假就为真
	//var isMarried bool
	//isMarried = true
	//var isMarried = true
	//var isMarried bool = true
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	//位运算符: 针对的是二进制数
	// 5的二进制表示: 101
	// 2的二进制表示: 010

	// &: 按位与
	fmt.Println(5 & 2) // 000

	// |: 按位或
	fmt.Println(5 | 2) // 111

	// ^: 按位异或 (两位不一样则为1)
	fmt.Println(5 ^ 2) // 111

	// <<: 将二进制位左移指定位数
	fmt.Println(5 << 1) //1010 -> 10
	fmt.Println(1 << 10)
	// >>:将二进制位右移制定的位数
	fmt.Println(5 >> 3) // 0
	fmt.Println(5 >> 2) // 1
	fmt.Println(5 >> 1) // 2

	m := int8(1)         //只能存八位
	fmt.Println(m << 10) // 0

	//192.168.1.1
	//权限 文件操作会讲为运算实际的应用
	//0644
	//赋值运算符,用来给变量赋值的
	var x int32
	x = 10
	x += 1  // x = x + 1
	x -= 1  // x = x - 1
	x *= 2  // x = x * 2
	x /= 2  // x = x / 2
	x %= 2  // x = x % 2
	x <<= 2 // x = x << 2
	x >>= 2 // x = x >> 2
	x |= 2  // x = x | 2
	x &= 2  // x = x & 2
	x ^= 2  // x = x ^ 2
	fmt.Println(x)

}
