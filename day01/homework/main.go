package main

import "fmt"

func main() {
	str := "hello,我是qzl"

	tmp := []rune(str)

	sum := 0

	for _, item := range tmp {
		if item > 256 {
			sum++
		}
	}
	fmt.Printf("有中文字 %d 个\n", sum)

	for i := 1; i <= 9; i++{
		for j := 1; j <= i ; j++{
			fmt.Printf("%d * %d = %d  ",j,i, i * j)
		}
		fmt.Println()

	}
}
