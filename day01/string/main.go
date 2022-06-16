package main

import (
	"fmt"
	"strings"
)

/**
	// 字符串
	s := "Hello 广州"
	c1 := 'h'
	c2 := '1'
	c3 := '粤'

	// 字节: 1字节 = 8Bit(8个二进制位)
	//1个字符 'A' = 1个字节
	//1个utf-8编码的汉字 '沙' = 一般占三个字节
 */

func main() {
	// \本来是具有特殊含义的,我应该告诉程序我写的\就是一个单独的\
	path := "\"E:\\Go\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	//多行的字符串

	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`

	fmt.Println(s2)

	s3 := `E:\Go`

	fmt.Println(s3)

	//字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "dsb"

	ss := name + world

	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s",name,world)

	fmt.Println(ss1)

	//分隔
	ret := strings.Split(s3,"\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss,"理性"))
	fmt.Println(strings.Contains(ss,"理想"))

	//前缀和后缀判断
	//前缀
	fmt.Println(strings.HasPrefix(ss,"理想"))
	//后缀
	fmt.Println(strings.HasSuffix(ss,"dsb"))

	//判断子串出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.Index(s4,"b"))
	fmt.Println(strings.LastIndex(s4,"b"))

	//拼接
	fmt.Println(strings.Join(ret,"\\"))


}
