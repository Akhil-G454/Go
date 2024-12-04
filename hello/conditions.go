package main

import "fmt"

func condition() {

	if num := 16; num < 10 {
		fmt.Println("less than 10")

	} else if num < 15 {
		fmt.Println("less than 15")
	} else if num < 20 {
		fmt.Println("less than 20")
	} else {
		fmt.Println("greater than 20")
	}

}
