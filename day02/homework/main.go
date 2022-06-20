package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "how do you do"
	result := strings.Split(str," ")
	strMap := make(map[string]int,16)
	for _, value := range result {
		if v, ok := strMap[value]; ok {
			v++
			strMap[value] = v
			continue
		}
		strMap[value] = 1
	}
	fmt.Println(strMap)
}
