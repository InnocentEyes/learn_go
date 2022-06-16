package main

import "fmt"

/**
	uint8 无符号整型 8位整型
	uint16 无符号整型 16位整型
	uint32 无符号整型 32位整型
	uint64 无符号整型 64位整型
	int8 符号整型 8位整型
	int16 符号整型 16位整型
	int32 符号整型 32位整型
	int64 符号整型 64位整型
 */

/**
	uint 32位操作系统上就是uint32,64位操作系统上就是uint64
	int	32位操作系统上就是int32,64位操作系统上就是int64
	uintstr 无符号整型,用于存放一个指针
 */

func main() {

	//十进制
	i1 := 10 // 正常开头
	fmt.Printf("%d \n",i1)
	fmt.Printf("%o \n",i1) //将十进制转换成八进制
	fmt.Printf("%x \n",i1) //将十进制转换成十六进制
	fmt.Printf("%b \n",i1) //将十进制转换成二进制

	fmt.Println("-----------------------------------")

	//八进制
	i2 := 077 // 0...
	fmt.Printf("%d \n",i2) //将八进制转换成十进制
	fmt.Printf("%o \n",i2)
	fmt.Printf("%x \n",i2) //将八进制转换成十六进制
	fmt.Printf("%b \n",i2) //将八进制转换成二进制

	fmt.Println("-----------------------------------")

	//十六进制 0x...
	i3 := 0x1234567
	fmt.Printf("%d \n",i3) //将十六进制转换成十进制
	fmt.Printf("%o \n",i3) //将十六进制转换成八进制
	fmt.Printf("%x \n",i3)
	fmt.Printf("%b \n",i3) //将十六进制转换成二进制

	//查看变量的类型
	fmt.Printf("%T \n",i3)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8类型,否则就是默认为int类型
	fmt.Printf("%T \n",i4)

}
