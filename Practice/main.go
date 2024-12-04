package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "helloo"
	var str2 string = "hello"
	fmt.Println(strings.Compare(str1, str2))
	if str1[0] == 104 {
		fmt.Printf("true")
	}
}
