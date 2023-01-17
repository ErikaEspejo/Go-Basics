package main

import "fmt"

func main() {
	// Slice -> Dynamic Array
	numbers := []int{1, 2, 3} //[empty] it means it is a slice
	fmt.Println(numbers, len(numbers), cap(numbers))

	// Add elements to slice
	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers, len(numbers), cap(numbers))

	// Sub slice
	subSlice := numbers[:2]
	fmt.Println(subSlice)

	/* Subslices are considered as children of a father slice, so when there is any modification to the father slice,
	the children also are going to be modified */

	numbers[0] = 100 // it's been modified after the subslice is created

	fmt.Println(subSlice, len(subSlice), cap(subSlice)) // Prints [100 2]
	fmt.Println(numbers)                                // Prints [100 2 3 4 5 6]

	/* Both were affected by the modification because subSlice is a child of numbers. Even when the modification
	has been made after it was created. */

	newSliceToAppend := []int{7, 8, 9}
	numbers = append(numbers, newSliceToAppend...) // To append an array/slice into another slice, it should be detructured with "..." to match with the slice datatype
	subSlice = append(subSlice, newSliceToAppend...)

	fmt.Println(numbers)
	fmt.Println(subSlice)

	// Slices with for ... range
	darthListStarWars := []string{"Darth Vader", "Darth Maul", "Darth Sidius", "Darth Revan"}

	for index, value := range darthListStarWars {
		fmt.Println("Index:", index, "- Name:", value)
	}

	// To escape index, just use "_" instead of using a variable
	for _, value := range darthListStarWars {
		fmt.Println("Name:", value)
	}

	// To escape value data, just don't use any variable
	for index := range darthListStarWars {
		fmt.Println("Index:", index)
	}
}
