package main

import "fmt"

func customPrint(n *int) int {
	var a int = 78
	var b bool = true
	var c string = "akhil"
	fmt.Println(a, b, c)
	fmt.Println("the value of the parameter is being changed")
	*n = 200
	fmt.Printf("the value is changed to %v \n", *n)

	return 0
}

func customValue(n *int) (square int, cube int) {
	square = *n * *n
	cube = *n * *n * *n
	return
}
