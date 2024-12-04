package main

import "fmt"

func update(name, street, city, state string, pincode int, a *address) {
	(*a).name = name
	(*a).street = street
	(*a).city = city
	(*a).state = state
	(*a).pincode = pincode
	fmt.Println("updated successfully")
	fmt.Println(*a)
}
