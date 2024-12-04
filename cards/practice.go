package main

import "fmt"

func customPrint() int { // Rename function to avoid conflict
	var a string = "akhil"
	var b int = 78
	var c bool = true
	fmt.Println(a, b, c)
	return 0
}
