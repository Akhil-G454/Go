package main

import "fmt"

func arrays() {
	var array = [...]int{1, 2, 3, 4, 5}
	arrays := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arrays)

	array[2] = 40
	fmt.Println(array)
	for _, j := range arrays {
		fmt.Println(j)
	}

	slice := arrays[1:4]

	fmt.Println("Memory addresses of array elements:")
	for i := 0; i < len(arrays); i++ {
		fmt.Printf("arrays[%d]: %p\n", i, &arrays[i])
	}

	fmt.Println("\nMemory addresses of slice elements:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]: %p\n", i, &slice[i])
	}

}
