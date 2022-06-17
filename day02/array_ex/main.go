package main

import "fmt"

//array数组练习题
//求数组[1,3,5,7,8]所有元素的和
// 找出数组中和为指定值得两个元素的下标,比如从数组[1,3,5,7,8]中 找出和为8的两个元素的下标分别为(0,3)和(1,2)

func main() {

	array := [5]int{1, 3, 5, 7, 8}
	sum := 0
	count := 0
	var result [2][2]int
	for _, value := range array {
		sum += value
	}
	for i := 0; i < len(array); i++ {
		for j := len(array) - 1; j > i; j-- {
			if array[i]+array[j] == 8 {
				result[count] = [2]int{i, j}
				count++
			}
		}
	}

	fmt.Printf("最终的结果为: %d\n", sum)
	fmt.Println(result)

}
