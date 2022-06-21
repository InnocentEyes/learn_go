package main

import "fmt"

//回文字符串
func main() {
	var str string = "上海自来水来自海上"
	s := []rune(str)
	j := len(s) - 1
	var target bool = true
	for i := 0; i < len(s); i++ {
		if i > j {
			break
		}
		if s[i] != s[j]{
			target = false
			break
		}
		j--
	}
	if target {
		fmt.Println("是回文字符串")
	}else{
		fmt.Println("不是回文字符串")
	}

}
