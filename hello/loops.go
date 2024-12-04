package main

import "fmt"

func loop() {
	var list = []string{"first", "second", "third", "fourth", "fifth", "sixth"}
	for index, value := range list {
		fmt.Println(index, value)
		if value == "third" {
			break
		}
	}

}
