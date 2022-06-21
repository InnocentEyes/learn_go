package main

import (
	"fmt"
	"strings"
)

//
func makeSuffixFuc(suffix string) func(string) string {
	return func(name string) (realName string) {
		if strings.HasSuffix(name,suffix) {
			return name
		}
		realName = name + suffix
		return
	}
}

func main() {
	jpgFunc := makeSuffixFuc(".jpg")
	pngFunc := makeSuffixFuc(".png")
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(pngFunc("test"))

}
