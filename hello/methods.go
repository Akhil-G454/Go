package main

import "fmt"

type details struct {
	name string
	roll int
}

type home struct {
	houseName string
	houseNO   int
}

func (a *details) updateName(s string) {
	(*a).name = s
	fmt.Println("updated name successfully")
	fmt.Println("updated struct is: ")
	fmt.Println(*a)
}

// struct and method should be in the same file.
//methods receives the details of struct as receiver.
// same like methods in classes. Here methods with receiver is used.
