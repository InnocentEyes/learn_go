package main

import "fmt"

/**
	uint8类型 byte 就是ascall

 */

func main() {
	s := "hello,world"

	n := len(s)

	fmt.Println(n)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%c \n",s[i])
	}


	for _,c := range s {
		fmt.Printf("%c\n",c)
	}

	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	//字符串修改
	s2 := "白萝卜" // => [白 萝 卜]
	s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c1 += "黑"
	c2 := '红'

	fmt.Println("c1: %T c4: %T \n",c1,c2)

	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3: %T c4:%T \n",c3,c4)
	fmt.Printf("%d\n",c4)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Println("%T\n",f)


}
