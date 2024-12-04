package main

import "fmt"

//example of nested structs and methods with different names

type student struct {
	name     string
	roll     int
	location string
	c        class
}

type class struct {
	name  string
	floor int
}

func (s student) show() {
	fmt.Println("Name: ", s.name)
	fmt.Println("roll: ", s.roll)
	fmt.Println("location: ", s.location)
	fmt.Println("Class Name: ", s.c.name)
	fmt.Println("Class Floor: ", s.c.floor)

}

func (c class) show() {
	fmt.Println("Class Name: ", c.name)
	fmt.Println("class Floor: ", c.floor)
}

func (s *student) updateStudentName(name string) {
	(*s).name = name
	fmt.Println("updated the name successfully")
	(*s).show()
}
