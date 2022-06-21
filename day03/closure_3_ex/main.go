package main

import "fmt"

func calc(base int) (add func(int)int,sub func(int)int){
	add = func(i int) int {
		base += i
		return base
	}
	sub = func(i int) int {
		base -= i
		return base
	}
	return
}

func main() {
	add,sub := calc(100)
	fmt.Println(add(1),sub(1))
	fmt.Println(add(2),sub(2))
	fmt.Println(add(3),sub(3))
	fmt.Println(add(4),sub(4))
	fmt.Println(add(5),sub(5))
}
