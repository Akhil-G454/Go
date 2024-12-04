package main

import "fmt"

func sl() {

	// slice := []int{1, 2, 3, 4}
	// fmt.Println(len(slice)) // Output: 3
	// fmt.Println(cap(slice)) // Output: 3
	// var output1, output2 = fmt.Println(len(slice))
	// fmt.Println(output1, output2)

	// slice = append(slice, 4)

	// slice = append(slice, 10, 11, 12, 13)
	// fmt.Println(len(slice)) // Output: 4
	// fmt.Println(cap(slice)) // Output: 6 (new array allocated)

	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(array, len(array))
	//creates an array of size 20 and it is fixed.

	slice := array[0:5]
	fmt.Println(slice, len(slice), cap(slice))
	//creates a slice of the original array which is just a reference of the original array
	// "Slice" var doesnt take any space in memory for its elements.

	slice[2] = 10
	fmt.Println(array, len(array))
	fmt.Println(slice, len(slice), cap(slice))
	//changes made to this slice are reflected in the original array because its
	//referencing the original array

	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	//This is the technique you have to follow to copy the slice to a new slice
	//it now only contains the elements you need and the capacity of the newslice is same
	//as number of elements in it.

}
